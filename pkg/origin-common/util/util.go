package util

import (
	"strings"

	"github.com/docker/distribution"
	"github.com/docker/distribution/context"
	"github.com/docker/distribution/digest"
	"github.com/docker/distribution/registry/api/errcode"
	disterrors "github.com/docker/distribution/registry/api/v2"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	imageapiv1 "github.com/openshift/api/image/v1"
	"github.com/openshift/image-registry/pkg/dockerregistry/server/client"
)

// ImageWithMetadata mutates the given image. It parses raw DockerImageManifest data stored in the image and
// fills its DockerImageMetadata and other fields.
func ImageWithMetadata(image *imageapiv1.Image) error {
	if len(image.DockerImageManifest) == 0 {
		return nil
	}

	ReorderImageLayers(image)

	if len(image.DockerImageLayers) > 0 && image.DockerImageMetadata.Size > 0 && len(image.DockerImageManifestMediaType) > 0 {
		glog.V(5).Infof("Image metadata already filled for %s", image.Name)
		return nil
	}

	manifest := imageapiv1.DockerImageManifest{}
	if err := json.Unmarshal([]byte(image.DockerImageManifest), &manifest); err != nil {
		return err
	}

	err := fillImageLayers(image, manifest)
	if err != nil {
		return err
	}

	switch manifest.SchemaVersion {
	case 1:
		image.DockerImageManifestMediaType = schema1.MediaTypeManifest

		if len(manifest.History) == 0 {
			// It should never have an empty history, but just in case.
			return fmt.Errorf("the image %s (%s) has a schema 1 manifest, but it doesn't have history", image.Name, image.DockerImageReference)
		}

		v1Metadata := imageapiv1.DockerV1CompatibilityImage{}
		if err := json.Unmarshal([]byte(manifest.History[0].DockerV1Compatibility), &v1Metadata); err != nil {
			return err
		}

		image.DockerImageMetadata.ID = v1Metadata.ID
		image.DockerImageMetadata.Parent = v1Metadata.Parent
		image.DockerImageMetadata.Comment = v1Metadata.Comment
		image.DockerImageMetadata.Created = v1Metadata.Created
		image.DockerImageMetadata.Container = v1Metadata.Container
		image.DockerImageMetadata.ContainerConfig = v1Metadata.ContainerConfig
		image.DockerImageMetadata.DockerVersion = v1Metadata.DockerVersion
		image.DockerImageMetadata.Author = v1Metadata.Author
		image.DockerImageMetadata.Config = v1Metadata.Config
		image.DockerImageMetadata.Architecture = v1Metadata.Architecture
	case 2:
		image.DockerImageManifestMediaType = schema2.MediaTypeManifest

		if len(image.DockerImageConfig) == 0 {
			return fmt.Errorf("dockerImageConfig must not be empty for manifest schema 2")
		}

		config := imageapiv1.DockerImageConfig{}
		if err := json.Unmarshal([]byte(image.DockerImageConfig), &config); err != nil {
			return fmt.Errorf("failed to parse dockerImageConfig: %v", err)
		}

		image.DockerImageMetadata.ID = manifest.Config.Digest
		image.DockerImageMetadata.Parent = config.Parent
		image.DockerImageMetadata.Comment = config.Comment
		image.DockerImageMetadata.Created = config.Created
		image.DockerImageMetadata.Container = config.Container
		image.DockerImageMetadata.ContainerConfig = config.ContainerConfig
		image.DockerImageMetadata.DockerVersion = config.DockerVersion
		image.DockerImageMetadata.Author = config.Author
		image.DockerImageMetadata.Config = config.Config
		image.DockerImageMetadata.Architecture = config.Architecture
	default:
		return fmt.Errorf("unrecognized Docker image manifest schema %d for %q (%s)", manifest.SchemaVersion, image.Name, image.DockerImageReference)
	}

	layerSet := sets.NewString()
	if manifest.SchemaVersion == 2 {
		layerSet.Insert(manifest.Config.Digest)
		image.DockerImageMetadata.Size = int64(len(image.DockerImageConfig))
	} else {
		image.DockerImageMetadata.Size = 0
	}
	for _, layer := range image.DockerImageLayers {
		if layerSet.Has(layer.Name) {
			continue
		}
		layerSet.Insert(layer.Name)
		image.DockerImageMetadata.Size += layer.LayerSize
	}

	return nil
}

// ReorderImageLayers mutates the given image. It reorders the layers in ascending order.
// Ascending order matches the order of layers in schema 2. Schema 1 has reversed (descending) order of layers.
func ReorderImageLayers(image *imageapiv1.Image) {
	if len(image.DockerImageLayers) == 0 {
		return
	}

	layersOrder, ok := image.Annotations[imageapiv1.DockerImageLayersOrderAnnotation]
	if !ok {
		switch image.DockerImageManifestMediaType {
		case schema1.MediaTypeManifest, schema1.MediaTypeSignedManifest:
			layersOrder = imageapiv1.DockerImageLayersOrderAscending
		case schema2.MediaTypeManifest:
			layersOrder = imageapiv1.DockerImageLayersOrderDescending
		default:
			return
		}
	}

	if layersOrder == imageapiv1.DockerImageLayersOrderDescending {
		// reverse order of the layers (lowest = 0, highest = i)
		for i, j := 0, len(image.DockerImageLayers)-1; i < j; i, j = i+1, j-1 {
			image.DockerImageLayers[i], image.DockerImageLayers[j] = image.DockerImageLayers[j], image.DockerImageLayers[i]
		}
	}

	if image.Annotations == nil {
		image.Annotations = map[string]string{}
	}

	image.Annotations[imageapiv1.DockerImageLayersOrderAnnotation] = imageapiv1.DockerImageLayersOrderAscending
}

// SchemeHost returns the scheme and host used to make this request.
// Suitable for use to compute scheme/host in returned 302 redirect Location.
// Note the returned host is not normalized, and may or may not contain a port.
// Returned values are based on the following information:
//
// Host:
// * X-Forwarded-Host/X-Forwarded-Port headers
// * Host field on the request (parsed from Host header)
// * Host in the request's URL (parsed from Request-Line)
//
// Scheme:
// * X-Forwarded-Proto header
// * Existence of TLS information on the request implies https
// * Scheme in the request's URL (parsed from Request-Line)
// * Port (if included in calculated Host value, 443 implies https)
// * Otherwise, defaults to "http"
func SchemeHost(req *http.Request) (string /*scheme*/, string /*host*/) {
	forwarded := func(attr string) string {
		// Get the X-Forwarded-<attr> value
		value := req.Header.Get("X-Forwarded-" + attr)
		// Take the first comma-separated value, if multiple exist
		value = strings.SplitN(value, ",", 2)[0]
		// Trim whitespace
		return strings.TrimSpace(value)
	}

	hasExplicitHost := func(h string) bool {
		_, _, err := net.SplitHostPort(h)
		return err == nil
	}

	forwardedHost := forwarded("Host")
	host := ""
	hostHadExplicitPort := false
	switch {
	case len(forwardedHost) > 0:
		host = forwardedHost
		hostHadExplicitPort = hasExplicitHost(host)

		// If both X-Forwarded-Host and X-Forwarded-Port are sent, use the explicit port info
		if forwardedPort := forwarded("Port"); len(forwardedPort) > 0 {
			if h, _, err := net.SplitHostPort(forwardedHost); err == nil {
				host = net.JoinHostPort(h, forwardedPort)
			} else {
				host = net.JoinHostPort(forwardedHost, forwardedPort)
			}
		}

	case len(req.Host) > 0:
		host = req.Host
		hostHadExplicitPort = hasExplicitHost(host)

	case len(req.URL.Host) > 0:
		host = req.URL.Host
		hostHadExplicitPort = hasExplicitHost(host)
	}

	port := ""
	if _, p, err := net.SplitHostPort(host); err == nil {
		port = p
	}

	forwardedProto := forwarded("Proto")
	scheme := ""
	switch {
	case len(forwardedProto) > 0:
		scheme = forwardedProto
	case req.TLS != nil:
		scheme = "https"
	case len(req.URL.Scheme) > 0:
		scheme = req.URL.Scheme
	case port == "443":
		scheme = "https"
	default:
		scheme = "http"
	}

	if !hostHadExplicitPort {
		if (scheme == "https" && port == "443") || (scheme == "http" && port == "80") {
			if hostWithoutPort, _, err := net.SplitHostPort(host); err == nil {
				host = hostWithoutPort
			}
		}
	}

	return scheme, host
}

// errMessageString is a part of error message copied from quotaAdmission.Admit() method in
// k8s.io/kubernetes/plugin/pkg/admission/resourcequota/admission.go module
const errQuotaMessageString = `exceeded quota:`
const errQuotaUnknownMessageString = `status unknown for quota:`
const errLimitsMessageString = `exceeds the maximum limit`

// IsErrorQuotaExceeded returns true if the given error stands for a denied request caused by detected quota
// abuse.
func IsErrorQuotaExceeded(err error) bool {
	if isForbidden := apierrs.IsForbidden(err); isForbidden || apierrs.IsInvalid(err) {
		lowered := strings.ToLower(err.Error())
		// the limit error message can be accompanied only by Invalid reason
		if strings.Contains(lowered, errLimitsMessageString) {
			return true
		}
		// the quota error message can be accompanied only by Forbidden reason
		if isForbidden && (strings.Contains(lowered, errQuotaMessageString) || strings.Contains(lowered, errQuotaUnknownMessageString)) {
			return true
		}
	}
	return false
}

// ParseDockerImageReference parses a Docker pull spec string into a
// DockerImageReference.
func ParseDockerImageReference(spec string) (DockerImageReference, error) {
	var ref DockerImageReference

	namedRef, err := parseNamedDockerImageReference(spec)
	if err != nil {
		return ref, err
	}

	ref.Registry = namedRef.Registry
	ref.Namespace = namedRef.Namespace
	ref.Name = namedRef.Name
	ref.Tag = namedRef.Tag
	ref.ID = namedRef.ID

	return ref, nil
}

// NamedDockerImageReference points to a Docker image.
type namedDockerImageReference struct {
	Registry  string
	Namespace string
	Name      string
	Tag       string
	ID        string
}

// parseNamedDockerImageReference parses a Docker pull spec string into a
// NamedDockerImageReference.
func parseNamedDockerImageReference(spec string) (namedDockerImageReference, error) {
	var ref namedDockerImageReference

	namedRef, err := reference.ParseNamed(spec)
	if err != nil {
		return ref, err
	}

	name := namedRef.Name()
	i := strings.IndexRune(name, '/')
	if i == -1 || (!strings.ContainsAny(name[:i], ":.") && name[:i] != "localhost") {
		ref.Name = name
	} else {
		ref.Registry, ref.Name = name[:i], name[i+1:]
	}

	if named, ok := namedRef.(reference.NamedTagged); ok {
		ref.Tag = named.Tag()
	}

	if named, ok := namedRef.(reference.Canonical); ok {
		ref.ID = named.Digest().String()
	}

	// It's not enough just to use the reference.ParseNamed(). We have to fill
	// ref.Namespace from ref.Name
	if i := strings.IndexRune(ref.Name, '/'); i != -1 {
		ref.Namespace, ref.Name = ref.Name[:i], ref.Name[i+1:]
	}

	return ref, nil
}

// JoinImageStreamImage creates a name for image stream image object from an image stream name and an id.
func JoinImageStreamImage(name, id string) string {
	return fmt.Sprintf("%s@%s", name, id)
}

// SplitImageStreamTag turns the name of an ImageStreamTag into Name and Tag.
// It returns false if the tag was not properly specified in the name.
func SplitImageStreamTag(nameAndTag string) (name string, tag string, ok bool) {
	parts := strings.SplitN(nameAndTag, ":", 2)
	name = parts[0]
	if len(parts) > 1 {
		tag = parts[1]
	}
	if len(tag) == 0 {
		tag = DefaultImageTag
	}
	return name, tag, len(parts) == 2
}

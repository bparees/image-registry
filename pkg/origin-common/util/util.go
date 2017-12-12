package util

import (
	"strings"

	"github.com/docker/distribution"
	"github.com/docker/distribution/context"
	"github.com/docker/distribution/digest"
	"github.com/docker/distribution/registry/api/errcode"
	disterrors "github.com/docker/distribution/registry/api/v2"
	quotautil "github.com/openshift/origin/pkg/quota/util"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	imageapiv1 "github.com/openshift/api/image/v1"
	"github.com/openshift/image-registry/pkg/dockerregistry/server/client"
)

const (
	// TODO: move these into github/openshift/api/image

	// DockerImageLayersOrderAnnotation describes layers order in the docker image.
	DockerImageLayersOrderAnnotation = "image.openshift.io/dockerLayersOrder"

	// DockerImageLayersOrderAscending indicates that image layers are sorted in
	// the order of their addition (from oldest to latest)
	DockerImageLayersOrderAscending = "ascending"

	// DockerImageLayersOrderDescending indicates that layers are sorted in
	// reversed order of their addition (from newest to oldest).
	DockerImageLayersOrderDescending = "descending"
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

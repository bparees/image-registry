package util

import (
	//"fmt"
	//"strings"

	//"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/util/sets"
	api "k8s.io/kubernetes/pkg/api"
	//imageapi "github.com/openshift/origin/pkg/image/apis/image"

	imageapiv1 "github.com/openshift/api/image/v1"
	imageapi "github.com/openshift/image-registry/pkg/origin-common/image/apis/image"
)

// ImageWithMetadata mutates the given image. It parses raw DockerImageManifest data stored in the image and
// fills its DockerImageMetadata and other fields.
// Copied from v3.7 github.com/openshift/origin/pkg/image/apis/image/v1/helpers.go
func ImageWithMetadata(image *imageapiv1.Image) error {
	// Check if the metadata are already filled in for this image.
	meta, hasMetadata := image.DockerImageMetadata.Object.(*imageapi.DockerImage)
	if hasMetadata && meta.Size > 0 {
		return nil
	}

	version := image.DockerImageMetadataVersion
	if len(version) == 0 {
		version = "1.0"
	}

	obj := &imageapi.DockerImage{}
	if err := runtime.DecodeInto(api.Codecs.UniversalDecoder(), image.DockerImageMetadata.Raw, obj); err != nil {
		return err
	}
	image.DockerImageMetadata.Object = obj
	image.DockerImageMetadataVersion = version
	return nil
}

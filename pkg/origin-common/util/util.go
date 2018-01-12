package util

import (
	"encoding/json"
	"fmt"
	"strings"

	//"k8s.io/apimachinery/pkg/api/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	//api "k8s.io/kubernetes/pkg/api"
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
	/*
		if err := runtime.DecodeInto(api.Codecs.UniversalDecoder(), image.DockerImageMetadata.Raw, obj); err != nil {
			return err
		}
	*/
	// TODO - bparees - should probably be using a scheme to decode this?
	if err := json.Unmarshal(image.DockerImageMetadata.Raw, obj); err != nil {
		return err
	}

	image.DockerImageMetadata.Object = obj
	image.DockerImageMetadataVersion = version
	return nil
}

// LatestImageTagEvent returns the most recent TagEvent and the tag for the specified
// image.
// Copied from v3.7 github.com/openshift/origin/pkg/image/apis/image/v1/helpers.go
func LatestImageTagEvent(stream *imageapiv1.ImageStream, imageID string) (string, *imageapiv1.TagEvent) {
	var (
		latestTagEvent *imageapiv1.TagEvent
		latestTag      string
	)
	for _, events := range stream.Status.Tags {
		if len(events.Items) == 0 {
			continue
		}
		tag := events.Tag
		for i, event := range events.Items {
			if imageapi.DigestOrImageMatch(event.Image, imageID) &&
				(latestTagEvent == nil || latestTagEvent != nil && event.Created.After(latestTagEvent.Created.Time)) {
				latestTagEvent = &events.Items[i]
				latestTag = tag
			}
		}
	}
	return latestTag, latestTagEvent
}

// ResolveImageID returns latest TagEvent for specified imageID and an error if
// there's more than one image matching the ID or when one does not exist.
// Copied from v3.7 github.com/openshift/origin/pkg/image/apis/image/v1/helpers.go
func ResolveImageID(stream *imageapiv1.ImageStream, imageID string) (*imageapiv1.TagEvent, error) {
	var event *imageapiv1.TagEvent
	set := sets.NewString()
	for _, history := range stream.Status.Tags {
		for i := range history.Items {
			tagging := &history.Items[i]
			if imageapi.DigestOrImageMatch(tagging.Image, imageID) {
				event = tagging
				set.Insert(tagging.Image)
			}
		}
	}
	switch len(set) {
	case 1:
		return &imageapiv1.TagEvent{
			Created:              metav1.Now(),
			DockerImageReference: event.DockerImageReference,
			Image:                event.Image,
		}, nil
	case 0:
		return nil, kerrors.NewNotFound(imageapiv1.Resource("imagestreamimage"), imageID)
	default:
		return nil, kerrors.NewConflict(imageapiv1.Resource("imagestreamimage"), imageID, fmt.Errorf("multiple images match the prefix %q: %s", imageID, strings.Join(set.List(), ", ")))
	}
}

// LatestTaggedImage returns the most recent TagEvent for the specified image
// repository and tag. Will resolve lookups for the empty tag. Returns nil
// if tag isn't present in stream.status.tags.
func LatestTaggedImage(stream *imageapiv1.ImageStream, tag string) *imageapiv1.TagEvent {
	if len(tag) == 0 {
		tag = imageapi.DefaultImageTag
	}
	for _, tagEventList := range stream.Status.Tags {
		if tagEventList.Tag != tag {
			continue
		}
		if len(tagEventList.Items) == 0 {
			return nil
		}
		return &tagEventList.Items[0]
	}

	return nil
}

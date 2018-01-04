package consts

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	// Limit that applies to images. Used with a max["storage"] LimitRangeItem to set
	// the maximum size of an image.
	LimitTypeImage corev1.LimitType = "openshift.io/Image"

	// DockerImageLayersOrderAscending indicates that image layers are sorted in
	// the order of their addition (from oldest to latest)
	DockerImageLayersOrderAscending = "ascending"

	// ManagedByOpenShiftAnnotation indicates that an image is managed by OpenShift's registry.
	ManagedByOpenShiftAnnotation = "openshift.io/image.managed"

	// ImageManifestBlobStoredAnnotation indicates that manifest and config blobs of image are stored in on
	// storage of integrated Docker registry.
	ImageManifestBlobStoredAnnotation = "image.openshift.io/manifestBlobStored"

	// DockerImageLayersOrderAnnotation describes layers order in the docker image.
	DockerImageLayersOrderAnnotation = "image.openshift.io/dockerLayersOrder"

	// The supported type of image signature.
	ImageSignatureTypeAtomicImageV1 string = "AtomicImageV1"
)

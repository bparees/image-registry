package consts

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	// Limit that applies to images. Used with a max["storage"] LimitRangeItem to set
	// the maximum size of an image.
	LimitTypeImage corev1.LimitType = "openshift.io/Image"
)

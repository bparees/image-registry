// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	v1 "github.com/openshift/api/security/v1"
	security "github.com/openshift/origin/pkg/security/apis/security"
	core_v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume,
		Convert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume,
		Convert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions,
		Convert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions,
		Convert_v1_IDRange_To_security_IDRange,
		Convert_security_IDRange_To_v1_IDRange,
		Convert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview,
		Convert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview,
		Convert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec,
		Convert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec,
		Convert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus,
		Convert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus,
		Convert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview,
		Convert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview,
		Convert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec,
		Convert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec,
		Convert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview,
		Convert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview,
		Convert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec,
		Convert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec,
		Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus,
		Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus,
		Convert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions,
		Convert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions,
		Convert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions,
		Convert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions,
		Convert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints,
		Convert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints,
		Convert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList,
		Convert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList,
		Convert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus,
		Convert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus,
		Convert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions,
		Convert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions,
	)
}

func autoConvert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume(in *v1.AllowedFlexVolume, out *security.AllowedFlexVolume, s conversion.Scope) error {
	out.Driver = in.Driver
	return nil
}

// Convert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume is an autogenerated conversion function.
func Convert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume(in *v1.AllowedFlexVolume, out *security.AllowedFlexVolume, s conversion.Scope) error {
	return autoConvert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume(in, out, s)
}

func autoConvert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume(in *security.AllowedFlexVolume, out *v1.AllowedFlexVolume, s conversion.Scope) error {
	out.Driver = in.Driver
	return nil
}

// Convert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume is an autogenerated conversion function.
func Convert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume(in *security.AllowedFlexVolume, out *v1.AllowedFlexVolume, s conversion.Scope) error {
	return autoConvert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume(in, out, s)
}

func autoConvert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(in *v1.FSGroupStrategyOptions, out *security.FSGroupStrategyOptions, s conversion.Scope) error {
	out.Type = security.FSGroupStrategyType(in.Type)
	out.Ranges = *(*[]security.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions is an autogenerated conversion function.
func Convert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(in *v1.FSGroupStrategyOptions, out *security.FSGroupStrategyOptions, s conversion.Scope) error {
	return autoConvert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(in, out, s)
}

func autoConvert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(in *security.FSGroupStrategyOptions, out *v1.FSGroupStrategyOptions, s conversion.Scope) error {
	out.Type = v1.FSGroupStrategyType(in.Type)
	out.Ranges = *(*[]v1.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions is an autogenerated conversion function.
func Convert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(in *security.FSGroupStrategyOptions, out *v1.FSGroupStrategyOptions, s conversion.Scope) error {
	return autoConvert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(in, out, s)
}

func autoConvert_v1_IDRange_To_security_IDRange(in *v1.IDRange, out *security.IDRange, s conversion.Scope) error {
	out.Min = in.Min
	out.Max = in.Max
	return nil
}

// Convert_v1_IDRange_To_security_IDRange is an autogenerated conversion function.
func Convert_v1_IDRange_To_security_IDRange(in *v1.IDRange, out *security.IDRange, s conversion.Scope) error {
	return autoConvert_v1_IDRange_To_security_IDRange(in, out, s)
}

func autoConvert_security_IDRange_To_v1_IDRange(in *security.IDRange, out *v1.IDRange, s conversion.Scope) error {
	out.Min = in.Min
	out.Max = in.Max
	return nil
}

// Convert_security_IDRange_To_v1_IDRange is an autogenerated conversion function.
func Convert_security_IDRange_To_v1_IDRange(in *security.IDRange, out *v1.IDRange, s conversion.Scope) error {
	return autoConvert_security_IDRange_To_v1_IDRange(in, out, s)
}

func autoConvert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview(in *v1.PodSecurityPolicyReview, out *security.PodSecurityPolicyReview, s conversion.Scope) error {
	if err := Convert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview is an autogenerated conversion function.
func Convert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview(in *v1.PodSecurityPolicyReview, out *security.PodSecurityPolicyReview, s conversion.Scope) error {
	return autoConvert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview(in, out, s)
}

func autoConvert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview(in *security.PodSecurityPolicyReview, out *v1.PodSecurityPolicyReview, s conversion.Scope) error {
	if err := Convert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview is an autogenerated conversion function.
func Convert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview(in *security.PodSecurityPolicyReview, out *v1.PodSecurityPolicyReview, s conversion.Scope) error {
	return autoConvert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview(in, out, s)
}

func autoConvert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(in *v1.PodSecurityPolicyReviewSpec, out *security.PodSecurityPolicyReviewSpec, s conversion.Scope) error {
	if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.ServiceAccountNames = *(*[]string)(unsafe.Pointer(&in.ServiceAccountNames))
	return nil
}

// Convert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec is an autogenerated conversion function.
func Convert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(in *v1.PodSecurityPolicyReviewSpec, out *security.PodSecurityPolicyReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(in, out, s)
}

func autoConvert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(in *security.PodSecurityPolicyReviewSpec, out *v1.PodSecurityPolicyReviewSpec, s conversion.Scope) error {
	if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.ServiceAccountNames = *(*[]string)(unsafe.Pointer(&in.ServiceAccountNames))
	return nil
}

// Convert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec is an autogenerated conversion function.
func Convert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(in *security.PodSecurityPolicyReviewSpec, out *v1.PodSecurityPolicyReviewSpec, s conversion.Scope) error {
	return autoConvert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(in, out, s)
}

func autoConvert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(in *v1.PodSecurityPolicyReviewStatus, out *security.PodSecurityPolicyReviewStatus, s conversion.Scope) error {
	if in.AllowedServiceAccounts != nil {
		in, out := &in.AllowedServiceAccounts, &out.AllowedServiceAccounts
		*out = make([]security.ServiceAccountPodSecurityPolicyReviewStatus, len(*in))
		for i := range *in {
			if err := Convert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AllowedServiceAccounts = nil
	}
	return nil
}

// Convert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus is an autogenerated conversion function.
func Convert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(in *v1.PodSecurityPolicyReviewStatus, out *security.PodSecurityPolicyReviewStatus, s conversion.Scope) error {
	return autoConvert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(in, out, s)
}

func autoConvert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(in *security.PodSecurityPolicyReviewStatus, out *v1.PodSecurityPolicyReviewStatus, s conversion.Scope) error {
	if in.AllowedServiceAccounts != nil {
		in, out := &in.AllowedServiceAccounts, &out.AllowedServiceAccounts
		*out = make([]v1.ServiceAccountPodSecurityPolicyReviewStatus, len(*in))
		for i := range *in {
			if err := Convert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AllowedServiceAccounts = nil
	}
	return nil
}

// Convert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus is an autogenerated conversion function.
func Convert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(in *security.PodSecurityPolicyReviewStatus, out *v1.PodSecurityPolicyReviewStatus, s conversion.Scope) error {
	return autoConvert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(in, out, s)
}

func autoConvert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview(in *v1.PodSecurityPolicySelfSubjectReview, out *security.PodSecurityPolicySelfSubjectReview, s conversion.Scope) error {
	if err := Convert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview is an autogenerated conversion function.
func Convert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview(in *v1.PodSecurityPolicySelfSubjectReview, out *security.PodSecurityPolicySelfSubjectReview, s conversion.Scope) error {
	return autoConvert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview(in, out, s)
}

func autoConvert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview(in *security.PodSecurityPolicySelfSubjectReview, out *v1.PodSecurityPolicySelfSubjectReview, s conversion.Scope) error {
	if err := Convert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview is an autogenerated conversion function.
func Convert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview(in *security.PodSecurityPolicySelfSubjectReview, out *v1.PodSecurityPolicySelfSubjectReview, s conversion.Scope) error {
	return autoConvert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview(in, out, s)
}

func autoConvert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(in *v1.PodSecurityPolicySelfSubjectReviewSpec, out *security.PodSecurityPolicySelfSubjectReviewSpec, s conversion.Scope) error {
	if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec is an autogenerated conversion function.
func Convert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(in *v1.PodSecurityPolicySelfSubjectReviewSpec, out *security.PodSecurityPolicySelfSubjectReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(in, out, s)
}

func autoConvert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(in *security.PodSecurityPolicySelfSubjectReviewSpec, out *v1.PodSecurityPolicySelfSubjectReviewSpec, s conversion.Scope) error {
	if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec is an autogenerated conversion function.
func Convert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(in *security.PodSecurityPolicySelfSubjectReviewSpec, out *v1.PodSecurityPolicySelfSubjectReviewSpec, s conversion.Scope) error {
	return autoConvert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(in, out, s)
}

func autoConvert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview(in *v1.PodSecurityPolicySubjectReview, out *security.PodSecurityPolicySubjectReview, s conversion.Scope) error {
	if err := Convert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview is an autogenerated conversion function.
func Convert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview(in *v1.PodSecurityPolicySubjectReview, out *security.PodSecurityPolicySubjectReview, s conversion.Scope) error {
	return autoConvert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview(in, out, s)
}

func autoConvert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview(in *security.PodSecurityPolicySubjectReview, out *v1.PodSecurityPolicySubjectReview, s conversion.Scope) error {
	if err := Convert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview is an autogenerated conversion function.
func Convert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview(in *security.PodSecurityPolicySubjectReview, out *v1.PodSecurityPolicySubjectReview, s conversion.Scope) error {
	return autoConvert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview(in, out, s)
}

func autoConvert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(in *v1.PodSecurityPolicySubjectReviewSpec, out *security.PodSecurityPolicySubjectReviewSpec, s conversion.Scope) error {
	if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec is an autogenerated conversion function.
func Convert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(in *v1.PodSecurityPolicySubjectReviewSpec, out *security.PodSecurityPolicySubjectReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(in, out, s)
}

func autoConvert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(in *security.PodSecurityPolicySubjectReviewSpec, out *v1.PodSecurityPolicySubjectReviewSpec, s conversion.Scope) error {
	if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec is an autogenerated conversion function.
func Convert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(in *security.PodSecurityPolicySubjectReviewSpec, out *v1.PodSecurityPolicySubjectReviewSpec, s conversion.Scope) error {
	return autoConvert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(in, out, s)
}

func autoConvert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(in *v1.PodSecurityPolicySubjectReviewStatus, out *security.PodSecurityPolicySubjectReviewStatus, s conversion.Scope) error {
	if in.AllowedBy != nil {
		in, out := &in.AllowedBy, &out.AllowedBy
		*out = new(api.ObjectReference)
		if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AllowedBy = nil
	}
	out.Reason = in.Reason
	if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus is an autogenerated conversion function.
func Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(in *v1.PodSecurityPolicySubjectReviewStatus, out *security.PodSecurityPolicySubjectReviewStatus, s conversion.Scope) error {
	return autoConvert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(in, out, s)
}

func autoConvert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(in *security.PodSecurityPolicySubjectReviewStatus, out *v1.PodSecurityPolicySubjectReviewStatus, s conversion.Scope) error {
	if in.AllowedBy != nil {
		in, out := &in.AllowedBy, &out.AllowedBy
		*out = new(core_v1.ObjectReference)
		if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AllowedBy = nil
	}
	out.Reason = in.Reason
	if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus is an autogenerated conversion function.
func Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(in *security.PodSecurityPolicySubjectReviewStatus, out *v1.PodSecurityPolicySubjectReviewStatus, s conversion.Scope) error {
	return autoConvert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(in, out, s)
}

func autoConvert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(in *v1.RunAsUserStrategyOptions, out *security.RunAsUserStrategyOptions, s conversion.Scope) error {
	out.Type = security.RunAsUserStrategyType(in.Type)
	out.UID = (*int64)(unsafe.Pointer(in.UID))
	out.UIDRangeMin = (*int64)(unsafe.Pointer(in.UIDRangeMin))
	out.UIDRangeMax = (*int64)(unsafe.Pointer(in.UIDRangeMax))
	return nil
}

// Convert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions is an autogenerated conversion function.
func Convert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(in *v1.RunAsUserStrategyOptions, out *security.RunAsUserStrategyOptions, s conversion.Scope) error {
	return autoConvert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(in, out, s)
}

func autoConvert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(in *security.RunAsUserStrategyOptions, out *v1.RunAsUserStrategyOptions, s conversion.Scope) error {
	out.Type = v1.RunAsUserStrategyType(in.Type)
	out.UID = (*int64)(unsafe.Pointer(in.UID))
	out.UIDRangeMin = (*int64)(unsafe.Pointer(in.UIDRangeMin))
	out.UIDRangeMax = (*int64)(unsafe.Pointer(in.UIDRangeMax))
	return nil
}

// Convert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions is an autogenerated conversion function.
func Convert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(in *security.RunAsUserStrategyOptions, out *v1.RunAsUserStrategyOptions, s conversion.Scope) error {
	return autoConvert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(in, out, s)
}

func autoConvert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(in *v1.SELinuxContextStrategyOptions, out *security.SELinuxContextStrategyOptions, s conversion.Scope) error {
	out.Type = security.SELinuxContextStrategyType(in.Type)
	if in.SELinuxOptions != nil {
		in, out := &in.SELinuxOptions, &out.SELinuxOptions
		*out = new(api.SELinuxOptions)
		if err := api_v1.Convert_v1_SELinuxOptions_To_api_SELinuxOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.SELinuxOptions = nil
	}
	return nil
}

// Convert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions is an autogenerated conversion function.
func Convert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(in *v1.SELinuxContextStrategyOptions, out *security.SELinuxContextStrategyOptions, s conversion.Scope) error {
	return autoConvert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(in, out, s)
}

func autoConvert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(in *security.SELinuxContextStrategyOptions, out *v1.SELinuxContextStrategyOptions, s conversion.Scope) error {
	out.Type = v1.SELinuxContextStrategyType(in.Type)
	if in.SELinuxOptions != nil {
		in, out := &in.SELinuxOptions, &out.SELinuxOptions
		*out = new(core_v1.SELinuxOptions)
		if err := api_v1.Convert_api_SELinuxOptions_To_v1_SELinuxOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.SELinuxOptions = nil
	}
	return nil
}

// Convert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions is an autogenerated conversion function.
func Convert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(in *security.SELinuxContextStrategyOptions, out *v1.SELinuxContextStrategyOptions, s conversion.Scope) error {
	return autoConvert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(in, out, s)
}

func autoConvert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints(in *v1.SecurityContextConstraints, out *security.SecurityContextConstraints, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Priority = (*int32)(unsafe.Pointer(in.Priority))
	out.AllowPrivilegedContainer = in.AllowPrivilegedContainer
	out.DefaultAddCapabilities = *(*[]api.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
	out.RequiredDropCapabilities = *(*[]api.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
	out.AllowedCapabilities = *(*[]api.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
	// INFO: in.AllowHostDirVolumePlugin opted out of conversion generation
	out.Volumes = *(*[]security.FSType)(unsafe.Pointer(&in.Volumes))
	out.AllowedFlexVolumes = *(*[]security.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
	out.AllowHostNetwork = in.AllowHostNetwork
	out.AllowHostPorts = in.AllowHostPorts
	out.AllowHostPID = in.AllowHostPID
	out.AllowHostIPC = in.AllowHostIPC
	if err := Convert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(&in.SELinuxContext, &out.SELinuxContext, s); err != nil {
		return err
	}
	if err := Convert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, s); err != nil {
		return err
	}
	if err := Convert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, s); err != nil {
		return err
	}
	if err := Convert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, s); err != nil {
		return err
	}
	out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.SeccompProfiles = *(*[]string)(unsafe.Pointer(&in.SeccompProfiles))
	return nil
}

func autoConvert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(in *security.SecurityContextConstraints, out *v1.SecurityContextConstraints, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Priority = (*int32)(unsafe.Pointer(in.Priority))
	out.AllowPrivilegedContainer = in.AllowPrivilegedContainer
	out.DefaultAddCapabilities = *(*[]core_v1.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
	out.RequiredDropCapabilities = *(*[]core_v1.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
	out.AllowedCapabilities = *(*[]core_v1.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
	out.Volumes = *(*[]v1.FSType)(unsafe.Pointer(&in.Volumes))
	out.AllowedFlexVolumes = *(*[]v1.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
	out.AllowHostNetwork = in.AllowHostNetwork
	out.AllowHostPorts = in.AllowHostPorts
	out.AllowHostPID = in.AllowHostPID
	out.AllowHostIPC = in.AllowHostIPC
	if err := Convert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(&in.SELinuxContext, &out.SELinuxContext, s); err != nil {
		return err
	}
	if err := Convert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, s); err != nil {
		return err
	}
	if err := Convert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, s); err != nil {
		return err
	}
	if err := Convert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, s); err != nil {
		return err
	}
	out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
	out.SeccompProfiles = *(*[]string)(unsafe.Pointer(&in.SeccompProfiles))
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

func autoConvert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList(in *v1.SecurityContextConstraintsList, out *security.SecurityContextConstraintsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]security.SecurityContextConstraints, len(*in))
		for i := range *in {
			if err := Convert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList is an autogenerated conversion function.
func Convert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList(in *v1.SecurityContextConstraintsList, out *security.SecurityContextConstraintsList, s conversion.Scope) error {
	return autoConvert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList(in, out, s)
}

func autoConvert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList(in *security.SecurityContextConstraintsList, out *v1.SecurityContextConstraintsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.SecurityContextConstraints, len(*in))
		for i := range *in {
			if err := Convert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList is an autogenerated conversion function.
func Convert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList(in *security.SecurityContextConstraintsList, out *v1.SecurityContextConstraintsList, s conversion.Scope) error {
	return autoConvert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList(in, out, s)
}

func autoConvert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(in *v1.ServiceAccountPodSecurityPolicyReviewStatus, out *security.ServiceAccountPodSecurityPolicyReviewStatus, s conversion.Scope) error {
	if err := Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(&in.PodSecurityPolicySubjectReviewStatus, &out.PodSecurityPolicySubjectReviewStatus, s); err != nil {
		return err
	}
	out.Name = in.Name
	return nil
}

// Convert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus is an autogenerated conversion function.
func Convert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(in *v1.ServiceAccountPodSecurityPolicyReviewStatus, out *security.ServiceAccountPodSecurityPolicyReviewStatus, s conversion.Scope) error {
	return autoConvert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(in, out, s)
}

func autoConvert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(in *security.ServiceAccountPodSecurityPolicyReviewStatus, out *v1.ServiceAccountPodSecurityPolicyReviewStatus, s conversion.Scope) error {
	if err := Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(&in.PodSecurityPolicySubjectReviewStatus, &out.PodSecurityPolicySubjectReviewStatus, s); err != nil {
		return err
	}
	out.Name = in.Name
	return nil
}

// Convert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus is an autogenerated conversion function.
func Convert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(in *security.ServiceAccountPodSecurityPolicyReviewStatus, out *v1.ServiceAccountPodSecurityPolicyReviewStatus, s conversion.Scope) error {
	return autoConvert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(in, out, s)
}

func autoConvert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(in *v1.SupplementalGroupsStrategyOptions, out *security.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	out.Type = security.SupplementalGroupsStrategyType(in.Type)
	out.Ranges = *(*[]security.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions is an autogenerated conversion function.
func Convert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(in *v1.SupplementalGroupsStrategyOptions, out *security.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	return autoConvert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(in, out, s)
}

func autoConvert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(in *security.SupplementalGroupsStrategyOptions, out *v1.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	out.Type = v1.SupplementalGroupsStrategyType(in.Type)
	out.Ranges = *(*[]v1.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions is an autogenerated conversion function.
func Convert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(in *security.SupplementalGroupsStrategyOptions, out *v1.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	return autoConvert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(in, out, s)
}

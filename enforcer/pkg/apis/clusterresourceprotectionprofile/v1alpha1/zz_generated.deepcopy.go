// +build !ignore_autogenerated

/*
Copyright 2020 IBM Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	protect "github.com/IBM/integrity-enforcer/enforcer/pkg/protect"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceProtectionProfile) DeepCopyInto(out *ClusterResourceProtectionProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceProtectionProfile.
func (in *ClusterResourceProtectionProfile) DeepCopy() *ClusterResourceProtectionProfile {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceProtectionProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceProtectionProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceProtectionProfileList) DeepCopyInto(out *ClusterResourceProtectionProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceProtectionProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceProtectionProfileList.
func (in *ClusterResourceProtectionProfileList) DeepCopy() *ClusterResourceProtectionProfileList {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceProtectionProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceProtectionProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceProtectionProfileSpec) DeepCopyInto(out *ClusterResourceProtectionProfileSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]*protect.Rule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.IgnoreServiceAccount != nil {
		in, out := &in.IgnoreServiceAccount, &out.IgnoreServiceAccount
		*out = make([]*protect.ServieAccountPattern, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.ProtectAttrs != nil {
		in, out := &in.ProtectAttrs, &out.ProtectAttrs
		*out = make([]*protect.AttrsPattern, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.UnprotectAttrs != nil {
		in, out := &in.UnprotectAttrs, &out.UnprotectAttrs
		*out = make([]*protect.AttrsPattern, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.IgnoreAttrs != nil {
		in, out := &in.IgnoreAttrs, &out.IgnoreAttrs
		*out = make([]*protect.AttrsPattern, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceProtectionProfileSpec.
func (in *ClusterResourceProtectionProfileSpec) DeepCopy() *ClusterResourceProtectionProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceProtectionProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceProtectionProfileStatus) DeepCopyInto(out *ClusterResourceProtectionProfileStatus) {
	*out = *in
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]*protect.Result, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceProtectionProfileStatus.
func (in *ClusterResourceProtectionProfileStatus) DeepCopy() *ClusterResourceProtectionProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceProtectionProfileStatus)
	in.DeepCopyInto(out)
	return out
}
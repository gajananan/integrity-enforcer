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
func (in *ResourceProtectionProfile) DeepCopyInto(out *ResourceProtectionProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceProtectionProfile.
func (in *ResourceProtectionProfile) DeepCopy() *ResourceProtectionProfile {
	if in == nil {
		return nil
	}
	out := new(ResourceProtectionProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceProtectionProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceProtectionProfileList) DeepCopyInto(out *ResourceProtectionProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceProtectionProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceProtectionProfileList.
func (in *ResourceProtectionProfileList) DeepCopy() *ResourceProtectionProfileList {
	if in == nil {
		return nil
	}
	out := new(ResourceProtectionProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceProtectionProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceProtectionProfileSpec) DeepCopyInto(out *ResourceProtectionProfileSpec) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceProtectionProfileSpec.
func (in *ResourceProtectionProfileSpec) DeepCopy() *ResourceProtectionProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceProtectionProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceProtectionProfileStatus) DeepCopyInto(out *ResourceProtectionProfileStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceProtectionProfileStatus.
func (in *ResourceProtectionProfileStatus) DeepCopy() *ResourceProtectionProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceProtectionProfileStatus)
	in.DeepCopyInto(out)
	return out
}

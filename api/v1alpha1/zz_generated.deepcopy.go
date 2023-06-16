//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupCPUBoost) DeepCopyInto(out *StartupCPUBoost) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Selector.DeepCopyInto(&out.Selector)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupCPUBoost.
func (in *StartupCPUBoost) DeepCopy() *StartupCPUBoost {
	if in == nil {
		return nil
	}
	out := new(StartupCPUBoost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StartupCPUBoost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupCPUBoostList) DeepCopyInto(out *StartupCPUBoostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StartupCPUBoost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupCPUBoostList.
func (in *StartupCPUBoostList) DeepCopy() *StartupCPUBoostList {
	if in == nil {
		return nil
	}
	out := new(StartupCPUBoostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StartupCPUBoostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupCPUBoostSpec) DeepCopyInto(out *StartupCPUBoostSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupCPUBoostSpec.
func (in *StartupCPUBoostSpec) DeepCopy() *StartupCPUBoostSpec {
	if in == nil {
		return nil
	}
	out := new(StartupCPUBoostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupCPUBoostStatus) DeepCopyInto(out *StartupCPUBoostStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupCPUBoostStatus.
func (in *StartupCPUBoostStatus) DeepCopy() *StartupCPUBoostStatus {
	if in == nil {
		return nil
	}
	out := new(StartupCPUBoostStatus)
	in.DeepCopyInto(out)
	return out
}
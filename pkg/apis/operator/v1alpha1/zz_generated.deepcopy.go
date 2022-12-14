//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The OpenShift Pipelines Authors

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftPipelinesConfig) DeepCopyInto(out *OpenShiftPipelinesConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftPipelinesConfig.
func (in *OpenShiftPipelinesConfig) DeepCopy() *OpenShiftPipelinesConfig {
	if in == nil {
		return nil
	}
	out := new(OpenShiftPipelinesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenShiftPipelinesConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftPipelinesConfigList) DeepCopyInto(out *OpenShiftPipelinesConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenShiftPipelinesConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftPipelinesConfigList.
func (in *OpenShiftPipelinesConfigList) DeepCopy() *OpenShiftPipelinesConfigList {
	if in == nil {
		return nil
	}
	out := new(OpenShiftPipelinesConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenShiftPipelinesConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftPipelinesConfigSpec) DeepCopyInto(out *OpenShiftPipelinesConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftPipelinesConfigSpec.
func (in *OpenShiftPipelinesConfigSpec) DeepCopy() *OpenShiftPipelinesConfigSpec {
	if in == nil {
		return nil
	}
	out := new(OpenShiftPipelinesConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftPipelinesConfigStatus) DeepCopyInto(out *OpenShiftPipelinesConfigStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftPipelinesConfigStatus.
func (in *OpenShiftPipelinesConfigStatus) DeepCopy() *OpenShiftPipelinesConfigStatus {
	if in == nil {
		return nil
	}
	out := new(OpenShiftPipelinesConfigStatus)
	in.DeepCopyInto(out)
	return out
}

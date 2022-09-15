/*
Copyright 2020 The OpenShift Pipelines Authors

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

package v1alpha1

import (
	upstreamv1alpha1 "github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

var (
// _ upstreamv1alpha1.TektonComponent = (*OpenShiftPipelinesConfig)(nil)
)

// OpenShiftPipelinesConfig is a Knative abstraction that encapsulates the interface by which Knative
// components express a desire to have a particular image cached.
//
// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
type OpenShiftPipelinesConfig struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the OpenShiftPipelinesConfig (from the client).
	// +optional
	Spec OpenShiftPipelinesConfigSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the OpenShiftPipelinesConfig (from the controller).
	// +optional
	Status OpenShiftPipelinesConfigStatus `json:"status,omitempty"`
}

var (
	// Check that AddressableService can be validated and defaulted.
	_ apis.Validatable   = (*OpenShiftPipelinesConfig)(nil)
	_ apis.Defaultable   = (*OpenShiftPipelinesConfig)(nil)
	_ kmeta.OwnerRefable = (*OpenShiftPipelinesConfig)(nil)
	// Check that the type conforms to the duck Knative Resource shape.
	_ duckv1.KRShaped = (*OpenShiftPipelinesConfig)(nil)
)

// OpenShiftPipelinesConfigSpec holds the desired state of the OpenShiftPipelinesConfig (from the client).
type OpenShiftPipelinesConfigSpec struct {
	Profile string `json:"profile,omitempty"`
	// This is where we can shape our API just the way we want
	// Like something for PipelineAsCode
}

const (
	// OpenShiftPipelinesConfigConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	OpenShiftPipelinesConfigConditionReady = apis.ConditionReady
)

// OpenShiftPipelinesConfigStatus communicates the observed state of the OpenShiftPipelinesConfig (from the controller).
type OpenShiftPipelinesConfigStatus struct {
	duckv1.Status `json:",inline"`

	ReadyReplicas int32 `json:"readyReplicas"`
}

// OpenShiftPipelinesConfigList is a list of AddressableService resources
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type OpenShiftPipelinesConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []OpenShiftPipelinesConfig `json:"items"`
}

// GetStatus retrieves the status of the resource. Implements the KRShaped interface.
func (d *OpenShiftPipelinesConfig) GetStatus() *duckv1.Status {
	return &d.Status.Status
}

// func (d *OpenShiftPipelinesConfig) GetStatus() upstreamv1alpha1.TektonComponentStatus {
// 	return &d.Status
// }

func (d *OpenShiftPipelinesConfig) GetSpec() upstreamv1alpha1.TektonComponentSpec {
	return &d.Spec
}

func (s *OpenShiftPipelinesConfigSpec) GetTargetNamespace() string {
	// FIXME: make it configurable
	return "openshift-pipelines"
}
func (s *OpenShiftPipelinesConfigStatus) GetManifests() []string {
	return []string{}
}

func (s *OpenShiftPipelinesConfigStatus) GetVersion() string {
	// FIXME: do not hardcode this
	return "0.0.1"
}
func (s *OpenShiftPipelinesConfigStatus) SetVersion(string) {
}

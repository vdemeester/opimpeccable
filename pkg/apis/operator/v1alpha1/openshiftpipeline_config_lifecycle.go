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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"
)

const (
	PreInstall      apis.ConditionType = "PreInstall"
	ComponentsReady apis.ConditionType = "ComponentsReady"
	PostInstall     apis.ConditionType = "PostInstall"
)

var (
	configCondSet = apis.NewLivingConditionSet(
		PreInstall,
		ComponentsReady,
		PostInstall,
	)
)

// GetGroupVersionKind implements kmeta.OwnerRefable
func (*OpenShiftPipelinesConfig) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("OpenShiftPipelinesConfig")
}

// GetConditionSet retrieves the condition set for this resource. Implements the KRShaped interface.
func (d *OpenShiftPipelinesConfig) GetConditionSet() apis.ConditionSet {
	return configCondSet
}

// InitializeConditions sets the initial values to the conditions.
func (ospcs *OpenShiftPipelinesConfigStatus) InitializeConditions() {
	configCondSet.Manage(ospcs).InitializeConditions()
}

func (ospcs *OpenShiftPipelinesConfigStatus) MarkNotReady(msg string) {
	configCondSet.Manage(ospcs).MarkFalse(
		apis.ConditionReady,
		"Error",
		"Ready: %s", msg)
}

func (ospcs *OpenShiftPipelinesConfigStatus) IsReady() bool {
	return configCondSet.Manage(ospcs).IsHappy()
}

// To implement (or not)
func (ospcs *OpenShiftPipelinesConfigStatus) MarkInstallSucceeded() {
}
func (ospcs *OpenShiftPipelinesConfigStatus) MarkInstallFailed(msg string) {
}
func (ospcs *OpenShiftPipelinesConfigStatus) MarkDeploymentsAvailable() {
}
func (ospcs *OpenShiftPipelinesConfigStatus) MarkDeploymentsNotReady() {
}
func (ospcs *OpenShiftPipelinesConfigStatus) MarkDependenciesInstalled() {
}
func (ospcs *OpenShiftPipelinesConfigStatus) MarkDependencyInstalling(msg string) {
}
func (ospcs *OpenShiftPipelinesConfigStatus) MarkDependencyMissing(msg string) {
}

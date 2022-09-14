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

package opconfig

import (
	"context"
	"fmt"

	// corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	// corev1listers "k8s.io/client-go/listers/core/v1"

	// "github.com/vdemeester/opimpeccable/pkg/apis/operator"
	operatorv1alpha1 "github.com/vdemeester/opimpeccable/pkg/apis/operator/v1alpha1"
	clientset "github.com/vdemeester/opimpeccable/pkg/client/clientset/versioned"
	openshiftpipelinesconfigreconciler "github.com/vdemeester/opimpeccable/pkg/client/injection/reconciler/operator/v1alpha1/openshiftpipelinesconfig"
	// "knative.dev/pkg/kmeta"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"
)

// Reconciler implements openshiftpipelinesconfigreconciler.Interface for
// OpenShiftPipelinesConfig resources.
type Reconciler struct {
	kubeclient     kubernetes.Interface
	operatorclient clientset.Interface
}

// Check that our Reconciler implements Interface
var _ openshiftpipelinesconfigreconciler.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, c *operatorv1alpha1.OpenShiftPipelinesConfig) reconciler.Event {
	// This logger has all the context necessary to identify which resource is being reconciled.
	logger := logging.FromContext(ctx)
	logger.Infof("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")

	// As of today, the job is relatively simple:
	// From an OpenShiftPipelineConfig, create or update a TektonConfig (upstream)
	// and more in the future.
	c.Status.InitializeConditions()
	logger.Infow("Reconciling OpenShiftPipelineConfig", "status", c.Status)

	c.SetDefaults(ctx)
	if c.GetName() != "config" {
		msg := fmt.Sprintf("Resource ignored, Expected Name: %s, Got Name: %s",
			"config",
			c.GetName(),
		)
		logger.Error(msg)
		c.Status.MarkNotReady(msg)
		return nil
	}

	return nil
}

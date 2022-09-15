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

	upstreamv1alpha1 "github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	// corev1listers "k8s.io/client-go/listers/core/v1"

	// "github.com/vdemeester/opimpeccable/pkg/apis/operator"
	upstreamclientset "github.com/tektoncd/operator/pkg/client/clientset/versioned"
	operatorv1alpha1 "github.com/vdemeester/opimpeccable/pkg/apis/operator/v1alpha1"
	clientset "github.com/vdemeester/opimpeccable/pkg/client/clientset/versioned"
	openshiftpipelinesconfigreconciler "github.com/vdemeester/opimpeccable/pkg/client/injection/reconciler/operator/v1alpha1/openshiftpipelinesconfig"
	// "knative.dev/pkg/kmeta"
	"github.com/tektoncd/operator/pkg/reconciler/shared/tektonconfig/pipeline"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"
)

// Reconciler implements openshiftpipelinesconfigreconciler.Interface for
// OpenShiftPipelinesConfig resources.
type Reconciler struct {
	kubeclient           kubernetes.Interface
	operatorclient       clientset.Interface
	tektonoperatorclient upstreamclientset.Interface
}

// Check that our Reconciler implements Interface
var _ openshiftpipelinesconfigreconciler.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, c *operatorv1alpha1.OpenShiftPipelinesConfig) reconciler.Event {
	// This logger has all the context necessary to identify which resource is being reconciled.
	logger := logging.FromContext(ctx)

	logger.Infof("ObjectMeta: %+v", c.ObjectMeta)
	logger.Infof("TypeMeta: %+v", c.TypeMeta)

	// Ugly ass hack
	c.TypeMeta = metav1.TypeMeta{
		APIVersion: operatorv1alpha1.SchemeGroupVersion.String(),
		Kind:       "OpenShiftPipelineConfig",
	}
	logger.Infof("TypeMeta: %+v", c.TypeMeta)

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

	if c.Spec.Profile != "lite" {
		msg := fmt.Sprintf("Only lite profile is supported today, %s is not supported", c.Spec.Profile)
		logger.Error(msg)
		c.Status.MarkNotReady(msg)
		return nil
	}

	// Create a simple TektonPipeline
	// FIXME: ideally EnsureTektonPipelineExists doesn't depend on TektonConfig
	tc := &upstreamv1alpha1.TektonConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: "config",
		},
		Spec: upstreamv1alpha1.TektonConfigSpec{
			Profile: c.Spec.Profile,
			CommonSpec: upstreamv1alpha1.CommonSpec{
				TargetNamespace: "openshift-pipelines",
			},
		},
	}
	if err := r.ensureTargetNamespaceExists(ctx, tc, c); err != nil {
		return err
	}
	// Ensure if the pipeline CR already exists, if not create Pipeline CR
	if _, err := pipeline.EnsureTektonPipelineExists(ctx, r.tektonoperatorclient.OperatorV1alpha1().TektonPipelines(), tc, c); err != nil {
		c.Status.MarkNotReady(fmt.Sprintf("TektonPipeline: %s", err.Error()))
		if err == upstreamv1alpha1.RECONCILE_AGAIN_ERR {
			return upstreamv1alpha1.REQUEUE_EVENT_AFTER
		}
		return nil
	}
	return nil
}
func (r *Reconciler) ensureTargetNamespaceExists(ctx context.Context, tc *upstreamv1alpha1.TektonConfig, c *operatorv1alpha1.OpenShiftPipelinesConfig) error {

	ns, err := r.kubeclient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("operator.tekton.dev/targetNamespace=%s", "true"),
	})

	if err != nil {
		return err
	}

	if len(ns.Items) > 0 {
		for _, namespace := range ns.Items {
			if namespace.Name != tc.GetSpec().GetTargetNamespace() {
				namespace.Labels["operator.tekton.dev/mark-for-deletion"] = "true"
				_, err = r.kubeclient.CoreV1().Namespaces().Update(ctx, &namespace, metav1.UpdateOptions{})
				if err != nil {
					return err
				}

			} else {
				return nil
			}
		}
	} else {
		if err := createTargetNamespace(ctx, nil, tc, c, r.kubeclient); err != nil {
			if errors.IsAlreadyExists(err) {
				return r.addTargetNamespaceLabel(ctx, tc.GetSpec().GetTargetNamespace())
			}
			return err
		}
	}
	return nil
}

func createTargetNamespace(ctx context.Context, labels map[string]string, obj upstreamv1alpha1.TektonComponent, ownerObj *operatorv1alpha1.OpenShiftPipelinesConfig, kubeClientSet kubernetes.Interface) error {
	// ownerRef := *metav1.NewControllerRef(ownerObj, ownerObj.GroupVersionKind())
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: obj.GetSpec().GetTargetNamespace(),
			Labels: map[string]string{
				"operator.tekton.dev/targetNamespace": "true",
			},
			// OwnerReferences: []metav1.OwnerReference{ownerRef},
		},
	}

	if len(labels) > 0 {
		for key, value := range labels {
			namespace.Labels[key] = value
		}
	}

	if _, err := kubeClientSet.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}

func (r *Reconciler) addTargetNamespaceLabel(ctx context.Context, targetNamespace string) error {
	ns, err := r.kubeclient.CoreV1().Namespaces().Get(ctx, targetNamespace, v1.GetOptions{})
	if err != nil {
		return err
	}
	labels := ns.GetLabels()
	if labels == nil {
		labels = map[string]string{
			"operator.tekton.dev/targetNamespace": "true",
		}
	}
	ns.SetLabels(labels)
	_, err = r.kubeclient.CoreV1().Namespaces().Update(ctx, ns, v1.UpdateOptions{})
	return err
}

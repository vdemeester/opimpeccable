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
	"regexp"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/kmeta"
	"knative.dev/pkg/logging"

	// "github.com/vdemeester/opimpeccable/pkg/apis/operator/v1alpha1"
	operatorclient "github.com/vdemeester/opimpeccable/pkg/client/injection/client"
	openshiftpipelinesconfiginformer "github.com/vdemeester/opimpeccable/pkg/client/injection/informers/operator/v1alpha1/openshiftpipelinesconfig"
	openshiftpipelinesconfigreconciler "github.com/vdemeester/opimpeccable/pkg/client/injection/reconciler/operator/v1alpha1/openshiftpipelinesconfig"
	"k8s.io/apimachinery/pkg/types"
	kubeclient "knative.dev/pkg/client/injection/kube/client"
	namespaceinformer "knative.dev/pkg/client/injection/kube/informers/core/v1/namespace"
)

// NewController creates a Reconciler and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)
	logger.Infof("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	logger.Infof("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	logger.Infof("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	logger.Infof("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	logger.Infof("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	logger.Infof("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	// Obtain an informer to both the main and child resources. These will be started by
	// the injection framework automatically. They'll keep a cached representation of the
	// cluster's state of the respective resource at all times.
	openshiftpipelinesconfigInformer := openshiftpipelinesconfiginformer.Get(ctx)

	r := &Reconciler{
		// The client will be needed to create/delete Pods via the API.
		kubeclient:     kubeclient.Get(ctx),
		operatorclient: operatorclient.Get(ctx),
	}
	impl := openshiftpipelinesconfigreconciler.NewImpl(ctx, r, func(impl *controller.Impl) controller.Options {
		return controller.Options{
			AgentName: "OpenShiftPipelinesConfig",
		}
	})

	// Listen for events on the main resource and enqueue themselves.
	openshiftpipelinesconfigInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))
	namespaceinformer.Get(ctx).Informer().AddEventHandler(controller.HandleAll(enqueueCustomName(impl, "config")))

	return impl
}

// enqueueCustomName adds an event with name `config` in work queue so that
// whenever a namespace event occurs, the TektonConfig reconciler get triggered.
// This is required because we want to get our TektonConfig reconciler triggered
// for already existing and new namespaces, without manual intervention like adding
// a label/annotation on namespace to make it manageable by Tekton controller.
// This will also filter the namespaces by regex `^(openshift|kube)-`
// and enqueue only when namespace doesn't match the regex
func enqueueCustomName(impl *controller.Impl, name string) func(obj interface{}) {
	return func(obj interface{}) {
		var nsRegex = regexp.MustCompile("^(openshift|kube)-")
		object, err := kmeta.DeletionHandlingAccessor(obj)
		if err == nil && !nsRegex.MatchString(object.GetName()) {
			impl.EnqueueKey(types.NamespacedName{Namespace: "", Name: name})
		}
	}
}

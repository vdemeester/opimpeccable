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

// Code generated by injection-gen. DO NOT EDIT.

package openshiftpipelinesconfig

import (
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1alpha1 "github.com/vdemeester/opimpeccable/pkg/apis/operator/v1alpha1"
	versioned "github.com/vdemeester/opimpeccable/pkg/client/clientset/versioned"
	operatorv1alpha1 "github.com/vdemeester/opimpeccable/pkg/client/listers/operator/v1alpha1"
	zap "go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	equality "k8s.io/apimachinery/pkg/api/equality"
	errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	sets "k8s.io/apimachinery/pkg/util/sets"
	record "k8s.io/client-go/tools/record"
	controller "knative.dev/pkg/controller"
	kmp "knative.dev/pkg/kmp"
	logging "knative.dev/pkg/logging"
	reconciler "knative.dev/pkg/reconciler"
)

// Interface defines the strongly typed interfaces to be implemented by a
// controller reconciling v1alpha1.OpenShiftPipelinesConfig.
type Interface interface {
	// ReconcileKind implements custom logic to reconcile v1alpha1.OpenShiftPipelinesConfig. Any changes
	// to the objects .Status or .Finalizers will be propagated to the stored
	// object. It is recommended that implementors do not call any update calls
	// for the Kind inside of ReconcileKind, it is the responsibility of the calling
	// controller to propagate those properties. The resource passed to ReconcileKind
	// will always have an empty deletion timestamp.
	ReconcileKind(ctx context.Context, o *v1alpha1.OpenShiftPipelinesConfig) reconciler.Event
}

// Finalizer defines the strongly typed interfaces to be implemented by a
// controller finalizing v1alpha1.OpenShiftPipelinesConfig.
type Finalizer interface {
	// FinalizeKind implements custom logic to finalize v1alpha1.OpenShiftPipelinesConfig. Any changes
	// to the objects .Status or .Finalizers will be ignored. Returning a nil or
	// Normal type reconciler.Event will allow the finalizer to be deleted on
	// the resource. The resource passed to FinalizeKind will always have a set
	// deletion timestamp.
	FinalizeKind(ctx context.Context, o *v1alpha1.OpenShiftPipelinesConfig) reconciler.Event
}

// ReadOnlyInterface defines the strongly typed interfaces to be implemented by a
// controller reconciling v1alpha1.OpenShiftPipelinesConfig if they want to process resources for which
// they are not the leader.
type ReadOnlyInterface interface {
	// ObserveKind implements logic to observe v1alpha1.OpenShiftPipelinesConfig.
	// This method should not write to the API.
	ObserveKind(ctx context.Context, o *v1alpha1.OpenShiftPipelinesConfig) reconciler.Event
}

type doReconcile func(ctx context.Context, o *v1alpha1.OpenShiftPipelinesConfig) reconciler.Event

// reconcilerImpl implements controller.Reconciler for v1alpha1.OpenShiftPipelinesConfig resources.
type reconcilerImpl struct {
	// LeaderAwareFuncs is inlined to help us implement reconciler.LeaderAware.
	reconciler.LeaderAwareFuncs

	// Client is used to write back status updates.
	Client versioned.Interface

	// Listers index properties about resources.
	Lister operatorv1alpha1.OpenShiftPipelinesConfigLister

	// Recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	Recorder record.EventRecorder

	// configStore allows for decorating a context with config maps.
	// +optional
	configStore reconciler.ConfigStore

	// reconciler is the implementation of the business logic of the resource.
	reconciler Interface

	// finalizerName is the name of the finalizer to reconcile.
	finalizerName string

	// skipStatusUpdates configures whether or not this reconciler automatically updates
	// the status of the reconciled resource.
	skipStatusUpdates bool
}

// Check that our Reconciler implements controller.Reconciler.
var _ controller.Reconciler = (*reconcilerImpl)(nil)

// Check that our generated Reconciler is always LeaderAware.
var _ reconciler.LeaderAware = (*reconcilerImpl)(nil)

func NewReconciler(ctx context.Context, logger *zap.SugaredLogger, client versioned.Interface, lister operatorv1alpha1.OpenShiftPipelinesConfigLister, recorder record.EventRecorder, r Interface, options ...controller.Options) controller.Reconciler {
	// Check the options function input. It should be 0 or 1.
	if len(options) > 1 {
		logger.Fatal("Up to one options struct is supported, found: ", len(options))
	}

	// Fail fast when users inadvertently implement the other LeaderAware interface.
	// For the typed reconcilers, Promote shouldn't take any arguments.
	if _, ok := r.(reconciler.LeaderAware); ok {
		logger.Fatalf("%T implements the incorrect LeaderAware interface. Promote() should not take an argument as genreconciler handles the enqueuing automatically.", r)
	}

	rec := &reconcilerImpl{
		LeaderAwareFuncs: reconciler.LeaderAwareFuncs{
			PromoteFunc: func(bkt reconciler.Bucket, enq func(reconciler.Bucket, types.NamespacedName)) error {
				all, err := lister.List(labels.Everything())
				if err != nil {
					return err
				}
				for _, elt := range all {
					// TODO: Consider letting users specify a filter in options.
					enq(bkt, types.NamespacedName{
						Namespace: elt.GetNamespace(),
						Name:      elt.GetName(),
					})
				}
				return nil
			},
		},
		Client:        client,
		Lister:        lister,
		Recorder:      recorder,
		reconciler:    r,
		finalizerName: defaultFinalizerName,
	}

	for _, opts := range options {
		if opts.ConfigStore != nil {
			rec.configStore = opts.ConfigStore
		}
		if opts.FinalizerName != "" {
			rec.finalizerName = opts.FinalizerName
		}
		if opts.SkipStatusUpdates {
			rec.skipStatusUpdates = true
		}
		if opts.DemoteFunc != nil {
			rec.DemoteFunc = opts.DemoteFunc
		}
	}

	return rec
}

// Reconcile implements controller.Reconciler
func (r *reconcilerImpl) Reconcile(ctx context.Context, key string) error {
	logger := logging.FromContext(ctx)

	// Initialize the reconciler state. This will convert the namespace/name
	// string into a distinct namespace and name, determine if this instance of
	// the reconciler is the leader, and any additional interfaces implemented
	// by the reconciler. Returns an error is the resource key is invalid.
	s, err := newState(key, r)
	if err != nil {
		logger.Error("Invalid resource key: ", key)
		return nil
	}

	// If we are not the leader, and we don't implement either ReadOnly
	// observer interfaces, then take a fast-path out.
	if s.isNotLeaderNorObserver() {
		return controller.NewSkipKey(key)
	}

	// If configStore is set, attach the frozen configuration to the context.
	if r.configStore != nil {
		ctx = r.configStore.ToContext(ctx)
	}

	// Add the recorder to context.
	ctx = controller.WithEventRecorder(ctx, r.Recorder)

	// Get the resource with this namespace/name.

	getter := r.Lister.OpenShiftPipelinesConfigs(s.namespace)

	original, err := getter.Get(s.name)

	if errors.IsNotFound(err) {
		// The resource may no longer exist, in which case we stop processing and call
		// the ObserveDeletion handler if appropriate.
		logger.Debugf("Resource %q no longer exists", key)
		if del, ok := r.reconciler.(reconciler.OnDeletionInterface); ok {
			return del.ObserveDeletion(ctx, types.NamespacedName{
				Namespace: s.namespace,
				Name:      s.name,
			})
		}
		return nil
	} else if err != nil {
		return err
	}

	// Don't modify the informers copy.
	resource := original.DeepCopy()

	var reconcileEvent reconciler.Event

	name, do := s.reconcileMethodFor(resource)
	// Append the target method to the logger.
	logger = logger.With(zap.String("targetMethod", name))
	switch name {
	case reconciler.DoReconcileKind:
		// Set and update the finalizer on resource if r.reconciler
		// implements Finalizer.
		if resource, err = r.setFinalizerIfFinalizer(ctx, resource); err != nil {
			return fmt.Errorf("failed to set finalizers: %w", err)
		}

		if !r.skipStatusUpdates {
			reconciler.PreProcessReconcile(ctx, resource)
		}

		// Reconcile this copy of the resource and then write back any status
		// updates regardless of whether the reconciliation errored out.
		reconcileEvent = do(ctx, resource)

		if !r.skipStatusUpdates {
			reconciler.PostProcessReconcile(ctx, resource, original)
		}

	case reconciler.DoFinalizeKind:
		// For finalizing reconcilers, if this resource being marked for deletion
		// and reconciled cleanly (nil or normal event), remove the finalizer.
		reconcileEvent = do(ctx, resource)

		if resource, err = r.clearFinalizer(ctx, resource, reconcileEvent); err != nil {
			return fmt.Errorf("failed to clear finalizers: %w", err)
		}

	case reconciler.DoObserveKind:
		// Observe any changes to this resource, since we are not the leader.
		reconcileEvent = do(ctx, resource)

	}

	// Synchronize the status.
	switch {
	case r.skipStatusUpdates:
		// This reconciler implementation is configured to skip resource updates.
		// This may mean this reconciler does not observe spec, but reconciles external changes.
	case equality.Semantic.DeepEqual(original.Status, resource.Status):
		// If we didn't change anything then don't call updateStatus.
		// This is important because the copy we loaded from the injectionInformer's
		// cache may be stale and we don't want to overwrite a prior update
		// to status with this stale state.
	case !s.isLeader:
		// High-availability reconcilers may have many replicas watching the resource, but only
		// the elected leader is expected to write modifications.
		logger.Warn("Saw status changes when we aren't the leader!")
	default:
		if err = r.updateStatus(ctx, original, resource); err != nil {
			logger.Warnw("Failed to update resource status", zap.Error(err))
			r.Recorder.Eventf(resource, v1.EventTypeWarning, "UpdateFailed",
				"Failed to update status for %q: %v", resource.Name, err)
			return err
		}
	}

	// Report the reconciler event, if any.
	if reconcileEvent != nil {
		var event *reconciler.ReconcilerEvent
		if reconciler.EventAs(reconcileEvent, &event) {
			logger.Infow("Returned an event", zap.Any("event", reconcileEvent))
			r.Recorder.Event(resource, event.EventType, event.Reason, event.Error())

			// the event was wrapped inside an error, consider the reconciliation as failed
			if _, isEvent := reconcileEvent.(*reconciler.ReconcilerEvent); !isEvent {
				return reconcileEvent
			}
			return nil
		}

		if controller.IsSkipKey(reconcileEvent) {
			// This is a wrapped error, don't emit an event.
		} else if ok, _ := controller.IsRequeueKey(reconcileEvent); ok {
			// This is a wrapped error, don't emit an event.
		} else {
			logger.Errorw("Returned an error", zap.Error(reconcileEvent))
			r.Recorder.Event(resource, v1.EventTypeWarning, "InternalError", reconcileEvent.Error())
		}
		return reconcileEvent
	}

	return nil
}

func (r *reconcilerImpl) updateStatus(ctx context.Context, existing *v1alpha1.OpenShiftPipelinesConfig, desired *v1alpha1.OpenShiftPipelinesConfig) error {
	existing = existing.DeepCopy()
	return reconciler.RetryUpdateConflicts(func(attempts int) (err error) {
		// The first iteration tries to use the injectionInformer's state, subsequent attempts fetch the latest state via API.
		if attempts > 0 {

			getter := r.Client.OperatorV1alpha1().OpenShiftPipelinesConfigs(desired.Namespace)

			existing, err = getter.Get(ctx, desired.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
		}

		// If there's nothing to update, just return.
		if equality.Semantic.DeepEqual(existing.Status, desired.Status) {
			return nil
		}

		if diff, err := kmp.SafeDiff(existing.Status, desired.Status); err == nil && diff != "" {
			logging.FromContext(ctx).Debug("Updating status with: ", diff)
		}

		existing.Status = desired.Status

		updater := r.Client.OperatorV1alpha1().OpenShiftPipelinesConfigs(existing.Namespace)

		_, err = updater.UpdateStatus(ctx, existing, metav1.UpdateOptions{})
		return err
	})
}

// updateFinalizersFiltered will update the Finalizers of the resource.
// TODO: this method could be generic and sync all finalizers. For now it only
// updates defaultFinalizerName or its override.
func (r *reconcilerImpl) updateFinalizersFiltered(ctx context.Context, resource *v1alpha1.OpenShiftPipelinesConfig) (*v1alpha1.OpenShiftPipelinesConfig, error) {

	getter := r.Lister.OpenShiftPipelinesConfigs(resource.Namespace)

	actual, err := getter.Get(resource.Name)
	if err != nil {
		return resource, err
	}

	// Don't modify the informers copy.
	existing := actual.DeepCopy()

	var finalizers []string

	// If there's nothing to update, just return.
	existingFinalizers := sets.NewString(existing.Finalizers...)
	desiredFinalizers := sets.NewString(resource.Finalizers...)

	if desiredFinalizers.Has(r.finalizerName) {
		if existingFinalizers.Has(r.finalizerName) {
			// Nothing to do.
			return resource, nil
		}
		// Add the finalizer.
		finalizers = append(existing.Finalizers, r.finalizerName)
	} else {
		if !existingFinalizers.Has(r.finalizerName) {
			// Nothing to do.
			return resource, nil
		}
		// Remove the finalizer.
		existingFinalizers.Delete(r.finalizerName)
		finalizers = existingFinalizers.List()
	}

	mergePatch := map[string]interface{}{
		"metadata": map[string]interface{}{
			"finalizers":      finalizers,
			"resourceVersion": existing.ResourceVersion,
		},
	}

	patch, err := json.Marshal(mergePatch)
	if err != nil {
		return resource, err
	}

	patcher := r.Client.OperatorV1alpha1().OpenShiftPipelinesConfigs(resource.Namespace)

	resourceName := resource.Name
	updated, err := patcher.Patch(ctx, resourceName, types.MergePatchType, patch, metav1.PatchOptions{})
	if err != nil {
		r.Recorder.Eventf(existing, v1.EventTypeWarning, "FinalizerUpdateFailed",
			"Failed to update finalizers for %q: %v", resourceName, err)
	} else {
		r.Recorder.Eventf(updated, v1.EventTypeNormal, "FinalizerUpdate",
			"Updated %q finalizers", resource.GetName())
	}
	return updated, err
}

func (r *reconcilerImpl) setFinalizerIfFinalizer(ctx context.Context, resource *v1alpha1.OpenShiftPipelinesConfig) (*v1alpha1.OpenShiftPipelinesConfig, error) {
	if _, ok := r.reconciler.(Finalizer); !ok {
		return resource, nil
	}

	finalizers := sets.NewString(resource.Finalizers...)

	// If this resource is not being deleted, mark the finalizer.
	if resource.GetDeletionTimestamp().IsZero() {
		finalizers.Insert(r.finalizerName)
	}

	resource.Finalizers = finalizers.List()

	// Synchronize the finalizers filtered by r.finalizerName.
	return r.updateFinalizersFiltered(ctx, resource)
}

func (r *reconcilerImpl) clearFinalizer(ctx context.Context, resource *v1alpha1.OpenShiftPipelinesConfig, reconcileEvent reconciler.Event) (*v1alpha1.OpenShiftPipelinesConfig, error) {
	if _, ok := r.reconciler.(Finalizer); !ok {
		return resource, nil
	}
	if resource.GetDeletionTimestamp().IsZero() {
		return resource, nil
	}

	finalizers := sets.NewString(resource.Finalizers...)

	if reconcileEvent != nil {
		var event *reconciler.ReconcilerEvent
		if reconciler.EventAs(reconcileEvent, &event) {
			if event.EventType == v1.EventTypeNormal {
				finalizers.Delete(r.finalizerName)
			}
		}
	} else {
		finalizers.Delete(r.finalizerName)
	}

	resource.Finalizers = finalizers.List()

	// Synchronize the finalizers filtered by r.finalizerName.
	return r.updateFinalizersFiltered(ctx, resource)
}

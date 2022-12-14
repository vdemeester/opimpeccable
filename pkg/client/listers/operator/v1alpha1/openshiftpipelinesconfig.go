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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vdemeester/opimpeccable/pkg/apis/operator/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OpenShiftPipelinesConfigLister helps list OpenShiftPipelinesConfigs.
// All objects returned here must be treated as read-only.
type OpenShiftPipelinesConfigLister interface {
	// List lists all OpenShiftPipelinesConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OpenShiftPipelinesConfig, err error)
	// Get retrieves the OpenShiftPipelinesConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OpenShiftPipelinesConfig, error)
	OpenShiftPipelinesConfigListerExpansion
}

// openShiftPipelinesConfigLister implements the OpenShiftPipelinesConfigLister interface.
type openShiftPipelinesConfigLister struct {
	indexer cache.Indexer
}

// NewOpenShiftPipelinesConfigLister returns a new OpenShiftPipelinesConfigLister.
func NewOpenShiftPipelinesConfigLister(indexer cache.Indexer) OpenShiftPipelinesConfigLister {
	return &openShiftPipelinesConfigLister{indexer: indexer}
}

// List lists all OpenShiftPipelinesConfigs in the indexer.
func (s *openShiftPipelinesConfigLister) List(selector labels.Selector) (ret []*v1alpha1.OpenShiftPipelinesConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpenShiftPipelinesConfig))
	})
	return ret, err
}

// Get retrieves the OpenShiftPipelinesConfig from the index for a given name.
func (s *openShiftPipelinesConfigLister) Get(name string) (*v1alpha1.OpenShiftPipelinesConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("openshiftpipelinesconfig"), name)
	}
	return obj.(*v1alpha1.OpenShiftPipelinesConfig), nil
}

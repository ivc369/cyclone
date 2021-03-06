/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceLister helps list Resources.
type ResourceLister interface {
	// List lists all Resources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Resource, err error)
	// Resources returns an object that can list and get Resources.
	Resources(namespace string) ResourceNamespaceLister
	ResourceListerExpansion
}

// resourceLister implements the ResourceLister interface.
type resourceLister struct {
	indexer cache.Indexer
}

// NewResourceLister returns a new ResourceLister.
func NewResourceLister(indexer cache.Indexer) ResourceLister {
	return &resourceLister{indexer: indexer}
}

// List lists all Resources in the indexer.
func (s *resourceLister) List(selector labels.Selector) (ret []*v1alpha1.Resource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Resource))
	})
	return ret, err
}

// Resources returns an object that can list and get Resources.
func (s *resourceLister) Resources(namespace string) ResourceNamespaceLister {
	return resourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceNamespaceLister helps list and get Resources.
type ResourceNamespaceLister interface {
	// List lists all Resources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Resource, err error)
	// Get retrieves the Resource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Resource, error)
	ResourceNamespaceListerExpansion
}

// resourceNamespaceLister implements the ResourceNamespaceLister
// interface.
type resourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Resources in the indexer for a given namespace.
func (s resourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Resource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Resource))
	})
	return ret, err
}

// Get retrieves the Resource from the indexer for a given namespace and name.
func (s resourceNamespaceLister) Get(name string) (*v1alpha1.Resource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.GroupResource("resource"), name)
	}
	return obj.(*v1alpha1.Resource), nil
}

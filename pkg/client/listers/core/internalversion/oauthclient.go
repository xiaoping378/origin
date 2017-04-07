// This file was automatically generated by lister-gen with arguments: --input-dirs=[github.com/openshift/origin/pkg/authorization/api,github.com/openshift/origin/pkg/authorization/api/v1,github.com/openshift/origin/pkg/build/api,github.com/openshift/origin/pkg/build/api/v1,github.com/openshift/origin/pkg/deploy/api,github.com/openshift/origin/pkg/deploy/api/v1,github.com/openshift/origin/pkg/image/api,github.com/openshift/origin/pkg/image/api/v1,github.com/openshift/origin/pkg/oauth/api,github.com/openshift/origin/pkg/oauth/api/v1,github.com/openshift/origin/pkg/project/api,github.com/openshift/origin/pkg/project/api/v1,github.com/openshift/origin/pkg/quota/api,github.com/openshift/origin/pkg/quota/api/v1,github.com/openshift/origin/pkg/route/api,github.com/openshift/origin/pkg/route/api/v1,github.com/openshift/origin/pkg/sdn/api,github.com/openshift/origin/pkg/sdn/api/v1,github.com/openshift/origin/pkg/template/api,github.com/openshift/origin/pkg/template/api/v1,github.com/openshift/origin/pkg/user/api,github.com/openshift/origin/pkg/user/api/v1] --logtostderr=true

package internalversion

import (
	api "github.com/openshift/origin/pkg/oauth/api"
	"k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// OAuthClientLister helps list OAuthClients.
type OAuthClientLister interface {
	// List lists all OAuthClients in the indexer.
	List(selector labels.Selector) (ret []*api.OAuthClient, err error)
	// OAuthClients returns an object that can list and get OAuthClients.
	OAuthClients(namespace string) OAuthClientNamespaceLister
	OAuthClientListerExpansion
}

// oAuthClientLister implements the OAuthClientLister interface.
type oAuthClientLister struct {
	indexer cache.Indexer
}

// NewOAuthClientLister returns a new OAuthClientLister.
func NewOAuthClientLister(indexer cache.Indexer) OAuthClientLister {
	return &oAuthClientLister{indexer: indexer}
}

// List lists all OAuthClients in the indexer.
func (s *oAuthClientLister) List(selector labels.Selector) (ret []*api.OAuthClient, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*api.OAuthClient))
	})
	return ret, err
}

// OAuthClients returns an object that can list and get OAuthClients.
func (s *oAuthClientLister) OAuthClients(namespace string) OAuthClientNamespaceLister {
	return oAuthClientNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OAuthClientNamespaceLister helps list and get OAuthClients.
type OAuthClientNamespaceLister interface {
	// List lists all OAuthClients in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*api.OAuthClient, err error)
	// Get retrieves the OAuthClient from the indexer for a given namespace and name.
	Get(name string) (*api.OAuthClient, error)
	OAuthClientNamespaceListerExpansion
}

// oAuthClientNamespaceLister implements the OAuthClientNamespaceLister
// interface.
type oAuthClientNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OAuthClients in the indexer for a given namespace.
func (s oAuthClientNamespaceLister) List(selector labels.Selector) (ret []*api.OAuthClient, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*api.OAuthClient))
	})
	return ret, err
}

// Get retrieves the OAuthClient from the indexer for a given namespace and name.
func (s oAuthClientNamespaceLister) Get(name string) (*api.OAuthClient, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("oauthclient"), name)
	}
	return obj.(*api.OAuthClient), nil
}

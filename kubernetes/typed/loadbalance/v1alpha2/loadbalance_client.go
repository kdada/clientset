/*
Copyright 2018 caicloud authors. All rights reserved.
*/

package v1alpha2

import (
	"github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha2 "github.com/caicloud/clientset/pkg/apis/loadbalance/v1alpha2"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type LoadbalanceV1alpha2Interface interface {
	RESTClient() rest.Interface
	LoadBalancersGetter
}

// LoadbalanceV1alpha2Client is used to interact with features provided by the loadbalance.caicloud.io group.
type LoadbalanceV1alpha2Client struct {
	restClient rest.Interface
}

func (c *LoadbalanceV1alpha2Client) LoadBalancers(namespace string) LoadBalancerInterface {
	return newLoadBalancers(c, namespace)
}

// NewForConfig creates a new LoadbalanceV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*LoadbalanceV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &LoadbalanceV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new LoadbalanceV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *LoadbalanceV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new LoadbalanceV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *LoadbalanceV1alpha2Client {
	return &LoadbalanceV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *LoadbalanceV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

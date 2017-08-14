/*
Copyright 2017 caicloud authors. All rights reserved.
*/

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GrayReleases returns a GrayReleaseInformer.
	GrayReleases() GrayReleaseInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// GrayReleases returns a GrayReleaseInformer.
func (v *version) GrayReleases() GrayReleaseInformer {
	return &grayReleaseInformer{factory: v.SharedInformerFactory}
}

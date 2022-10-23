package fn

import (
	rctxv1 "github.com/yndd/lcnc-runtime/pkg/api/resourcecontext/v1"
)

// ResourceContextProcessor is implemented by configuration functions built with this framework
// to conform to the Configuration Functions Specification:
// https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/docs/api-conventions/functions-spec.md
type ResourceContextProcessor interface {
	Process(rctx *rctxv1.ResourceContext) (bool, error)
}

// ResourceContextProcessorFunc converts a compatible function to a ResourceContextProcessor.
type ResourceContextProcessorFunc func(rctx *rctxv1.ResourceContext) (bool, error)

func (p ResourceContextProcessorFunc) Process(rctx *rctxv1.ResourceContext) (bool, error) {
	return p(rctx)
}
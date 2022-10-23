package fn

import (
	rctxv1 "github.com/yndd/lcnc-runtime/pkg/api/resourcecontext/v1"
)

type Runner interface {
	// Returns:
	//    return a boolean to tell whether the execution should be considered as PASS or FAIL.
	Run(rctx *rctxv1.ResourceContext) bool
}

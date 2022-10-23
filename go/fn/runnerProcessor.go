package fn

import (
	rctxv1 "github.com/yndd/lcnc-runtime/pkg/api/resourcecontext/v1"
)

type runnerProcessor struct {
	//ctx      context.Context
	fnRunner Runner
}

func (r runnerProcessor) Process(rctx *rctxv1.ResourceContext) (bool, error) {

	// Run the main function.
	//fnCtx := &Context{Context: r.ctx}
	//results := new(Results)
	shouldPass := r.fnRunner.Run(rctx)
	// If running in a pipeline, the ResourceList may already have results from previous function runs.
	// Thus, we only append new results to the end.
	//rl.Results = append(rl.Results, *results...)
	return shouldPass, nil
}

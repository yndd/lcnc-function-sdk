package fn

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
)

func WithContext(ctx context.Context, runner Runner) ResourceContextProcessor {
	return runnerProcessor{ctx: ctx, fnRunner: runner}
}

type runnerProcessor struct {
	ctx      context.Context
	fnRunner Runner
}

func (r runnerProcessor) Process(rCtx *ResourceContext) (bool, error) {
	// TBD if we need to process the function config or not
	if rCtx.Results == nil {
		rCtx.Results = &Results{}
	}
	if rCtx.Resources == nil {
		rCtx.Results.Errorf("expecting some resource input, got none")
		return false, nil
	}

	// Run the main function.
	fnCtx := &Context{Context: r.ctx}
	// validate and initialize the output and the conditions

	// initialize the result
	results := &Results{}
	// initialize the conditions and output
	rCtx.Resources.Output = map[string][]runtime.RawExtension{}
	rCtx.Resources.Conditions = map[string][]runtime.RawExtension{}

	shouldPass := r.fnRunner.Run(fnCtx, rCtx.FunctionConfig, rCtx.Resources, results)
	// copy the results in the resourceContext
	rCtx.Results = results
	return shouldPass, nil
}

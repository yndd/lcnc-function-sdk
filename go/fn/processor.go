package fn

// ResourceContextProcessor is implemented by configuration functions built with this framework
type ResourceContextProcessor interface {
	Process(fnCtx *ResourceContext) (bool, error)
}

// ResourceContextProcessorFunc converts a compatible function to a ResourceContextProcessor.
type ResourceContextProcessorFunc func(fnCtx *ResourceContext) (bool, error)

func (p ResourceContextProcessorFunc) Process(fnCtx *ResourceContext) (bool, error) {
	return p(fnCtx)
}

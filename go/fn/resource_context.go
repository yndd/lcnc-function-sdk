package fn

import (
	"encoding/json"
	"fmt"
)

type ResourceContext struct {
	// fnconfig provides additional configuration for the function
	FunctionConfig map[string]string `json:"functionConfig,omitempty" yaml:"functionConfig,omitempty"`
	// Resources contain the resource on which this function operates
	Resources *Resources `json:"resources,omitempty" yaml:"resources,omitempty"`
	// results provide a structured
	Results *Results `json:"results,omitempty" yaml:"results,omitempty"`
}

type Resources struct {
	// holds the input KRM resources with the key being GVK in string format
	Input map[string][]string `json:"input,omitempty" yaml:"input,omitempty"`
	// holds the output KRM resources with the key being GVK in string format
	Output map[string][]string `json:"output,omitempty" yaml:"output,omitempty"`
	// holds the conditional KRM resources with the key being GVK in string format
	Conditions map[string][]string `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

func ParseResourceContext(input []byte) (*ResourceContext, error) {
	rCtx := &ResourceContext{}
	if err := json.Unmarshal(input, rCtx); err != nil {
		return nil, fmt.Errorf("error: %s with bytes: %s", err.Error(), string(input))
	}
	return rCtx, nil
}

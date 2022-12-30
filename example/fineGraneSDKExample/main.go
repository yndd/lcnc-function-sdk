package main

import (
	"context"
	"os"

	"github.com/yndd/lcnc-function-sdk/go/fn"
)

var _ fn.Runner = &Topology{}

type Topology struct {
}

func main() {
	ctx := context.TODO()
	if err := fn.AsMain(fn.WithContext(ctx, &Topology{})); err != nil {
		os.Exit(1)
	}
}

func (r *Topology) Run(ctx *fn.Context, functionConfig map[string]string, resources *fn.Resources, results *fn.Results) bool {
	if resources.Conditions == nil {
		resources.Conditions = make(map[string][]string)
	}
	if _, ok := resources.Output["topo.yndd.io.v1alpha1.Topology"]; !ok {
		resources.Output["topo.yndd.io.v1alpha1.Topology"] = make([]string, 0)
	}
	resources.Output["topo.yndd.io.v1alpha1.Topology"] = append(
		resources.Output["topo.yndd.io.v1alpha1.Topology"], "dummy topo spec")

	return true
}

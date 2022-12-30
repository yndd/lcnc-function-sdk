package main

import (
	"context"
	"os"

	"github.com/yndd/lcnc-function-sdk/go/fn"
	"k8s.io/apimachinery/pkg/runtime"
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

func (r *Topology) Run(ctx *fn.Context, functionConfig map[string]runtime.RawExtension, resources *fn.Resources, results *fn.Results) bool {

	if _, ok := resources.Output["topo.yndd.io.v1alpha1.Topology"]; !ok {
		resources.Output["topo.yndd.io.v1alpha1.Topology"] = make([]runtime.RawExtension, 0)
	}

	resources.Output["topo.yndd.io.v1alpha1.Topology"] = append(
		resources.Output["topo.yndd.io.v1alpha1.Topology"], runtime.RawExtension{Raw: []byte("{\"a\": \"b\"}")})
	resources.Output["topo.yndd.io.v1alpha1.Topology"] = append(
		resources.Output["topo.yndd.io.v1alpha1.Topology"], runtime.RawExtension{Raw: []byte("{\"a\": \"b\"}")})
	return true
}

package main

import (
	"context"
	"os"

	//topov1alpha1 "github.com/henderiw-k8s-lcnc/topology/apis/topo/v1alpha1"
	"github.com/yndd/lcnc-function-sdk/go/fn"
	"k8s.io/apimachinery/pkg/runtime"
	//"sigs.k8s.io/yaml"
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

	/*
		b, _ := os.ReadFile("../definition.yaml")
		d := &topov1alpha1.Definition{}
		yaml.Unmarshal(b, d)
	*/

	if _, ok := resources.Output["topo.yndd.io.v1alpha1.Topology"]; !ok {
		resources.Output["topo.yndd.io.v1alpha1.Topology"] = make([]runtime.RawExtension, 0)
	}

	resources.Output["topo.yndd.io.v1alpha1.Topology"] = append(
		resources.Output["topo.yndd.io.v1alpha1.Topology"], runtime.RawExtension{Raw: []byte("{\"a\": \"b\"}")})
	resources.Output["topo.yndd.io.v1alpha1.Topology"] = append(
		resources.Output["topo.yndd.io.v1alpha1.Topology"], runtime.RawExtension{Raw: []byte("{\"a\": \"b\"}")})

	//resources.AddOutput(d)
	return true
}

package main

import (
	"fmt"
	"os"

	"github.com/yndd/lcnc-function-sdk/go/fn"
	"k8s.io/apimachinery/pkg/runtime"
)

func main() {
	if err := fn.AsMain(fn.ResourceContextProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}

func Run(fnCtx *fn.ResourceContext) (bool, error) {
	fmt.Printf("topology process start\n")
	t := Topology{}
	t.Process(fnCtx)
	return true, nil
}

type Topology struct {
}

func (t *Topology) Process(fnCtx *fn.ResourceContext) {
	// WITH THIS SDK USAGE YOU NEED TO TAKE CARE OF ALL THE PRECAUTIONS
	// INIT OUTPUT/CONDITIONS/RESULT
	if fnCtx.Resources.Conditions == nil {
		fnCtx.Resources.Conditions = make(map[string][]runtime.RawExtension)
	}
	if fnCtx.Resources.Output == nil {
		fnCtx.Resources.Output = make(map[string][]runtime.RawExtension)
	}
	if _, ok := fnCtx.Resources.Output["topology.topo.yndd.io"]; !ok {
		fnCtx.Resources.Output["topology.topo.yndd.io"] = make([]runtime.RawExtension, 0)
	}
	fnCtx.Resources.Output["topo.yndd.io.v1alpha1.Topology"] = append(
		fnCtx.Resources.Output["topo.yndd.io.v1alpha1.Topology"], runtime.RawExtension{Raw: []byte("{\"a\": \"b\"}")})
	fnCtx.Resources.Output["topo.yndd.io.v1alpha1.Topology"] = append(
		fnCtx.Resources.Output["topo.yndd.io.v1alpha1.Topology"], runtime.RawExtension{Raw: []byte("{\"a\": \"b\"}")})
}

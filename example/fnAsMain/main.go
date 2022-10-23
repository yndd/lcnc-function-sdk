package main

import (
	"fmt"
	"os"

	"github.com/yndd/lcnc-function-sdk/go/fn"
	rctxv1 "github.com/yndd/lcnc-runtime/pkg/api/resourcecontext/v1"
)

func main() {
	if err := fn.AsMain(fn.ResourceContextProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}

func Run(rctx *rctxv1.ResourceContext) (bool, error) {
	fmt.Printf("topology process start\n")
	t := Topology{}
	t.Process(rctx)
	return true, nil
}

type Topology struct {
}

func (t *Topology) Process(rctx *rctxv1.ResourceContext) {
	if rctx.Spec.Properties.Allocations == nil {
		rctx.Spec.Properties.Allocations = make(map[string][]rctxv1.KRMResource)
	}
	if _, ok := rctx.Spec.Properties.Allocations["topology.topo.yndd.io"]; !ok {
		rctx.Spec.Properties.Allocations["topology.topo.yndd.io"] = make([]rctxv1.KRMResource, 0)
	}
	rctx.Spec.Properties.Allocations["topology.topo.yndd.io"] = append(
		rctx.Spec.Properties.Allocations["topology.topo.yndd.io"],
		"dummy topo spec",
	)
	fmt.Printf("topology process result: %v", rctx)
}

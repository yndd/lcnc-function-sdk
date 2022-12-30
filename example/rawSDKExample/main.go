package main

import (
	"fmt"
	"os"

	"github.com/yndd/lcnc-function-sdk/go/fn"
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
		fnCtx.Resources.Conditions = make(map[string][]string)
	}
	if _, ok := fnCtx.Resources.Output["topology.topo.yndd.io"]; !ok {
		fnCtx.Resources.Output["topology.topo.yndd.io"] = make([]string, 0)
	}
	fnCtx.Resources.Output["topology.topo.yndd.io"] = append(
		fnCtx.Resources.Output["topology.topo.yndd.io"], "dummy topo spec",
	)
}

package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/graphs/callGraph"
	"wasma/pkg/wasmp/modules"
)

type CallGraphAnalysis struct{}

// Analyze
// This analysis creates a call graph of the given
// wasm file an saves it as a dot file.
func (callGraphAnalysis *CallGraphAnalysis) Analyze(module *modules.Module, args map[string]string) {
	ic := true
	if args["ic"] == "false" {
		ic = false
	}

	log.Println(fmt.Sprintf("find indirect calls: %v", ic))
	// create call graph
	callGraph, err := callGraph.NewCallGraph(module, ic)
	if err != nil {
		log.Fatal(err.Error())
	}

	// save call graph as dot file
	fileName := strings.TrimSuffix(filepath.Base(args["file"]), filepath.Ext(args["file"]))
	err = callGraph.SaveDot(filepath.Join(args["out"], fileName) + ".dot")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (callGraphAnalysis *CallGraphAnalysis) Name() string {
	return "call graph analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&CallGraphAnalysis{})
}

package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/graphs/controlFlowGraph"
	"wasma/pkg/wasmp/modules"
)

type ControlFlowGraphAnalysis struct{}

// Analyze
// This analysis creates a control flow graph of the given.
func (controlFlowGraphAnalysis *ControlFlowGraphAnalysis) Analyze(module *modules.Module, args map[string]string) {
	var cfgs map[uint32]*controlFlowGraph.CFG

	if funcIdxStr, found := args["fi"]; found {
		funcIdx, err := strconv.Atoi(funcIdxStr)
		if err != nil {
			log.Fatal(err.Error())
		}
		if args["fi"] != "-1" {
			cfgs = controlFlowGraph.NewControlFlowGraph(module, false, uint32(funcIdx))
			fileName := strings.TrimSuffix(filepath.Base(args["file"]), filepath.Ext(args["file"])) + fmt.Sprintf("_%v", funcIdx)
			err := cfgs[uint32(funcIdx)].SaveDot(filepath.Join(args["out"], fileName) + ".dot")
			if err != nil {
				log.Fatal(err.Error())
			}
			return
		} else {
			cfgs = controlFlowGraph.NewControlFlowGraph(module, true, 0)
		}
	} else {
		cfgs = controlFlowGraph.NewControlFlowGraph(module, true, 0)
	}

	// save control flow graphs as dot file
	if codeSection, err := module.GetCodeSection(); err == nil {
		for funcIdx, _ := range codeSection.Codes {
			fileName := strings.TrimSuffix(filepath.Base(args["file"]), filepath.Ext(args["file"])) + fmt.Sprintf("_%v", funcIdx)
			err := cfgs[funcIdx].SaveDot(filepath.Join(args["out"], fileName) + ".dot")
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

func (controlFlowGraphAnalysis *ControlFlowGraphAnalysis) Name() string {
	return "control flow graph analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&ControlFlowGraphAnalysis{})
}

package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/graphs/dataFlowGraph"
	"wasma/pkg/wasmp/modules"
)

type DataFlowGraphAnalysis struct{}

// Analyze
// This analysis creates a control flow graph of the given.
func (dataFlowGraphAnalysis *DataFlowGraphAnalysis) Analyze(module *modules.Module, args map[string]string) {

	if funcIdxStr, found := args["fi"]; found {
		funcIdx, err := strconv.Atoi(funcIdxStr)
		if err != nil {
			log.Fatal(err.Error())
		}
		if args["fi"] == "-1" && strings.ToLower(args["cdfg"]) != "true" {
			log.Fatal("Parameter function index (-fi) is not set.")
		}

		complete := false

		if strings.ToLower(args["cdfg"]) == "true" {
			funcIdx = 0
			complete = true
		}

		start := time.Now()
		dfgs := dataFlowGraph.NewDataFlowGraph(module, complete, uint32(funcIdx))
		log.Printf("Data-flow graph construction for %v took %v\n", args["file"], time.Since(start))

		if complete {
			for _, dfg := range dfgs {
				fileName := strings.TrimSuffix(filepath.Base(args["file"]), filepath.Ext(args["file"])) + fmt.Sprintf("_%v", dfg.FuncIdx)
				err = dfg.SaveDot(filepath.Join(args["out"], fileName) + ".dot")
				if err != nil {
					log.Fatal(err.Error())
				}
			}
		} else {
			fileName := strings.TrimSuffix(filepath.Base(args["file"]), filepath.Ext(args["file"])) + fmt.Sprintf("_%v", funcIdx)
			err = dfgs[uint32(funcIdx)].SaveDot(filepath.Join(args["out"], fileName) + ".dot")
			if err != nil {
				log.Fatal(err.Error())
			}
		}

	} else {
		log.Fatal("Parameter function index (-fi) is not set.")
	}
}

func (dataFlowGraphAnalysis *DataFlowGraphAnalysis) Name() string {
	return "data flow graph analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&DataFlowGraphAnalysis{})
}

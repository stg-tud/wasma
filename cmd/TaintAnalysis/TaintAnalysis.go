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

type TaintAnalysis struct {
}

// Analyze
// This analysis creates a control flow graph of the given.
func (taintAnalysis *TaintAnalysis) Analyze(module *modules.Module, args map[string]string) {

	returnValueIsTainted := true

	if importSection, err := module.GetImportSection(); err == nil {
		for _, customSectionImport := range importSection.Imports {
			log.Printf("Sections Import: %v", customSectionImport.Imp.Name)
		}
	}
	if exportSection, err := module.GetExportSection(); err == nil {
		for _, customSectionExport := range exportSection.Exports {
			log.Printf("Sections Export: %v", customSectionExport.Name)
		}
	}

	if funcIdxStr, found := args["fi"]; found {
		funcIdx, err := strconv.Atoi(funcIdxStr)
		if err != nil {
			log.Fatal(err.Error())
		}
		if args["fi"] == "-1" {
			log.Fatal("Parameter function index (-fi) is not set.")
		}

		start := time.Now()
		// 2 paramter auf true setzen für alle funktionen
		dfgs := dataFlowGraph.NewDataFlowGraph(module, false, uint32(funcIdx))

		/*
			for _, dataFlowGraph := range dfgs {
				for _, dataFlowEdges := range dataFlowGraph.Tree {
					for _, dataFlowEdge := range dataFlowEdges {
						if dataFlowGraph.Disassembly.DisassembledInstrs[dataFlowEdge.Input].Tainted {
							dataFlowEdge.Variable.Tainted = true
							dataFlowGraph.Disassembly.DisassembledInstrs[dataFlowEdge.Output].Tainted = true

						}
					}
				}
			}
		*/

		log.Printf("Taint construction for %v took %v\n", args["file"], time.Since(start))

		fileName := strings.TrimSuffix(filepath.Base(args["file"]), filepath.Ext(args["file"])) + fmt.Sprintf("_%v", funcIdx)
		filePath := filepath.Join(args["out"], fileName) + ".dot"
		err = dfgs[uint32(funcIdx)].SaveDot(filePath)
		if err != nil {
			log.Fatal(err.Error())
		}
		filePathAbs, err := filepath.Abs(filePath)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("Taint construction saved to %v", filePathAbs)

		returnValueIsTainted = false

	} else {
		log.Fatal("Parameter function index (-fi) is not set.")
	}

	log.Printf("Return value is tainted: %v", returnValueIsTainted)

}

func (taintAnalysis *TaintAnalysis) Name() string {
	return "taint analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&TaintAnalysis{})
}

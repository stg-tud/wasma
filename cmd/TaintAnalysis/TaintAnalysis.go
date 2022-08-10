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

func (taintAnalysis *TaintAnalysis) FindSinks(dfgs map[uint32]*dataFlowGraph.DFG) []string {
	var foundSinks []string

	// get calls to sinks from flow
	// how are function params given to call
	// see if they are tainted
	// if var coresponding to call is tinted add to sinks
	for _, dataFlowGraph := range dfgs {
		for instrIdx, dataFlowEdge := range dataFlowGraph.Environment.Flow {
			instruction := dataFlowGraph.Disassembly.DisassembledInstrs[instrIdx].Instruction
			value := ""
			isCall := false
			isCallTainted := false

			switch instruction.Name() {
			case "call":
				{
					value = fmt.Sprintf("%v", instruction.ToString())
					isCall = true
				}
			case "call_indirect":
				{
					value = fmt.Sprintf("indirect %v", instruction.ToString())
					isCall = true
				}
			}

			if isCall {
				for _, varIn := range dataFlowEdge.Input {
					if varIn.Tainted {
						isCallTainted = true
						// log.Printf("Taint in call found %v %v %v\n", varIn.VariableName, dataFlowEdge.Input, varIn.Tainted)
					}
				}
			}

			if isCallTainted {
				foundSinks = append(foundSinks, value)
			}
		}
	}

	return foundSinks
}

// Analyze
// This analysis creates a control flow graph of the given.
func (taintAnalysis *TaintAnalysis) Analyze(module *modules.Module, args map[string]string) {

	returnValueIsTainted := false

	if importSection, err := module.GetImportSection(); err == nil {
		for _, customSectionImport := range importSection.Imports {
			log.Printf("Sections Import: %v %v", customSectionImport.Index, customSectionImport.Imp.Name)
		}
	}
	if exportSection, err := module.GetExportSection(); err == nil {
		for _, customSectionExport := range exportSection.Exports {
			log.Printf("Sections Export: %v %v", customSectionExport.ExportDesc.FuncIdx, customSectionExport.Name)
		}
	}

	if funcNameStr, found := args["fn"]; found {
		if exportSection, err := module.GetExportSection(); err == nil {
			for _, customSectionExport := range exportSection.Exports {
				if funcNameStr == customSectionExport.Name {
					log.Printf("Analyse function: %v %v", customSectionExport.ExportDesc.FuncIdx, customSectionExport.Name)

					praramsToCheck := []int{}
					if funcParams, found := args["fp"]; found {
						if funcParams != "" {
							praramsToCheckStr := strings.Split(funcParams, ",")
							for _, i := range praramsToCheckStr {
								j, err := strconv.Atoi(i)
								if err != nil {
									panic(err)
								}
								praramsToCheck = append(praramsToCheck, j)
							}
						}
					}
					funcIdx := customSectionExport.ExportDesc.FuncIdx

					start := time.Now()
					// 2 paramter auf true setzen für alle funktionen
					dfgs := dataFlowGraph.NewDataFlowGraphWithTaint(module, false, uint32(funcIdx), praramsToCheck)

					// get calls to sinks
					sinks := taintAnalysis.FindSinks(dfgs)
					if len(sinks) == 0 {
						log.Printf("No sinks found\n")
					} else {
						for _, sink := range sinks {
							log.Printf("Sink found %v\n", sink)
						}
					}

					//get taint of return value
					for _, dataFlowGraph := range dfgs {
						for _, dataFlowEdges := range dataFlowGraph.Tree {
							for _, dataFlowEdge := range dataFlowEdges {
								if varIns := dataFlowEdge.Input; varIns == "return" {
									log.Printf("Return found %v %v %v\n", dataFlowEdge.Variable.VariableName, dataFlowEdge.Output, dataFlowEdge.Tainted)
									if dataFlowEdge.Tainted {
										returnValueIsTainted = true
									}
								}
							}
						}
					}

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
				}
			}
		}
	} else {
		log.Fatal("Parameter function name (-fn) is not set.")
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

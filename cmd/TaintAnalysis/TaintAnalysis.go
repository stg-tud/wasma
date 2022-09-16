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

type FunctionCall struct {
	funcIdx     uint32
	instruction string
	name        string
}

type TaintAnalysis struct {
}

func (taintAnalysis *TaintAnalysis) GetIdsOfSourceNames(sources []string, module *modules.Module) []dataFlowGraph.Source {
	var funcIdxs []dataFlowGraph.Source
	if importSection, err := module.GetImportSection(); err == nil {
		for _, funcNameStr := range sources {
			for _, customImportExport := range importSection.Imports {
				if funcNameStr == customImportExport.Imp.Name {
					source := dataFlowGraph.Source{Name: funcNameStr, FuncIdx: customImportExport.Index}
					funcIdxs = append(funcIdxs, source)
				}
				//log.Printf("Analyse function: %v %v", funcIdx, customSectionExport.Name)
			}
		}
	}
	return funcIdxs
}

func (taintAnalysis *TaintAnalysis) GetIdOfEntrypointFunction(args map[string]string, module *modules.Module) uint32 {
	var funcIdx uint32
	if exportSection, err := module.GetExportSection(); err == nil {
		if funcNameStr, found := args["fn"]; found && funcNameStr != "" {
			for _, customSectionExport := range exportSection.Exports {
				if funcNameStr == customSectionExport.Name {
					funcIdx = customSectionExport.ExportDesc.FuncIdx
				}
			}
		} else if funcNameStr, found := args["fi"]; found && funcNameStr != "" {
			funcIdxInt, err := strconv.Atoi(funcNameStr)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				funcIdx = uint32(funcIdxInt)
			}
		} else {
			log.Fatal("Parameter function name (-fn) or (-fi) is not set.")
		}
	}
	return funcIdx
}

func (taintAnalysis *TaintAnalysis) GetInitialTaintedParameters(args map[string]string) []uint32 {
	praramsToCheck := []uint32{}
	if funcParams, found := args["fp"]; found {
		if funcParams != "" {
			praramsToCheckStr := strings.Split(funcParams, ",")
			for _, i := range praramsToCheckStr {
				j, err := strconv.Atoi(i)
				if err != nil {
					panic(err)
				}
				praramsToCheck = append(praramsToCheck, uint32(j))
			}
		}
	}
	return praramsToCheck
}

func (taintAnalysis *TaintAnalysis) GetKnownSources() []string {
	knownSources := []string{"fd_read", "args_get"}
	return knownSources
}

func (taintAnalysis *TaintAnalysis) GetKnownSinks() []string {
	knownSinks := []string{"fd_write"}
	return knownSinks
}

func (taintAnalysis *TaintAnalysis) FindSinks(dfgs map[uint32]*dataFlowGraph.DFG, module *modules.Module) map[uint32][]FunctionCall {

	// get calls to sinks from flow
	// how are function params given to call
	// see if they are tainted
	// if var coresponding to call is tainted add to sinks
	foundSinks := make(map[uint32][]FunctionCall)
	for dfgId, dataFlowGraph := range dfgs {

		for instrIdx, dataFlowEdge := range dataFlowGraph.Environment.Flow {
			instruction := dataFlowGraph.Disassembly.DisassembledInstrs[instrIdx].Instruction

			var functionCall FunctionCall
			isCall := false
			isCallTainted := false

			switch instruction.Name() {
			case "call":
				{
					functionCall.instruction = instruction.ToString()
					isCall = true
				}
			case "call_indirect":
				{
					functionCall.instruction = fmt.Sprintf("indirect %v", instruction.ToString())
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
				functionIdx, error := instruction.Funcidx()
				if error != nil && error.Error() == "attribute not available" {
					// do nothing
				} else if error != nil {
					log.Printf("Error %v with instruction %v", error, instruction.Name())
				}

				if importSection, err := module.GetImportSection(); err == nil {
					if val, ok := importSection.FuncImports[functionIdx]; ok {
						functionCall.name = val.Name
					}

					if val, ok := importSection.TableImports[functionIdx]; ok {
						functionCall.name = val.Name
					}

					if val, ok := importSection.MemImports[functionIdx]; ok {
						functionCall.name = val.Name
					}

					if val, ok := importSection.GlobalImports[functionIdx]; ok {
						functionCall.name = val.Name
					}

					functionCall.funcIdx = functionIdx

				} else {
					log.Printf("Error %v with instruction %v", error, instruction.Name())
				}

				for _, knownSink := range taintAnalysis.GetKnownSinks() {
					if functionCall.name == knownSink {
						foundSinks[dfgId] = append(foundSinks[dfgId], functionCall)
					}
				}
			} else {
				functionIdx, error := instruction.Funcidx()
				if error != nil && error.Error() == "attribute not available" {
					// do nothing
				} else if error != nil {
					log.Printf("Error %v with instruction %v", error, instruction.Name())
				}

				if importSection, err := module.GetImportSection(); err == nil {
					if val, ok := importSection.FuncImports[functionIdx]; ok {
						functionCall.name = val.Name
					}

					if val, ok := importSection.TableImports[functionIdx]; ok {
						functionCall.name = val.Name
					}

					if val, ok := importSection.MemImports[functionIdx]; ok {
						functionCall.name = val.Name
					}

					if val, ok := importSection.GlobalImports[functionIdx]; ok {
						functionCall.name = val.Name
					}

					functionCall.funcIdx = functionIdx

				} else {
					log.Printf("Error %v with instruction %v", error, instruction.Name())
				}

				for _, knownSink := range taintAnalysis.GetKnownSinks() {
					if functionCall.name == knownSink {
						log.Printf("Found no sink without taint in %v with instruction name %v", dfgId, instruction.Name())
					}
				}
			}
		}
	}

	return foundSinks
}

// Analyze
// This analysis creates a taint flow graph of the given.
func (taintAnalysis *TaintAnalysis) Analyze(module *modules.Module, args map[string]string) {

	returnValueIsTainted := false

	/*
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
	*/

	funcIdx := taintAnalysis.GetIdOfEntrypointFunction(args, module)
	paramsToCheck := taintAnalysis.GetInitialTaintedParameters(args)
	sourcesName := taintAnalysis.GetKnownSources()
	sourceIdxs := taintAnalysis.GetIdsOfSourceNames(sourcesName, module)

	// find all calls and indirect calls that come after funcIdx of entry point

	start := time.Now()
	// 2 paramter auf true setzen für alle funktionen
	dfgs := dataFlowGraph.NewDataFlowGraphWithTaint(module, true, uint32(funcIdx), paramsToCheck, sourceIdxs)

	// get calls to sinks
	sinks := taintAnalysis.FindSinks(dfgs, module)
	if len(sinks) == 0 {
		log.Printf("No sinks found\n")
	} else {
		for funcId, sink := range sinks {
			log.Printf("Sink with id %v and object %v found\n", funcId, sink)
		}
	}

	//get taint of return value
	for _, dataFlowEdges := range dfgs[funcIdx].Tree {
		for _, dataFlowEdge := range dataFlowEdges {
			if varIns := dataFlowEdge.Input; varIns == "return" {
				if dataFlowEdge.Tainted {
					returnValueIsTainted = true
					//log.Printf("Return found %v %v %v\n", dataFlowEdge.Variable.VariableName, dataFlowEdge.Output, dataFlowEdge.Tainted)
				}
			}
		}
	}

	log.Printf("Taint construction for %v took %v\n", args["file"], time.Since(start))

	fileName := strings.TrimSuffix(filepath.Base(args["file"]), filepath.Ext(args["file"])) + fmt.Sprintf("_%v", funcIdx)
	filePath := filepath.Join(args["out"], fileName) + ".dot"
	err := dfgs[uint32(funcIdx)].SaveDot(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	filePathAbs, err := filepath.Abs(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Taint construction saved to %v", filePathAbs)

	log.Printf("Return value is tainted: %v", returnValueIsTainted)

}

func (taintAnalysis *TaintAnalysis) Name() string {
	return "taint analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&TaintAnalysis{})
}

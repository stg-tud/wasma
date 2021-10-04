package main

import (
	"fmt"
	"log"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/code"
	"wasma/pkg/wasma/output"
	"wasma/pkg/wasmp/modules"
)

type OverapproximationCallIndirectAnalysis struct{}

func countTableFuncIdxs(funcIdxs []uint32, tableFuncIdxs map[uint32]bool) uint32 {
	var counter uint32
	for _, funcIdx := range funcIdxs {
		if _, found := tableFuncIdxs[funcIdx]; found {
			counter++
		}
	}
	return counter
}

func (overapproximationCallIndirectAnalysis *OverapproximationCallIndirectAnalysis) Analyze(module *modules.Module, args map[string]string) {
	file := args["file"]
	out := args["out"]
	var tableFuncIdxs map[uint32]bool //:= code.GetTableFuncIdxs(module)
	if elementSection, err := module.GetElementSection(); err == nil {
		tableFuncIdxs = elementSection.FuncIdxs
	}

	csvFile, err := output.OpenOrCreateCSV(out)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer csvFile.Close()

	if functionSection, err := module.GetFunctionSection(); err == nil {
		if codeSection, err := module.GetCodeSection(); err == nil {
			for funcIdx, _ := range codeSection.Codes {
				disassembly := code.DisassemblyFunction(funcIdx, module)
				for _, instrDisassembly := range disassembly.DisassembledInstrs {
					if instrDisassembly.Instruction.Name() == "call_indirect" {
						if typeIdx, err := instrDisassembly.Instruction.Typeidx(); err == nil {
							funcIdxs := functionSection.FuncIdxs[typeIdx]
							var funcIdxsImports []uint32

							if importSection, err := module.GetImportSection(); err == nil {
								funcIdxsImports = importSection.FuncIdxs[typeIdx]
							}

							csvFile.Write([]string{
								file,
								fmt.Sprintf("%v", funcIdx),
								"call_indirect",
								fmt.Sprintf("%v", typeIdx),
								fmt.Sprintf("%v", len(funcIdxs)+len(funcIdxsImports)),
								fmt.Sprintf("%v",
									countTableFuncIdxs(funcIdxs, tableFuncIdxs)+
										countTableFuncIdxs(funcIdxsImports, tableFuncIdxs))})
						}
					}
				}
			}

		}
	}
}

func (overapproximationCallIndirectAnalysis *OverapproximationCallIndirectAnalysis) Name() string {
	return "overapproximation call_indirect analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&OverapproximationCallIndirectAnalysis{})
}

package main

import (
	"fmt"
	"log"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/code"
	"wasma/pkg/wasma/output"
	"wasma/pkg/wasmp/modules"
)

type OverapproximationBrTableAnalysis struct{}

func (overapproximationBrTableAnalysis *OverapproximationBrTableAnalysis) Analyze(module *modules.Module, args map[string]string) {
	file := args["file"]
	out := args["out"]

	csvFile, err := output.OpenOrCreateCSV(out)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer csvFile.Close()

	if codeSection, err := module.GetCodeSection(); err == nil {
		for funcIdx, _ := range codeSection.Codes {
			disassembly := code.DisassemblyFunction(funcIdx, module)
			for _, instrDisassembly := range disassembly.DisassembledInstrs {
				if instrDisassembly.Instruction.Name() == "br_table" {
					vecLabelIdx, _ := instrDisassembly.Instruction.VecLabelidx()
					csvFile.Write([]string{
						file,
						fmt.Sprintf("%v", funcIdx),
						"br_table",
						fmt.Sprintf("%v", len(vecLabelIdx)+1)}) // +1 for the default case
				}
			}
		}
	}
}

func (overapproximationBrTableAnalysis *OverapproximationBrTableAnalysis) Name() string {
	return "overapproximation br_table analysis)"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&OverapproximationBrTableAnalysis{})
}

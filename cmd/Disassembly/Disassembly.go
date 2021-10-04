package main

import (
	"fmt"
	"strconv"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/code"
	"wasma/pkg/wasmp/modules"
)

type DisassemblyAnalysis struct{}

func (disassemblyAnalysis *DisassemblyAnalysis) Analyze(module *modules.Module, args map[string]string) {
	funcIdx, errStrConv := strconv.Atoi(args["fi"])

	if errStrConv == nil && funcIdx >= 0 {
		disassembly := code.DisassemblyFunction(uint32(funcIdx), module)
		for _, line := range disassembly.GetTextRepresentation() {
			fmt.Printf("%v\n", line)
		}
	}
}

func (disassemblyAnalysis *DisassemblyAnalysis) Name() string {
	return "disassembly analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&DisassemblyAnalysis{})
}

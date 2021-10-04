package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/code"
	"wasma/pkg/wasma/output"
	"wasma/pkg/wasmp/modules"
	"wasma/pkg/wasmp/values"
)

type InstructionCountAnalysis struct{}

func count(counts *map[string]Count, value string, opcode byte) {
	if entry, found := (*counts)[value]; found {
		entry.count = entry.count + 1
		(*counts)[value] = entry
	} else {
		(*counts)[value] = Count{value, opcode, 1}
	}
}

type Count struct {
	instruction string
	opcode      byte
	count       int
}

func (instructionCountAnalysis *InstructionCountAnalysis) Analyze(module *modules.Module, args map[string]string) {
	log.Printf("start instruction count analysis: %v", args["file"])
	outputFile, err := output.OpenOrCreateTXT(args["out"])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	instructionCounts := make(map[string]Count)
	instructionCountsWithImmediates := make(map[string]Count)
	var numberCounts = 0

	if values.GetElseCounter() > 0 {
		numberCounts = numberCounts + values.GetElseCounter()
		instructionCounts["else"] = Count{"else", 0x05, values.GetElseCounter()}
		instructionCountsWithImmediates["else"] = Count{"else", 0x05, values.GetElseCounter()}
	}

	if values.GetEndCounter() > 0 {
		numberCounts = numberCounts + values.GetEndCounter()
		instructionCounts["end"] = Count{"end", 0x0B, values.GetEndCounter()}
		instructionCountsWithImmediates["end"] = Count{"end", 0x0B, values.GetEndCounter()}
	}

	if functionSection, err := module.GetFunctionSection(); err == nil {
		for funcIdx, _ := range functionSection.TypeIdxs {
			for _, instrDisassembly := range code.DisassemblyFunction(funcIdx, module).DisassembledInstrs {
				count(&instructionCounts, instrDisassembly.Instruction.Name(), instrDisassembly.Instruction.Opcode())
				//count(&instructionCountsWithImmediates, instrDisassembly.Instruction.ToString(), instrDisassembly.Instruction.Opcode())
				numberCounts++
			}
		}
	}

	outputFile.WriteString(fmt.Sprintf("Total opcodes: %v\n\nOpcode counts:\n", numberCounts))

	outputData(outputFile, instructionCounts)

	outputFile.WriteString("\nOpcode counts with immediates:\n")

	//outputData(outputFile, instructionCountsWithImmediates)

	log.Printf("finished instruction count analysis: %v", args["file"])
}

func (instructionCountAnalysis *InstructionCountAnalysis) Name() string {
	return "instruction count analysis"
}

func outputData(outputFile *os.File, outputData map[string]Count) {
	// sort output data by count and opcode
	var outputs []Count
	for _, value := range outputData {
		outputs = append(outputs, value)
	}

	sort.Slice(outputs, func(i, j int) bool {
		//if outputs[i].count > outputs[j].count {
		//	return true
		//}
		//if outputs[i].count < outputs[j].count {
		//	return false
		//}
		//return outputs[i].opcode < outputs[j].opcode
		if outputs[i].count > outputs[j].count {
			return true
		}
		if outputs[i].count < outputs[j].count {
			return false
		}
		if outputs[i].opcode < outputs[j].opcode {
			return true
		}
		if outputs[i].opcode > outputs[j].opcode {
			return false
		}
		return outputs[i].instruction < outputs[j].instruction
	})

	// output data after sorting
	for _, output := range outputs {
		outputFile.WriteString(fmt.Sprintf("%v: %v\n", output.instruction, output.count))
	}
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&InstructionCountAnalysis{})
}

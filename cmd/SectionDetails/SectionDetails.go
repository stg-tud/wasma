package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/output"
	"wasma/pkg/wasmp/modules"
)

type SectionDetailsAnalysis struct{}

var head string

func sectionTypeDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Type.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if typeSection, err := module.GetTypeSection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Type[%v]:\n", len(typeSection.FunctionTypes)))
		var functionTypes []uint32
		for typeIdx, _ := range typeSection.FunctionTypes {
			functionTypes = append(functionTypes, typeIdx)
		}
		sort.Slice(functionTypes, func(i, j int) bool {
			return functionTypes[i] < functionTypes[j]
		})

		for _, typeIdx := range functionTypes {
			functionType := typeSection.FunctionTypes[typeIdx]
			outputFile.WriteString(fmt.Sprintf(" - type[%v] (%v) -> %v\n", typeIdx, functionType.ParameterTypesToString(), functionType.ResultTypeToString()))
		}
	}
}

func sectionImportDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Import.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if importSection, err := module.GetImportSection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Import[%v]:\n", len(importSection.Imports)))
		for _, importValue := range importSection.Imports {
			switch importValue.Imp.ImportDesc.ImportType {
			case 0x00:
				outputFile.WriteString(fmt.Sprintf(" - func[%v]\n", importValue.Index))
			case 0x01:
				outputFile.WriteString(fmt.Sprintf(" - table[%v]\n", importValue.Index))
			case 0x02:
				outputFile.WriteString(fmt.Sprintf(" - memory[%v]\n", importValue.Index))
			case 0x03:
				outputFile.WriteString(fmt.Sprintf(" - global[%v]\n", importValue.Index))
			}
		}
	}
}

func sectionFunctionDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Function.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if functionSection, err := module.GetFunctionSection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Function[%v]:\n", len(functionSection.TypeIdxs)))
		var functionTypes []uint32
		for typeIdx, _ := range functionSection.TypeIdxs {
			functionTypes = append(functionTypes, typeIdx)
		}
		sort.Slice(functionTypes, func(i, j int) bool {
			return functionTypes[i] < functionTypes[j]
		})

		for _, typeIdx := range functionTypes {
			functionType := functionSection.TypeIdxs[typeIdx]
			outputFile.WriteString(fmt.Sprintf(" - func[%v] sig=%v\n", typeIdx, functionType))
		}
	}
}

func sectionTableDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Table.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if tableSection, err := module.GetTableSection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Table[%v]:\n", len(tableSection.Tables)))
		var tableTypes []uint32
		for tableIdx, _ := range tableSection.Tables {
			tableTypes = append(tableTypes, tableIdx)
		}
		sort.Slice(tableTypes, func(i, j int) bool {
			return tableTypes[i] < tableTypes[j]
		})

		for _, tableIdx := range tableTypes {
			tableType := tableSection.Tables[tableIdx]
			if tableType.Limit.Type == 0x00 {
				outputFile.WriteString(fmt.Sprintf(" - table[%v] type=funcref initial=%v\n", tableIdx, tableType.Limit.Min))
			} else {
				outputFile.WriteString(fmt.Sprintf(" - table[%v] type=funcref initial=%v max=%v\n", tableIdx, tableType.Limit.Min, tableType.Limit.Max))
			}
		}
	}
}

func sectionMemoryDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Memory.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if memorySection, err := module.GetMemorySection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Memory[%v]:\n", len(memorySection.MemTypes)))
		var memoryTypes []uint32
		for memoryIdx, _ := range memorySection.MemTypes {
			memoryTypes = append(memoryTypes, memoryIdx)
		}
		sort.Slice(memoryTypes, func(i, j int) bool {
			return memoryTypes[i] < memoryTypes[j]
		})

		for _, memoryIdx := range memoryTypes {
			memoryType := memorySection.MemTypes[memoryIdx]
			if memoryType.Type == 0x00 {
				outputFile.WriteString(fmt.Sprintf(" - memory[%v] pages: initial=%v\n", memoryIdx, memoryType.Min))
			} else {
				outputFile.WriteString(fmt.Sprintf(" - memory[%v] pages: initial=%v max=%v\n", memoryIdx, memoryType.Min, memoryType.Max))
			}
		}
	}
}

func sectionGlobalDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Global.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if globalSection, err := module.GetGlobalSection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Global[%v]:\n", len(globalSection.Globals)))
		var globals []uint32
		for globalIdx, _ := range globalSection.Globals {
			globals = append(globals, globalIdx)
		}
		sort.Slice(globals, func(i, j int) bool {
			return globals[i] < globals[j]
		})

		for _, globalIdx := range globals {
			outputFile.WriteString(fmt.Sprintf(" - global[%v] %v mutable=%v\n", globalIdx, globalSection.Globals[globalIdx].GlobalType.ValType, globalSection.Globals[globalIdx].GlobalType.Mut))
		}

	}
}

func sectionExportDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Export.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if exportSection, err := module.GetExportSection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Export[%v]:\n", len(exportSection.Exports)))
		for _, export := range exportSection.Exports {
			outputFile.WriteString(fmt.Sprintf(" - %v\n", export.ToString()))
		}
	}
}

func sectionStartDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Start.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if startSection, err := module.GetStartSection(); err == nil {
		outputFile.WriteString("Start:\n")
		outputFile.WriteString(fmt.Sprintf(" - start function: %v\n", startSection.FuncIdx))
	}
}

func sectionElementDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Elem.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if elementSection, err := module.GetElementSection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Elem[%v]:\n", len(elementSection.Elements)))
		for _, element := range elementSection.Elements {
			outputFile.WriteString(fmt.Sprintf("count=%v\n", len(element.FuncIdxs)))
			for _, funcIdx := range element.FuncIdxs {
				outputFile.WriteString(fmt.Sprintf("func[%v]\n", funcIdx))
			}
		}
	}
}

func sectionCodeDetails(outputFileName string, module *modules.Module) {
	outputFile, err := output.OpenOrCreateTXT(outputFileName + "_Code.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(head)

	if codeSection, err := module.GetCodeSection(); err == nil {
		outputFile.WriteString(fmt.Sprintf("Code[%v]:\n", len(codeSection.Codes)))
		var funcIdxs []uint32
		for funcIdx, _ := range codeSection.Codes {
			funcIdxs = append(funcIdxs, funcIdx)
		}
		sort.Slice(funcIdxs, func(i, j int) bool {
			return funcIdxs[i] < funcIdxs[j]
		})

		for _, funcIdx := range funcIdxs {
			code := codeSection.Codes[funcIdx]
			outputFile.WriteString(fmt.Sprintf(" - func[%v] size=%v\n", funcIdx, code.Size))
		}
	}
}

func (sectionDetailsAnalysis *SectionDetailsAnalysis) Analyze(module *modules.Module, args map[string]string) {
	log.Printf("start section details analysis: %v", args["file"])

	fileName := strings.TrimSuffix(filepath.Base(args["file"]), filepath.Ext(args["file"]))
	fileName = filepath.Join(args["out"], fileName)

	head = fmt.Sprintf("\n%v:\tfile format wasm 0x1\n\nSection Details:\n\n", filepath.Base(args["file"]))
	sectionTypeDetails(fileName, module)
	sectionImportDetails(fileName, module)
	sectionFunctionDetails(fileName, module)
	sectionTableDetails(fileName, module)
	sectionMemoryDetails(fileName, module)
	sectionGlobalDetails(fileName, module)
	sectionExportDetails(fileName, module)
	sectionStartDetails(fileName, module)
	sectionElementDetails(fileName, module)
	sectionCodeDetails(fileName, module)

	log.Printf("finished section details analysis: %v", args["file"])
}

func (sectionDetailsAnalysis *SectionDetailsAnalysis) Name() string {
	return "section details analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&SectionDetailsAnalysis{})
}

package main

import (
	"fmt"
	"wasma/pkg/wasma"
	"wasma/pkg/wasmp/modules"
	"wasma/pkg/wasmp/types"
)

type ImportsAnalysis struct{}

func (importsAnalysis *ImportsAnalysis) Analyze(module *modules.Module, args map[string]string) {
	functionTypes := make(map[uint32]types.FunctionType)
	if typeSection, err := module.GetTypeSection(); err == nil {
		functionTypes = typeSection.FunctionTypes
	}
	if importSection, err := module.GetImportSection(); err == nil {
		for _, imp := range importSection.FuncImports {
			impType := imp.ImportDesc.ImportType
			switch impType {
			case 0x00:
				if functionType, found := functionTypes[imp.ImportDesc.TypeIdx]; found {
					fmt.Printf("Import[modName=\"%v\", name=\"%v\", importType=\"%v\", functionType=[typeIdx=%v, parameterTypes=%v, resultTypes=%v]]\n", imp.ModName, imp.Name, impType, imp.ImportDesc.TypeIdx, functionType.ParameterTypes, functionType.ResultTypes)
				} else {
					fmt.Printf("Import[modName=\"%v\", name=\"%v\", importType=\"%v\", typeIdx=%v]\n", imp.ModName, imp.Name, impType, imp.ImportDesc.TypeIdx)
				}
			case 0x01:
				fmt.Printf("Import[modName=\"%v\", name=\"%v\", importType=\"%v\", tableType=[min=%v, max=%v]\n", imp.ModName, imp.Name, impType, imp.ImportDesc.TableType.Limit.Min, imp.ImportDesc.TableType.Limit.Max)
			case 0x02:
				fmt.Printf("Import[modName=\"%v\", name=\"%v\", importType=\"%v\", memType=[min=%v, max=%v]\n", imp.ModName, imp.Name, impType, imp.ImportDesc.MemType.Min, imp.ImportDesc.MemType.Max)
			case 0x03:
				fmt.Printf("Import[modName=\"%v\", name=\"%v\", importType=\"%v\", globalType=[valType=\"%v\", mut=%v]\n", imp.ModName, imp.Name, impType, imp.ImportDesc.GlobalType.ValType, imp.ImportDesc.GlobalType.Mut)
			default:
				fmt.Printf("error: invalid import type, got \"%v\"\n", impType)
			}
		}
	}
}

func (importsAnalysis *ImportsAnalysis) Name() string {
	return "imports analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&ImportsAnalysis{})
}

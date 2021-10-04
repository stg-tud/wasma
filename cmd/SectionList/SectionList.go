package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/output"
	"wasma/pkg/wasmp/modules"
)

type SectionListAnalysis struct{}

func write(outputFile *os.File, sectionString string, start uint32, size uint32, value string) {
	outputFile.WriteString(fmt.Sprintf("%v start=0x%08x end=0x%08x (size=0x%08x) %v\n",
		sectionString,
		start,
		start+size,
		size,
		value))
}

func (sectionListAnalysis *SectionListAnalysis) Analyze(module *modules.Module, args map[string]string) {
	log.Printf("start section list analysis: %v", args["file"])
	outputFile, err := output.OpenOrCreateTXT(args["out"])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	outputFile.WriteString(fmt.Sprintf("\n%v:\tfile format wasm 0x1\n", filepath.Base(args["file"])))
	outputFile.WriteString("\nSections:\n\n")

	if customSections, err := module.GetCustomSections(); err == nil {
		for _, customSection := range customSections {
			write(outputFile, "   Custom", customSection.StartContent, customSection.Size, fmt.Sprintf("%v", fmt.Sprintf("\"%v\"", customSection.Name)))
		}
	}

	if typeSection, err := module.GetTypeSection(); err == nil {
		write(outputFile, "     Type", typeSection.StartContent, typeSection.Size, fmt.Sprintf("count: %v", len(typeSection.FunctionTypes)))
		if customSectionsType, err := module.GetCustomSectionsType(); err == nil {
			for _, customSectionType := range customSectionsType {
				write(outputFile, "   Custom", customSectionType.StartContent, customSectionType.Size, fmt.Sprintf("\"%v\"", customSectionType.Name))
			}
		}
	}

	if importSection, err := module.GetImportSection(); err == nil {
		write(outputFile, "   Import", importSection.StartContent, importSection.Size, fmt.Sprintf("count: %v", len(importSection.FuncImports)+
			len(importSection.TableImports)+
			len(importSection.MemImports)+
			len(importSection.GlobalImports)))
		if customSectionsImport, err := module.GetCustomSectionsImport(); err == nil {
			for _, customSectionImport := range customSectionsImport {
				write(outputFile, "   Custom", customSectionImport.StartContent, customSectionImport.Size, fmt.Sprintf("\"%v\"", customSectionImport.Name))
			}
		}
	}

	if functionSection, err := module.GetFunctionSection(); err == nil {
		write(outputFile, " Function", functionSection.StartContent, functionSection.Size, fmt.Sprintf("count: %v", len(functionSection.TypeIdxs)))
		if customSectionsFunction, err := module.GetCustomSectionsFunction(); err == nil {
			for _, customSectionFunction := range customSectionsFunction {
				write(outputFile, "   Custom", customSectionFunction.StartContent, customSectionFunction.Size, fmt.Sprintf("\"%v\"", customSectionFunction.Name))
			}
		}
	}

	if tableSection, err := module.GetTableSection(); err == nil {
		write(outputFile, "    Table", tableSection.StartContent, tableSection.Size, fmt.Sprintf("count: %v", len(tableSection.Tables)))
		if customSectionsTable, err := module.GetCustomSectionsTable(); err == nil {
			for _, customSectionTable := range customSectionsTable {
				write(outputFile, "   Custom", customSectionTable.StartContent, customSectionTable.Size, fmt.Sprintf("\"%v\"", customSectionTable.Name))
			}
		}
	}

	if memorySection, err := module.GetMemorySection(); err == nil {
		write(outputFile, "   Memory", memorySection.StartContent, memorySection.Size, fmt.Sprintf("count: %v", len(memorySection.MemTypes)))
		if customSectionsMemory, err := module.GetCustomSectionsMemory(); err == nil {
			for _, customSectionMemory := range customSectionsMemory {
				write(outputFile, "   Custom", customSectionMemory.StartContent, customSectionMemory.Size, fmt.Sprintf("\"%v\"", customSectionMemory.Name))
			}
		}
	}

	if globalSection, err := module.GetGlobalSection(); err == nil {
		write(outputFile, "   Global", globalSection.StartContent, globalSection.Size, fmt.Sprintf("count: %v", len(globalSection.Globals)))
		if customSectionsGlobal, err := module.GetCustomSectionsGlobal(); err == nil {
			for _, customSectionGlobal := range customSectionsGlobal {
				write(outputFile, "   Custom", customSectionGlobal.StartContent, customSectionGlobal.Size, fmt.Sprintf("\"%v\"", customSectionGlobal.Name))
			}
		}
	}

	if exportSection, err := module.GetExportSection(); err == nil {
		write(outputFile, "   Export", exportSection.StartContent, exportSection.Size, fmt.Sprintf("count: %v", len(exportSection.Exports)))
		if customSectionsExport, err := module.GetCustomSectionsExport(); err == nil {
			for _, customSectionExport := range customSectionsExport {
				write(outputFile, "   Custom", customSectionExport.StartContent, customSectionExport.Size, fmt.Sprintf("\"%v\"", customSectionExport.Name))
			}
		}
	}

	if startSection, err := module.GetStartSection(); err == nil {
		write(outputFile, "    Start", startSection.StartContent, startSection.Size, fmt.Sprintf("start: %v", startSection.FuncIdx))
		if customSectionsStart, err := module.GetCustomSectionsStart(); err == nil {
			for _, customSectionStart := range customSectionsStart {
				write(outputFile, "   Custom", customSectionStart.StartContent, customSectionStart.Size, fmt.Sprintf("\"%v\"", customSectionStart.Name))
			}
		}
	}

	if elementSection, err := module.GetElementSection(); err == nil {
		write(outputFile, "     Elem", elementSection.StartContent, elementSection.Size, fmt.Sprintf("count: %v", len(elementSection.Elements)))
		if customSectionsElement, err := module.GetCustomSectionsElement(); err == nil {
			for _, customSectionElement := range customSectionsElement {
				write(outputFile, "   Custom", customSectionElement.StartContent, customSectionElement.Size, fmt.Sprintf("\"%v\"", customSectionElement.Name))
			}
		}
	}

	if codeSection, err := module.GetCodeSection(); err == nil {
		write(outputFile, "     Code", codeSection.StartContent, codeSection.Size, fmt.Sprintf("count: %v", len(codeSection.Codes)))
		if customSectionsCode, err := module.GetCustomSectionsCode(); err == nil {
			for _, customSectionCode := range customSectionsCode {
				write(outputFile, "   Custom", customSectionCode.StartContent, customSectionCode.Size, fmt.Sprintf("\"%v\"", customSectionCode.Name))
			}
		}
	}

	if dataSection, err := module.GetDataSection(); err == nil {
		write(outputFile, "     Data", dataSection.StartContent, dataSection.Size, fmt.Sprintf("count: %v", len(dataSection.Datas)))
		if customSectionsData, err := module.GetCustomSectionsData(); err == nil {
			for _, customSectionData := range customSectionsData {
				write(outputFile, "   Custom", customSectionData.StartContent, customSectionData.Size, fmt.Sprintf("\"%v\"", customSectionData.Name))
			}
		}
	}

	log.Printf("finished section list analysis: %v", args["file"])
}

func (sectionListAnalysis *SectionListAnalysis) Name() string {
	return "section list analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&SectionListAnalysis{})
}

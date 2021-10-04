package main

import (
	"fmt"
	"wasma/pkg/wasma"
	"wasma/pkg/wasmp/modules"
)

type ExportsAnalysis struct{}

func (exportsAnalysis *ExportsAnalysis) Analyze(module *modules.Module, args map[string]string) {
	var funcIdxs []uint32
	if exportSection, err := module.GetExportSection(); err == nil {
		for _, export := range exportSection.Exports {
			if idx, idxTypeS, err := export.ExportDesc.GetIdx(); err == nil {
				fmt.Printf("Export[name=\"%v\", type=\"%v\", typeName=\"%v\", idx=\"%v\"]\n", export.Name, export.ExportDesc.ExportType, idxTypeS, idx)
				if idxTypeS == "funcIdx" {
					funcIdxs = append(funcIdxs, idx)
				}
			} else {
				fmt.Println(err.Error())
				fmt.Printf("Export[name=\"%v\", type=\"%v\"\n", export.Name, export.ExportDesc.ExportType)
			}
		}
	}
}

func (exportsAnalysis *ExportsAnalysis) Name() string {
	return "exports analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&ExportsAnalysis{})
}

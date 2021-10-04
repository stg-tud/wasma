package main

import (
	"log"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/output"
	"wasma/pkg/wasmp/modules"
)

type DataAnalysis struct{}

func (dataAnalysis *DataAnalysis) Analyze(module *modules.Module, args map[string]string) {
	out := args["out"]

	csvFile, err := output.OpenOrCreateCSV(out)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer csvFile.Close()

	if dataSection, err := module.GetDataSection(); err == nil {
		for _, data := range dataSection.Datas {
			csvFile.Write([]string{string(data.Bytes)})
		}
	}
}

func (dataAnalysis *DataAnalysis) Name() string {
	return "data analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&DataAnalysis{})
}

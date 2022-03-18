build:
	@echo "Build standard analyses ..."
	go build -o bin/CallGraph cmd/CallGraph/CallGraph.go
	go build -o bin/ControlFlowGraph cmd/ControlFlowGraph/ControlFlowGraph.go
	go build -o bin/Data cmd/Data/Data.go
	go build -o bin/DataFlowGraph cmd/DataFlowGraph/DataFlowGraph.go
	go build -o bin/Disassembly cmd/Disassembly/Disassembly.go
	go build -o bin/Exports cmd/Exports/Exports.go
	go build -o bin/Imports cmd/Imports/Imports.go
	go build -o bin/InstructionCount cmd/InstructionCount/InstructionCount.go
	go build -o bin/OverapproximationBrTable cmd/OverapproximationBrTable/OverapproximationBrTable.go
	go build -o bin/OverapproximationCallIndirect cmd/OverapproximationCallIndirect/OverapproximationCallIndirect.go
	go build -o bin/SectionDetails cmd/SectionDetails/SectionDetails.go
	go build -o bin/SectionList cmd/SectionList/SectionList.go
	go build -o bin/wasma-fmng cmd/wasma-fmng/wasma-fmng.go
	go build -o bin/wasma-fm cmd/wasma-fm/wasma-fm.go

build-win:
	@echo "Build standard analyses (Windows) ..."
	GOOS=windows go build -o bin/CallGraph.exe cmd/CallGraph/CallGraph.go
	GOOS=windows go build -o bin/ControlFlowGraph.exe cmd/ControlFlowGraph/ControlFlowGraph.go
	GOOS=windows go build -o bin/Data.exe cmd/Data/Data.go
	GOOS=windows go build -o bin/DataFlowGraph.exe cmd/DataFlowGraph/DataFlowGraph.go
	GOOS=windows go build -o bin/Disassembly.exe cmd/Disassembly/Disassembly.go
	GOOS=windows go build -o bin/Exports.exe cmd/Exports/Exports.go
	GOOS=windows go build -o bin/Imports.exe cmd/Imports/Imports.go
	GOOS=windows go build -o bin/InstructionCount.exe cmd/InstructionCount/InstructionCount.go
	GOOS=windows go build -o bin/OverapproximationBrTable.exe cmd/OverapproximationBrTable/OverapproximationBrTable.go
	GOOS=windows go build -o bin/OverapproximationCallIndirect.exe cmd/OverapproximationCallIndirect/OverapproximationCallIndirect.go
	GOOS=windows go build -o bin/SectionDetails.exe cmd/SectionDetails/SectionDetails.go
	GOOS=windows go build -o bin/SectionList.exe cmd/SectionList/SectionList.go
	GOOS=windows go build -o bin/wasma-fmng.exe cmd/wasma-fmng/wasma-fmng.go
	GOOS=windows go build -o bin/wasma-fm.exe cmd/wasma-fm/wasma-fm.go


package dataFlowGraph

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"wasma/pkg/wasma/code"
	"wasma/pkg/wasma/graphs/controlFlowGraph"
	structuresWasma "wasma/pkg/wasma/structures"
	"wasma/pkg/wasmp/modules"
	"wasma/pkg/wasmp/structures"
	structuresWasmp "wasma/pkg/wasmp/structures"
)

type Source struct {
	Name    string
	FuncIdx uint32
}

type FlowEdge struct {
	Variable structuresWasmp.Variable
	// instrIdx
	Output string
	// instrIdx
	Input   string
	Tainted bool
}

type DFG struct {
	FuncIdx     uint32
	Environment *structuresWasma.Environment
	// key: variableIdx
	Tree        map[uint32][]FlowEdge
	Disassembly code.Disassembly
}

func (dfg *DFG) SaveDot(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// key: instrIdx, value: parameter name
	params := make(map[string]string)

	file.WriteString("digraph G {\n")

	for _, local := range dfg.Environment.Locals {
		color := ""
		if local.Taint.Tainted {
			color = "style=filled, fillcolor=red,"
		}

		if local.VariableType == "P" {
			file.WriteString(fmt.Sprintf("%v [shape=box, "+color+" label=\"#%v: (param %v, tainted %v)\"];\n", local.VariableName, dfg.FuncIdx, local.VariableName, local.Taint.Tainted))
		} else if local.VariableType == "L" {
			file.WriteString(fmt.Sprintf("%v [shape=box, "+color+" label=\"#%v: (local %v)\"];\n", local.VariableName, dfg.FuncIdx, local.VariableName))
		}
	}

	for _, global := range dfg.Environment.Globals {
		color := ""
		if global.Taint.Tainted {
			color = "style=filled, fillcolor=red,"
		}
		file.WriteString(fmt.Sprintf("%v [shape=box, "+color+" label=\"#%v: (global %v)\"];\n", global.VariableName, dfg.FuncIdx, global.VariableName))
	}

	// return
	file.WriteString(fmt.Sprintf("return [shape=box, label=\"#%v: return\"];\n", dfg.FuncIdx))

	//for instrIdx, dataFlowEdge := range dfg.environment.Flow {
	for instrIdx := range dfg.Environment.Flow {
		instruction := dfg.Disassembly.DisassembledInstrs[instrIdx].Instruction
		var value = ""

		switch instruction.Name() {
		case "local.get", "local.set", "local.tee":
			{
				localidx, _ := instruction.Localidx()
				value = fmt.Sprintf("%v", localidx)
			}
		case "global.get", "global.set":
			{
				globalidx, _ := instruction.Globalidx()
				value = fmt.Sprintf("%v", globalidx)
			}
		case "call":
			{
				funcidx, _ := instruction.Funcidx()
				value = fmt.Sprintf("%v", funcidx)
			}
		case "call_indirect":
			{
				typeidx, _ := instruction.Typeidx()
				value = fmt.Sprintf("%v", typeidx)
			}
		case "i32.const":
			{
				i32, _ := instruction.I32()
				value = fmt.Sprintf("%v", i32)
			}
		case "i64.const":
			{
				i64, _ := instruction.I64()
				value = fmt.Sprintf("%v", i64)
			}
		case "f32.const":
			{
				f32, _ := instruction.F32()
				value = fmt.Sprintf("%v", f32)
			}
		case "f64.const":
			{
				f64, _ := instruction.F64()
				value = fmt.Sprintf("%v", f64)
			}
		}

		file.WriteString(fmt.Sprintf("%v [label=\"#%v+%v: %v %v\"];\n", instrIdx, dfg.FuncIdx, instrIdx, instruction.Name(), value))
	}

	for _, flowEdges := range dfg.Tree {
		for _, flowEdge := range flowEdges {
			if flowEdge.Variable.VariableType == "P" && !strings.HasPrefix(flowEdge.Input, "P") {
				params[flowEdge.Output] = flowEdge.Variable.VariableName
			}
		}
	}

	for _, flowEdges := range dfg.Tree {
		for _, flowEdge := range flowEdges {
			color := ""
			if flowEdge.Variable.Value != "unknown" {
				if flowEdge.Tainted {
					color = "color=red,"
				}
				file.WriteString(fmt.Sprintf("%v -> %v ["+color+" label=\"%v\"];\n", flowEdge.Output, flowEdge.Input, flowEdge.Variable.Value))
			} else {
				if flowEdge.Tainted {
					color = "color=red"
				}
				file.WriteString(fmt.Sprintf("%v -> %v ["+color+"];\n", flowEdge.Output, flowEdge.Input))
			}
		}
	}

	file.WriteString("}")
	return nil
}

var environment *structuresWasma.Environment
var isEntireMemoryTainted bool

func AnalyseTaintOfFunction(module *modules.Module,
	funcIdx uint32,
	funcParams map[uint32]structures.Taint,
	sources []Source,
	cfg map[uint32]*controlFlowGraph.CFG,
	dfg map[uint32]*DFG,
	visited map[uint32]bool,
	uses_indirect_call *bool,
	uses_memory *bool) map[uint32]*DFG {

	visited[funcIdx] = true

	subCfg := cfg[funcIdx]
	subEnvironment := structuresWasma.NewEnvironment(module)

	functionType, err := code.GetFuncParams(funcIdx, module)
	if err == nil {
		for i, param := range functionType.ParameterTypes {
			taint := structures.Taint{Tainted: false}
			for paramIndex, v := range funcParams {
				if paramIndex == uint32(i) {
					taint = v
				}
			}
			subEnvironment.NewParameterWithTaint(param, taint)
		}
	}

	for _, local := range code.GetFuncLocals(funcIdx, module) {
		subEnvironment.NewLocal(local)
	}

	if globals, startIndex, err := code.GetGlobalsList(module); err == nil {
		subEnvironment.SetGlobalIdx(startIndex)
		for _, global := range globals {
			subEnvironment.NewGlobal(global.GlobalType.Mut, global.GlobalType.ValType)
		}
	}

	if start, found := subCfg.Tree[0]; found {
		walk(subEnvironment, start, subCfg.Tree, make(map[uint32]bool))
	}

	// search for function calls and start analysis there
	for instrIdx, dataFlowEdgeOuter := range subEnvironment.Flow {

		instruction := subCfg.Disassembly.DisassembledInstrs[instrIdx].Instruction

		var foundNewFunctions []uint32

		// taint sources
		isSourceCall := false
		var sourceLocal Source

		instruction.Name()

		switch instruction.Name() {
		case "call":
			{
				for _, source := range sources {
					if funcidx, error := instruction.Funcidx(); funcidx == source.FuncIdx && error == nil {
						isSourceCall = true
						sourceLocal = source
					} else {
						instruction.Funcidx()
					}
				}

			}
		case "call_indirect":
			{
				*uses_indirect_call = true
			}
		case "i32.store", "i64.store", "f32.store", "f64.store", "i32.store8", "i32.store16", "i64.store8", "i64.store16", "i64.store32":
			{
				isEntireMemoryTainted = true
				*uses_memory = true
			}
		}

		if isSourceCall {
			// taint the instruction/variable itself

			for _, varOut := range dataFlowEdgeOuter.Output {
				varOut.Taint.Tainted = true
				varOut.Taint.Source.Name = sourceLocal.Name
				varOut.Taint.Source.FuncIdx = sourceLocal.FuncIdx
				varOut.Taint.Source.Instruction = instruction.Name()
				primaryVariableIdx := varOut.PrimaryVariableIdx
				variableName := varOut.VariableName

				subEnvironment.Variables[primaryVariableIdx] = varOut

				// also taint vars with same name
				for varIdx0, varOut0 := range subEnvironment.Variables {
					if varOut0.PrimaryVariableIdx == primaryVariableIdx || varOut0.VariableName == variableName {
						varOut0.Taint.Tainted = true
						varOut0.Taint.Source.Name = sourceLocal.Name
						varOut0.Taint.Source.FuncIdx = sourceLocal.FuncIdx
						varOut0.Taint.Source.Instruction = instruction.Name()
						subEnvironment.Variables[varIdx0] = varOut0
					}
				}

				for instrIdx2, dataFlowEdgeOuter := range subEnvironment.Flow {
					// instruction -> variable
					for varIdx2, varOut2 := range dataFlowEdgeOuter.Input {
						if varOut2.PrimaryVariableIdx == primaryVariableIdx || varOut2.VariableName == variableName {
							varOut2.Taint.Tainted = true
							varOut2.Taint.Source.Name = sourceLocal.Name
							varOut2.Taint.Source.FuncIdx = sourceLocal.FuncIdx
							varOut2.Taint.Source.Instruction = instruction.Name()
							subEnvironment.Flow[instrIdx2].Input[varIdx2] = varOut2
						}
					}
					for varIdx3, varOut3 := range dataFlowEdgeOuter.Output {
						if varOut3.PrimaryVariableIdx == primaryVariableIdx || varOut3.VariableName == variableName {
							varOut3.Taint.Tainted = true
							varOut3.Taint.Source.Name = sourceLocal.Name
							varOut3.Taint.Source.FuncIdx = sourceLocal.FuncIdx
							varOut3.Taint.Source.Instruction = instruction.Name()
							subEnvironment.Flow[instrIdx2].Output[varIdx3] = varOut3
						}
					}
				}

			}
		}

		// propagate taint
		// variable -> instruction
		for _, varIn := range dataFlowEdgeOuter.Input {
			if varIn.Taint.Tainted {
				for _, varOut := range dataFlowEdgeOuter.Output {
					varOut.Taint = varIn.Taint
					primaryVariableIdx := varOut.PrimaryVariableIdx
					variableName := varOut.VariableName
					//subEnvironment.Flow[instrIdx].Output[primaryVariableIdx] = varOut
					subEnvironment.Variables[primaryVariableIdx] = varOut

					// also taint vars with same name
					for varIdx0, varOut0 := range subEnvironment.Variables {
						if varOut0.PrimaryVariableIdx == primaryVariableIdx || varOut0.VariableName == variableName {
							varOut0.Taint = varIn.Taint
							subEnvironment.Variables[varIdx0] = varOut0
						}
					}

					for instrIdx2, dataFlowEdgeOuter := range subEnvironment.Flow {
						// instruction -> variable
						for varIdx2, varOut2 := range dataFlowEdgeOuter.Input {
							if varOut2.PrimaryVariableIdx == primaryVariableIdx || varOut2.VariableName == variableName {
								varOut2.Taint = varIn.Taint
								subEnvironment.Flow[instrIdx2].Input[varIdx2] = varOut2
							}
						}
						for varIdx3, varOut3 := range dataFlowEdgeOuter.Output {
							if varOut3.PrimaryVariableIdx == primaryVariableIdx || varOut3.VariableName == variableName {
								varOut3.Taint = varIn.Taint
								subEnvironment.Flow[instrIdx2].Output[varIdx3] = varOut3
							}
						}
					}

				}
			}
		}

		switch instruction.Name() {
		case "call":
			{
				if funcIdxFromCall, error := instruction.Funcidx(); !visited[funcIdxFromCall] && error == nil {
					foundNewFunctions = append(foundNewFunctions, funcIdxFromCall)
				}
			}
		case "call_indirect":
			{
				typeIdx, err := instruction.Typeidx()
				if err != nil {
					log.Fatal(err.Error())
				}
				if functionSection, err := module.GetFunctionSection(); err == nil {
					var funcIdxs []uint32
					// select all functions with the given type index
					if fIdxs, found := functionSection.FuncIdxs[typeIdx]; found {
						funcIdxs = fIdxs
					} else {
						funcIdxs = []uint32{}
					}

					// select all imported functions with the given type index
					if importSection, err := module.GetImportSection(); err == nil {
						if importFuncIdxs, found := importSection.FuncIdxs[typeIdx]; found {
							// add imported functions
							funcIdxs = append(funcIdxs, importFuncIdxs...)
						}

					}
					for _, funcIdx := range funcIdxs {
						if elementSection, err := module.GetElementSection(); err == nil {
							if _, found := elementSection.FuncIdxs[funcIdx]; found {
								// Only add the FuncIdx for indirect calls if the given
								// FuncIdx is an element of the funcRef table.
								if _, error := instruction.Typeidx(); !visited[funcIdx] && error == nil {
									foundNewFunctions = append(foundNewFunctions, funcIdx)
								}
							}
						}
					}
				}
			}
		case "i32.load",
			"i64.load",
			"f32.load",
			"f64.load",
			"i32.load8_s",
			"i32.load8_u",
			"i32.load16_s",
			"i32.load16_u",
			"i64.load8_s",
			"i64.load8_u",
			"i64.load16_s",
			"i64.load16_u",
			"i64.load32_s",
			"i64.load32_u":
			{
				if isEntireMemoryTainted {
					for _, varOut := range dataFlowEdgeOuter.Output {
						varOut.Taint.Tainted = true
						varOut.Taint.Source.Name = sourceLocal.Name
						varOut.Taint.Source.FuncIdx = sourceLocal.FuncIdx
						varOut.Taint.Source.Instruction = instruction.Name()
						primaryVariableIdx := varOut.PrimaryVariableIdx
						variableName := varOut.VariableName
						subEnvironment.Variables[primaryVariableIdx] = varOut

						// also taint vars with same name
						for varIdx0, varOut0 := range subEnvironment.Variables {
							if varOut0.PrimaryVariableIdx == primaryVariableIdx || varOut0.VariableName == variableName {
								varOut0.Taint.Tainted = true
								varOut0.Taint.Source.Name = sourceLocal.Name
								varOut0.Taint.Source.FuncIdx = sourceLocal.FuncIdx
								varOut0.Taint.Source.Instruction = instruction.Name()
								subEnvironment.Variables[varIdx0] = varOut0
							}
						}

						for instrIdx2, dataFlowEdge := range subEnvironment.Flow {
							// instruction -> variable
							for varIdx2, varOut2 := range dataFlowEdge.Input {
								if varOut2.PrimaryVariableIdx == primaryVariableIdx || varOut2.VariableName == variableName {
									varOut2.Taint.Tainted = true
									varOut2.Taint.Source.Name = sourceLocal.Name
									varOut2.Taint.Source.FuncIdx = sourceLocal.FuncIdx
									varOut2.Taint.Source.Instruction = instruction.Name()
									subEnvironment.Flow[instrIdx2].Input[varIdx2] = varOut2
								}
							}
							for varIdx3, varOut3 := range dataFlowEdge.Output {
								if varOut3.PrimaryVariableIdx == primaryVariableIdx || varOut3.VariableName == variableName {
									varOut3.Taint.Tainted = true
									varOut3.Taint.Source.Name = sourceLocal.Name
									varOut3.Taint.Source.FuncIdx = sourceLocal.FuncIdx
									varOut3.Taint.Source.Instruction = instruction.Name()
									subEnvironment.Flow[instrIdx2].Output[varIdx3] = varOut3
								}
							}
						}
					}
				}
			}
		}

		for _, foundNewFunction := range foundNewFunctions {

			// not an exported function
			var importedFunctions []uint32
			if importSection, err := module.GetImportSection(); err == nil {
				for _, customSectionImport := range importSection.Imports {
					importedFunctions = append(importedFunctions, customSectionImport.Index)
				}
			}
			contains := false
			for _, importedFunction := range importedFunctions {
				if importedFunction == foundNewFunction {
					contains = true
					visited[foundNewFunction] = true
				}
			}
			if !contains {

				// get tainted parameters
				taintedParams := make(map[uint32]structures.Taint)
				var oneParamTainted structures.Taint
				// get number of params

				// possible improvement add shadow stack and only taint param if stack value is tainted

				for _, varIn := range dataFlowEdgeOuter.Input {
					if varIn.Taint.Tainted {
						oneParamTainted = varIn.Taint
						//log.Printf("Function %v param %v is %v\n", foundNewFunction, varInIndex, varIn)
					}
				}

				for varInIndex := range dataFlowEdgeOuter.Input {
					if oneParamTainted.Tainted {
						taintedParams[uint32(varInIndex)] = oneParamTainted
						//log.Printf("Function %v param %v is %v\n", foundNewFunction, varInIndex, varIn)
					}
				}

				dfg = AnalyseTaintOfFunction(module,
					foundNewFunction,
					taintedParams,
					sources,
					cfg,
					dfg,
					visited,
					uses_indirect_call,
					uses_memory)

				// taint return value from call

				for _, dataFlowEdges := range dfg[foundNewFunction].Tree {
					for _, dataFlowEdge := range dataFlowEdges {
						if varIns := dataFlowEdge.Input; varIns == "return" {
							if dataFlowEdge.Tainted {
								// taint call
								for _, varOut := range dataFlowEdgeOuter.Output {
									varOut.Taint.Tainted = true
									primaryVariableIdx := varOut.PrimaryVariableIdx
									variableName := varOut.VariableName
									subEnvironment.Variables[primaryVariableIdx] = varOut

									// also taint vars with same name
									for varIdx0, varOut0 := range subEnvironment.Variables {
										if varOut0.PrimaryVariableIdx == primaryVariableIdx || varOut0.VariableName == variableName {
											varOut0.Taint.Tainted = true
											subEnvironment.Variables[varIdx0] = varOut0
										}
									}

									for instrIdx2, dataFlowEdge := range subEnvironment.Flow {
										// instruction -> variable
										for varIdx2, varOut2 := range dataFlowEdge.Input {
											if varOut2.PrimaryVariableIdx == primaryVariableIdx || varOut2.VariableName == variableName {
												varOut2.Taint.Tainted = true
												subEnvironment.Flow[instrIdx2].Input[varIdx2] = varOut2
											}
										}
										for varIdx3, varOut3 := range dataFlowEdge.Output {
											if varOut3.PrimaryVariableIdx == primaryVariableIdx || varOut3.VariableName == variableName {
												varOut3.Taint.Tainted = true
												subEnvironment.Flow[instrIdx2].Output[varIdx3] = varOut3
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	tree := GetFlowTreeWithTaint(subEnvironment, sources, subCfg)

	// remove duplicates
	edges := make(map[string]bool)

	for i := range tree {
		var flowEdges []FlowEdge
		for _, flowEdge := range tree[i] {
			edge := fmt.Sprintf("%v;%v", flowEdge.Input, flowEdge.Output)
			if _, found := edges[edge]; !found {
				flowEdges = append(flowEdges, flowEdge)
				edges[edge] = true
			}
		}
		if len(flowEdges) > 0 {
			tree[i] = flowEdges
		} else {
			delete(tree, i)
		}
	}

	dfg[funcIdx] = &DFG{funcIdx, subEnvironment, tree, subCfg.Disassembly}

	environment = subEnvironment

	return dfg
}

func NewDataFlowGraphWithTaint(module *modules.Module, complete bool,
	funcIdx uint32, funcParams map[uint32]structures.Taint,
	sources []Source,
	uses_indirect_call *bool, uses_memory *bool) map[uint32]*DFG {
	// key: FuncIdx
	dfg := make(map[uint32]*DFG)
	cfg := controlFlowGraph.NewControlFlowGraph(module, complete, funcIdx)
	visited := make(map[uint32]bool)

	dfg = AnalyseTaintOfFunction(module, funcIdx, funcParams, sources, cfg, dfg, visited, uses_indirect_call, uses_memory)

	// debug
	// print all tainted vars
	// log.Printf("Visited functions %v\n", visited)

	/*
		for visitedFunctionIdx, dataFlowGraph := range dfg {
			for _, tree := range dataFlowGraph.Tree {
				for flowIdx, flow := range tree {
					if flow.Tainted || flow.Variable.Tainted {
						log.Printf("Tainted function %v with flow %v. Flowname: %v \n", visitedFunctionIdx, flowIdx, flow.Variable.VariableName)
					}
				}
			}
		}
	*/

	// find calls
	return dfg
}

func NewDataFlowGraph(module *modules.Module, complete bool, funcIdx uint32) map[uint32]*DFG {
	// key: FuncIdx
	dfg := make(map[uint32]*DFG)
	cfg := controlFlowGraph.NewControlFlowGraph(module, complete, funcIdx)

	for funcIdx, subCfg := range cfg {
		environment = structuresWasma.NewEnvironment(module)

		functionType, err := code.GetFuncParams(funcIdx, module)
		if err == nil {
			for _, param := range functionType.ParameterTypes {
				environment.NewParameter(param)
			}
		}

		for _, local := range code.GetFuncLocals(funcIdx, module) {
			environment.NewLocal(local)
		}

		if globals, startIndex, err := code.GetGlobalsList(module); err == nil {
			environment.SetGlobalIdx(startIndex)
			for _, global := range globals {
				environment.NewGlobal(global.GlobalType.Mut, global.GlobalType.ValType)
			}
		}

		if start, found := subCfg.Tree[0]; found {
			walk(environment, start, subCfg.Tree, make(map[uint32]bool))
		}

		tree := GetFlowTree(environment)

		// remove duplicates
		edges := make(map[string]bool)

		for i := range tree {
			var flowEdges []FlowEdge
			for _, flowEdge := range tree[i] {
				edge := fmt.Sprintf("%v;%v", flowEdge.Input, flowEdge.Output)
				if _, found := edges[edge]; !found {
					flowEdges = append(flowEdges, flowEdge)
					edges[edge] = true
				}
			}
			if len(flowEdges) > 0 {
				tree[i] = flowEdges
			} else {
				delete(tree, i)
			}
		}

		dfg[funcIdx] = &DFG{funcIdx, environment, tree, subCfg.Disassembly}
	}
	return dfg
}

func GetFlowTreeWithTaint(environment *structuresWasma.Environment, sources []Source, cfg *controlFlowGraph.CFG) map[uint32][]FlowEdge {
	// key: variableIdx
	variables := make(map[uint32]structuresWasmp.Variable)
	// key: variableIdx
	input := make(map[uint32][]string)
	// key: variableIdx
	output := make(map[uint32][]string)
	// key: variableIdx
	tree := make(map[uint32][]FlowEdge)

	for instrIdx, dataFlowEdge := range environment.Flow {
		// variable -> instruction
		for _, varIn := range dataFlowEdge.Input {
			variables[varIn.PrimaryVariableIdx] = varIn
			input[varIn.PrimaryVariableIdx] = append(input[varIn.PrimaryVariableIdx], fmt.Sprintf("%v", instrIdx))
		}
	}

	for instrIdx, dataFlowEdge := range environment.Flow {
		// instruction -> variable
		for _, varOut := range dataFlowEdge.Output {
			variables[varOut.PrimaryVariableIdx] = varOut
			output[varOut.PrimaryVariableIdx] = append(output[varOut.PrimaryVariableIdx], fmt.Sprintf("%v", instrIdx))

			// add input for locals and globals
			if _, found := input[varOut.PrimaryVariableIdx]; !found {
				if varOut.LocalGlobalIn && (varOut.VariableType == "P" || varOut.VariableType == "L" || varOut.VariableType == "GC" || varOut.VariableType == "GM") {
					input[varOut.PrimaryVariableIdx] = append(input[varOut.PrimaryVariableIdx], varOut.VariableName)
				}
			}
		}
	}

	// add output for locals
	for _, local := range environment.Locals {
		output[local.PrimaryVariableIdx] = append(output[local.PrimaryVariableIdx], local.VariableName)
	}

	// add output for globals
	for _, global := range environment.Globals {
		output[global.PrimaryVariableIdx] = append(output[global.PrimaryVariableIdx], global.VariableName)
	}

	for primaryVariableIdx := range variables {
		if varOuts, found := output[primaryVariableIdx]; found {
			if varIns, found := input[primaryVariableIdx]; found {
				for _, varIn := range varIns {
					for _, varOut := range varOuts {
						tainted := false
						if variables[primaryVariableIdx].Taint.Tainted {
							tainted = true
						}
						tree[primaryVariableIdx] = append(
							tree[primaryVariableIdx],
							FlowEdge{variables[primaryVariableIdx],
								varOut,
								varIn,
								tainted})
					}
				}
			} else {
				for _, varOut := range varOuts {
					tainted := false
					if variables[primaryVariableIdx].Taint.Tainted {
						tainted = true
					}
					tree[primaryVariableIdx] = append(
						tree[primaryVariableIdx],
						FlowEdge{variables[primaryVariableIdx],
							varOut,
							"return",
							tainted})
				}
			}
		}
	}

	return tree
}

func GetFlowTree(environment *structuresWasma.Environment) map[uint32][]FlowEdge {
	// key: variableIdx
	variables := make(map[uint32]structuresWasmp.Variable)
	// key: variableIdx
	input := make(map[uint32][]string)
	// key: variableIdx
	output := make(map[uint32][]string)
	// key: variableIdx
	tree := make(map[uint32][]FlowEdge)

	for instrIdx, dataFlowEdge := range environment.Flow {
		// variable -> instruction
		for _, varIn := range dataFlowEdge.Input {
			variables[varIn.PrimaryVariableIdx] = varIn
			input[varIn.PrimaryVariableIdx] = append(input[varIn.PrimaryVariableIdx], fmt.Sprintf("%v", instrIdx))
		}
	}

	for instrIdx, dataFlowEdge := range environment.Flow {
		// instruction -> variable
		for _, varOut := range dataFlowEdge.Output {
			variables[varOut.PrimaryVariableIdx] = varOut
			output[varOut.PrimaryVariableIdx] = append(output[varOut.PrimaryVariableIdx], fmt.Sprintf("%v", instrIdx))

			// add input for locals and globals
			if _, found := input[varOut.PrimaryVariableIdx]; !found {
				if varOut.LocalGlobalIn && (varOut.VariableType == "P" || varOut.VariableType == "L" || varOut.VariableType == "GC" || varOut.VariableType == "GM") {
					input[varOut.PrimaryVariableIdx] = append(input[varOut.PrimaryVariableIdx], varOut.VariableName)
				}
			}
		}
	}

	// propagate taint
	for _, dataFlowEdge := range environment.Flow {
		// variable -> instruction
		for _, varIn := range dataFlowEdge.Input {
			if varIn.Taint.Tainted {
				for _, varOut := range dataFlowEdge.Output {
					varOut.Taint.Tainted = true
					primaryVariableIdx := varOut.PrimaryVariableIdx
					variableName := varOut.VariableName
					//environment.Flow[instrIdx].Output[primaryVariableIdx] = varOut
					environment.Variables[primaryVariableIdx] = varOut
					variables[primaryVariableIdx] = varOut

					// also taint vars with same name
					for varIdx0, varOut0 := range environment.Variables {
						if varOut0.PrimaryVariableIdx == primaryVariableIdx || varOut0.VariableName == variableName {
							varOut0.Taint.Tainted = true
							environment.Variables[varIdx0] = varOut0
						}
					}

					// also taint vars with same name
					for varIdx00, varOut00 := range variables {
						if varOut00.PrimaryVariableIdx == primaryVariableIdx || varOut00.VariableName == variableName {
							varOut00.Taint.Tainted = true
							variables[varIdx00] = varOut00
						}
					}

					for instrIdx2, dataFlowEdge := range environment.Flow {
						// instruction -> variable
						for varIdx2, varOut2 := range dataFlowEdge.Input {
							if varOut2.PrimaryVariableIdx == primaryVariableIdx || varOut2.VariableName == variableName {
								varOut2.Taint.Tainted = true
								environment.Flow[instrIdx2].Input[varIdx2] = varOut2
							}
						}
						for varIdx3, varOut3 := range dataFlowEdge.Output {
							if varOut3.PrimaryVariableIdx == primaryVariableIdx || varOut3.VariableName == variableName {
								varOut3.Taint.Tainted = true
								environment.Flow[instrIdx2].Output[varIdx3] = varOut3
							}
						}
					}

				}
			}
		}
	}

	// add output for locals
	for _, local := range environment.Locals {
		output[local.PrimaryVariableIdx] = append(output[local.PrimaryVariableIdx], local.VariableName)
	}

	// add output for globals
	for _, global := range environment.Globals {
		output[global.PrimaryVariableIdx] = append(output[global.PrimaryVariableIdx], global.VariableName)
	}

	for primaryVariableIdx := range variables {
		if varOuts, found := output[primaryVariableIdx]; found {
			if varIns, found := input[primaryVariableIdx]; found {
				for _, varIn := range varIns {
					for _, varOut := range varOuts {
						tainted := false
						if variables[primaryVariableIdx].Taint.Tainted {
							tainted = true
						}
						tree[primaryVariableIdx] = append(
							tree[primaryVariableIdx],
							FlowEdge{variables[primaryVariableIdx],
								varOut,
								varIn,
								tainted})
					}
				}
			} else {
				for _, varOut := range varOuts {
					tainted := false
					if variables[primaryVariableIdx].Taint.Tainted {
						tainted = true
					}
					tree[primaryVariableIdx] = append(
						tree[primaryVariableIdx],
						FlowEdge{variables[primaryVariableIdx],
							varOut,
							"return",
							tainted})
				}
			}
		}
	}

	return tree
}

/*
	Input a node that will be the start node in the tree
	Will run trough the wasm commands and set the environment
	Output visited map with functionIdx as key and a bool value that indicates if the node is on the path
	execute the commands like pop and push from stack
*/
func walk(environment *structuresWasma.Environment, node *controlFlowGraph.CFGNode, tree map[uint32]*controlFlowGraph.CFGNode, visited map[uint32]bool) {
	if _, found := visited[node.InstrIdx]; !found {

		if node.Control {
			node.Instruction.Executor()(node.InstrIdx, node.Instruction, environment)
		} else {
			var instrIdxs []uint32
			for instrIdx := range node.Block {
				instrIdxs = append(instrIdxs, instrIdx)
			}
			sort.Slice(instrIdxs, func(i, j int) bool {
				return instrIdxs[i] < instrIdxs[j]
			})

			for _, instrIdx := range instrIdxs {
				instr := node.Block[instrIdx]
				instr.Executor()(instrIdx, instr, environment)
			}
		}

		if len(node.Successors) == 1 {
			visited[node.InstrIdx] = true
			walk(environment, tree[node.Successors[0].TargetNode], tree, visited)
		} else if len(node.Successors) > 1 {
			stack := environment.Stack
			for _, successor := range node.Successors {
				environment.Stack = stack
				visited[node.InstrIdx] = true
				walk(environment, tree[successor.TargetNode], tree, visited)
			}
		}

	}
}

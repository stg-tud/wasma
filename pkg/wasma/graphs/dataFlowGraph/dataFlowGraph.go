package dataFlowGraph

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"wasma/pkg/wasma/code"
	"wasma/pkg/wasma/graphs/controlFlowGraph"
	structuresWasma "wasma/pkg/wasma/structures"
	"wasma/pkg/wasmp/modules"
	structuresWasmp "wasma/pkg/wasmp/structures"
)

type FlowEdge struct {
	Variable structuresWasmp.Variable
	// instrIdx
	Output string
	// instrIdx
	Input string
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
		if local.VariableType == "P" {
			file.WriteString(fmt.Sprintf("%v [shape=box, label=\"#%v: (param %v)\"];\n", local.VariableName, dfg.FuncIdx, local.VariableName))
		} else if local.VariableType == "L" {
			file.WriteString(fmt.Sprintf("%v [shape=box, label=\"#%v: (local %v)\"];\n", local.VariableName, dfg.FuncIdx, local.VariableName))
		}
	}

	for _, global := range dfg.Environment.Globals {
		file.WriteString(fmt.Sprintf("%v [shape=box, label=\"#%v: (global %v)\"];\n", global.VariableName, dfg.FuncIdx, global.VariableName))
	}

	// return
	file.WriteString(fmt.Sprintf("return [shape=box, label=\"#%v: return\"];\n", dfg.FuncIdx))

	//for instrIdx, dataFlowEdge := range dfg.environment.Flow {
	for instrIdx, _ := range dfg.Environment.Flow {
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
			if flowEdge.Variable.Value != "unknown" {
				file.WriteString(fmt.Sprintf("%v -> %v [label=\"%v\"];\n", flowEdge.Output, flowEdge.Input, flowEdge.Variable.Value))
			} else {
				file.WriteString(fmt.Sprintf("%v -> %v;\n", flowEdge.Output, flowEdge.Input))
			}
		}
	}

	file.WriteString("}")
	return nil
}

var environment *structuresWasma.Environment

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
			walk(start, subCfg.Tree, make(map[uint32]bool))
		}

		tree := GetFlowTree(environment)

		// remove duplicates
		edges := make(map[string]bool)

		for i, _ := range tree {
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

	// add output for locals
	for _, local := range environment.Locals {
		output[local.PrimaryVariableIdx] = append(output[local.PrimaryVariableIdx], local.VariableName)
	}

	// add output for globals
	for _, global := range environment.Globals {
		output[global.PrimaryVariableIdx] = append(output[global.PrimaryVariableIdx], global.VariableName)
	}

	for primaryVariableIdx, _ := range variables {
		if varOuts, found := output[primaryVariableIdx]; found {
			if varIns, found := input[primaryVariableIdx]; found {
				for _, varIn := range varIns {
					for _, varOut := range varOuts {
						tree[primaryVariableIdx] = append(
							tree[primaryVariableIdx],
							FlowEdge{variables[primaryVariableIdx],
								varOut,
								varIn})
					}
				}
			} else {
				for _, varOut := range varOuts {
					tree[primaryVariableIdx] = append(
						tree[primaryVariableIdx],
						FlowEdge{variables[primaryVariableIdx],
							varOut,
							"return"})
				}
			}
		}
	}

	return tree
}

func walk(node *controlFlowGraph.CFGNode, tree map[uint32]*controlFlowGraph.CFGNode, visited map[uint32]bool) {
	if _, found := visited[node.InstrIdx]; !found {

		if node.Control {
			node.Instruction.Executor()(node.InstrIdx, node.Instruction, environment)
		} else {
			var instrIdxs []uint32
			for instrIdx, _ := range node.Block {
				instrIdxs = append(instrIdxs, instrIdx)
			}
			sort.Slice(instrIdxs, func(i, j int) bool {
				return instrIdxs[i] < instrIdxs[j]
			})

			for _, instrIdx := range instrIdxs {
				instr, _ := node.Block[instrIdx]
				instr.Executor()(instrIdx, instr, environment)
			}
		}

		if len(node.Successors) == 1 {
			visited[node.InstrIdx] = true
			walk(tree[node.Successors[0].TargetNode], tree, visited)
		} else if len(node.Successors) > 1 {
			stack := environment.Stack
			for _, successor := range node.Successors {
				environment.Stack = stack
				visited[node.InstrIdx] = true
				walk(tree[successor.TargetNode], tree, visited)
			}
		}

	}
}

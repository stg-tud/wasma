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
	structuresWasmp "wasma/pkg/wasmp/structures"
)

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
		if local.Tainted {
			color = "style=filled, fillcolor=red,"
		}

		if local.VariableType == "P" {
			file.WriteString(fmt.Sprintf("%v [shape=box, "+color+" label=\"#%v: (param %v, tainted %v)\"];\n", local.VariableName, dfg.FuncIdx, local.VariableName, local.Tainted))
		} else if local.VariableType == "L" {
			file.WriteString(fmt.Sprintf("%v [shape=box, "+color+" label=\"#%v: (local %v)\"];\n", local.VariableName, dfg.FuncIdx, local.VariableName))
		}
	}

	for _, global := range dfg.Environment.Globals {
		color := ""
		if global.Tainted {
			color = "style=filled, fillcolor=red,"
		}
		file.WriteString(fmt.Sprintf("%v [shape=box, "+color+" label=\"#%v: (global %v)\"];\n", global.VariableName, dfg.FuncIdx, global.VariableName))
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

func NewDataFlowGraph(module *modules.Module, complete bool, funcIdx uint32) map[uint32]*DFG {
	// key: FuncIdx
	dfg := make(map[uint32]*DFG)
	cfg := controlFlowGraph.NewControlFlowGraph(module, complete, funcIdx)

	for funcIdx, subCfg := range cfg {
		environment = structuresWasma.NewEnvironment(module)

		functionType, err := code.GetFuncParams(funcIdx, module)
		if err == nil {
			// todo taint param
			for _, param := range functionType.ParameterTypes {
				environment.NewParameterWithTaint(param, true)
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
		} else {
			log.Fatalf("no root node found for function: %v", funcIdx)
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

	for i := 0; i < 1000; i++ {
		// propagate taint

		for _, dataFlowEdge := range environment.Flow {
			// variable -> instruction
			for _, varIn := range dataFlowEdge.Input {
				if varIn.Tainted {
					for _, varOut := range dataFlowEdge.Output {
						varOut.Tainted = true
						primaryVariableIdx := varOut.PrimaryVariableIdx
						variableName := varOut.VariableName
						//environment.Flow[instrIdx].Output[primaryVariableIdx] = varOut
						environment.Variables[primaryVariableIdx] = varOut
						variables[primaryVariableIdx] = varOut

						// also taint vars with same name
						for varIdx0, varOut0 := range environment.Variables {
							if varOut0.PrimaryVariableIdx == primaryVariableIdx || varOut0.VariableName == variableName {
								varOut0.Tainted = true
								environment.Variables[varIdx0] = varOut0
							}
						}

						// also taint vars with same name
						for varIdx00, varOut00 := range variables {
							if varOut00.PrimaryVariableIdx == primaryVariableIdx || varOut00.VariableName == variableName {
								varOut00.Tainted = true
								variables[varIdx00] = varOut00
							}
						}

						for instrIdx2, dataFlowEdge := range environment.Flow {
							// instruction -> variable
							for varIdx2, varOut2 := range dataFlowEdge.Input {
								if varOut2.PrimaryVariableIdx == primaryVariableIdx || varOut2.VariableName == variableName {
									varOut2.Tainted = true
									environment.Flow[instrIdx2].Input[varIdx2] = varOut2
								}
							}
							for varIdx3, varOut3 := range dataFlowEdge.Output {
								if varOut3.PrimaryVariableIdx == primaryVariableIdx || varOut3.VariableName == variableName {
									varOut3.Tainted = true
									environment.Flow[instrIdx2].Output[varIdx3] = varOut3
								}
							}
						}

					}
				}
			}
		}

		// propagate taint
		/*
			for _, dataFlowEdge := range environment.Flow {
				// instruction -> variable
				for _, varOut := range dataFlowEdge.Output {
					if varOut.Tainted {
						for _, varIn := range dataFlowEdge.Input {

							varIn.Tainted = true
							primaryVariableIdx := varIn.PrimaryVariableIdx
							//environment.Flow[instrIdx].Output[varIdx] = varIn
							environment.Variables[primaryVariableIdx] = varIn
							variables[primaryVariableIdx] = varIn

							for instrIdx2, dataFlowEdge := range environment.Flow {
								// instruction -> variable
								for varIdx2, varOut2 := range dataFlowEdge.Input {
									if varOut2.PrimaryVariableIdx == primaryVariableIdx {
										environment.Flow[instrIdx2].Input[varIdx2] = varIn
									}
								}
								for varIdx3, varOut3 := range dataFlowEdge.Output {
									if varOut3.PrimaryVariableIdx == primaryVariableIdx {
										environment.Flow[instrIdx2].Output[varIdx3] = varIn
									}
								}
							}
						}
					}
				}
			}
		*/
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
						tainted := false
						if variables[primaryVariableIdx].Tainted {
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
					if variables[primaryVariableIdx].Tainted {
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

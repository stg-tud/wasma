package callGraph

import (
	"errors"
	"fmt"
	"log"
	"os"
	"wasma/pkg/wasmp/instructions"
	"wasma/pkg/wasmp/modules"
)

type Function struct {
	FuncIdx    uint32
	IsIndirect bool
}

type CallGraphEdges struct {
	// func x() {
	//     calls y
	// }
	IsCalled    []Function // y <- x
	Calls       []Function // x -> y
	NumIndirect uint32     // x -> ?
}

type EntryPoint struct {
	FuncIdx     uint32
	name        string
	numIndirect uint32
}
type CallGraph struct {
	EntryPoints []EntryPoint
	edges       map[uint32]*CallGraphEdges
	module      *modules.Module
}

func (callGraph *CallGraph) GetEdges() map[uint32]*CallGraphEdges {
	return callGraph.edges
}

// Checks if a given FuncIdx is in a given list of entry points.
func (callGraph *CallGraph) IsEntryPoint(funcIdx uint32) bool {
	for _, entryPoint := range callGraph.EntryPoints {
		if funcIdx == entryPoint.FuncIdx {
			return true
		}
	}
	return false
}

// Saves the current call graph to a file in the dot format.
func (callGraph *CallGraph) SaveDot(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("digraph G {\n")

	for _, entryPoint := range callGraph.EntryPoints {
		file.WriteString(fmt.Sprintf("%v [style=filled, color=\"grey\", label=\"%v[FuncIdx: %v, numIndirect: %v]\"];\n",
			entryPoint.FuncIdx,
			entryPoint.name,
			entryPoint.FuncIdx,
			entryPoint.numIndirect))
	}

	for funcIdx, edges := range callGraph.edges {
		if edges.NumIndirect > 0 && !callGraph.IsEntryPoint(funcIdx) {
			file.WriteString(fmt.Sprintf("%v [color=\"red\", label=\"%v(%v)\"];\n", funcIdx, funcIdx, edges.NumIndirect))
		}
		for _, call := range edges.Calls {
			if call.IsIndirect {
				file.WriteString(fmt.Sprintf("%v -> %v [label=\"indirect\", color=red];\n", funcIdx, call.FuncIdx))
			} else {
				file.WriteString(fmt.Sprintf("%v -> %v;\n", funcIdx, call.FuncIdx))
			}

		}
	}
	file.WriteString("}")

	return nil
}

// Returns a list of functions that are called by the
// given function.
// The second return value is true if the list of called
// functions not is empty, otherwise false.
func (callGraph *CallGraph) GetChildren(function uint32) ([]Function, uint32, bool) {
	if edges, found := callGraph.edges[function]; found {
		return edges.Calls, edges.NumIndirect, found
	} else {
		return []Function{}, 0, found
	}
}

// Returns a list of functions that call the given function.
// The second return value is true if the list of callers
// not is empty, otherwise false.
func (callGraph *CallGraph) GetParents(function uint32) ([]Function, bool) {
	if edges, found := callGraph.edges[function]; found {
		return edges.IsCalled, found
	} else {
		return []Function{}, found
	}
}

// Prints the call graph on the console.
func (callGraph *CallGraph) Print() error {
	var printedNodesCounter = 0
	for _, entryPoint := range callGraph.EntryPoints {
		layer := "    |--"
		log.Printf("ENTRYPOINT = (func: %v, name: %v)", entryPoint.FuncIdx, entryPoint.name)
		if children, indirect, found := callGraph.GetChildren(entryPoint.FuncIdx); found {
			log.Printf("%v> ?: %v", layer, indirect)
			for _, child := range children {
				printedNodesCounter++
				log.Printf("%v> %v (indirect: %v)", layer, child.FuncIdx, child.IsIndirect)
				if err := callGraph.print(child.FuncIdx, []uint32{child.FuncIdx}, "     "+layer, &printedNodesCounter); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Prints child of the call graph node.
func (callGraph *CallGraph) print(currentChild uint32, visited []uint32, layer string, printedNodesCounter *int) error {
	if *printedNodesCounter > 1000 {
		return errors.New("stopping output: The call graph is to big for console output")
	}

	if children, indirect, found := callGraph.GetChildren(currentChild); found {
		log.Printf("%v> ?: %v", layer, indirect)
		for _, child := range children {
			*printedNodesCounter++
			if exists := inElems(child.FuncIdx, visited); exists {
				log.Printf("%v> %v (cycle -> %v, indirect: %v)", layer, child.FuncIdx, child.FuncIdx, child.IsIndirect)
				continue
			} else {
				log.Printf("%v> %v (indirect: %v)", layer, child.FuncIdx, child.IsIndirect)
			}
			if err := callGraph.print(child.FuncIdx, append(visited, child.FuncIdx), "     "+layer, printedNodesCounter); err != nil {
				return err
			}
		}
	}
	return nil
}

func inFuncs(function Function, functions []Function) bool {
	for _, f := range functions {
		if f.FuncIdx == function.FuncIdx && f.IsIndirect == function.IsIndirect {
			return true
		}
	}

	return false
}

// Checks whether a given element is in a list.
func inElems(element uint32, list []uint32) bool {
	for _, listElement := range list {
		if element == listElement {
			return true
		}
	}
	return false
}

func addCallToCallee(callee uint32, caller uint32, callGraphEdges *map[uint32]*CallGraphEdges, isIndirect bool) {
	function := Function{caller, isIndirect}
	if edges, found := (*callGraphEdges)[callee]; found {
		if !inFuncs(function, edges.IsCalled) {
			edges.IsCalled = append(edges.IsCalled, function)
			(*callGraphEdges)[callee] = edges
		}

	} else {
		(*callGraphEdges)[callee] = &CallGraphEdges{
			[]Function{function},
			[]Function{},
			0}
	}
}

func addCallToCaller(caller uint32, callee uint32, callGraphEdges *map[uint32]*CallGraphEdges, isIndirect bool) {
	function := Function{callee, isIndirect}
	if edges, found := (*callGraphEdges)[caller]; found {
		if !inFuncs(function, edges.Calls) {
			edges.Calls = append(edges.Calls, function)
			(*callGraphEdges)[caller] = edges
		}

	} else {
		(*callGraphEdges)[caller] = &CallGraphEdges{
			[]Function{},
			[]Function{function},
			0}
	}
}

func addCallIndirectToCaller(caller uint32, callGraphEdges *map[uint32]*CallGraphEdges) {
	edges := (*callGraphEdges)[caller]
	if edges != nil {
		edges.NumIndirect = edges.NumIndirect + 1
		(*callGraphEdges)[caller] = edges
	} else {
		(*callGraphEdges)[caller] = &CallGraphEdges{[]Function{}, []Function{}, 1}
	}
}

// Searches for all call instructions (call, call_indirect) in a function body
func findCalls(
	callerIdx uint32,
	instructions []instructions.Instruction,
	callGraphEdges *map[uint32]*CallGraphEdges,
	module *modules.Module,
	findIndirectCalls bool) {
	for _, instruction := range instructions {

		// check if an instruction has sub instructions, for example, 'block'
		if subInstructions, err := instruction.Instr(); err == nil {
			findCalls(callerIdx, subInstructions, callGraphEdges, module, findIndirectCalls)
		}

		// check if an instruction has an else block
		if subInstructionsElse, err := instruction.ElseInstr(); err == nil {
			findCalls(callerIdx, subInstructionsElse, callGraphEdges, module, findIndirectCalls)
		}

		// call instruction
		if instruction.Name() == "call" {
			calleeIdx, err := instruction.Funcidx()
			if err != nil {
				log.Fatal(err.Error())
			}
			addCallToCaller(callerIdx, calleeIdx, callGraphEdges, false)
			addCallToCallee(calleeIdx, callerIdx, callGraphEdges, false)

			// call_indirect instruction
		} else if instruction.Name() == "call_indirect" && findIndirectCalls {
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
							addCallToCaller(callerIdx, funcIdx, callGraphEdges, true)
							addCallToCallee(funcIdx, callerIdx, callGraphEdges, true)
						}
					}
				}
			}

			addCallIndirectToCaller(callerIdx, callGraphEdges)
		}
	}
}

// Returns a new call graph of the given module
func NewCallGraph(module *modules.Module, findIndirectCalls bool) (*CallGraph, error) {
	callGraphEdges := make(map[uint32]*CallGraphEdges)
	var entryPoints []EntryPoint

	// create call graph
	if codeSection, err := module.GetCodeSection(); err == nil {
		for callerFunctionIdx, code := range codeSection.Codes {
			findCalls(callerFunctionIdx, code.Function.Expr.Instructions, &callGraphEdges, module, findIndirectCalls)
		}
	} else {
		return nil, errors.New("no call graph could be generated because the module contains no code")
	}

	// read main entry point if exists
	if startSection, err := module.GetStartSection(); err == nil {
		if edge, found := callGraphEdges[startSection.FuncIdx]; found {
			entryPoints = append(entryPoints, EntryPoint{startSection.FuncIdx, "start section", edge.NumIndirect})
		} else {
			entryPoints = append(entryPoints, EntryPoint{startSection.FuncIdx, "start section", 0})
		}
	}

	// read functionIdxs of export functions as additional entry points if exist
	if exportSection, err := module.GetExportSection(); err == nil {
		for _, export := range exportSection.Exports {
			// 0x00 => FuncIdx
			if export.ExportDesc.ExportType == 0x00 {
				if edge, found := callGraphEdges[export.ExportDesc.FuncIdx]; found {
					entryPoints = append(entryPoints, EntryPoint{export.ExportDesc.FuncIdx, export.Name, edge.NumIndirect})
				} else {
					entryPoints = append(entryPoints, EntryPoint{export.ExportDesc.FuncIdx, export.Name, 0})
				}
			}
		}
	}

	return &CallGraph{entryPoints, callGraphEdges, module}, nil
}

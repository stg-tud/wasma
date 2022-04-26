package controlFlowGraph

import (
	"errors"
	"fmt"
	"os"
	"wasma/pkg/wasma/code"
	"wasma/pkg/wasmp/instructions"
	"wasma/pkg/wasmp/modules"
)

type Edge struct {
	// instrIdx
	TargetNode uint32
	Tag        string
}
type CFGNode struct {
	Control     bool
	Name        string
	InstrIdx    uint32
	Instruction instructions.Instruction
	FuncOffset  uint32
	Block       map[uint32]instructions.Instruction
	Successors  []Edge
}

func NewTree(disassembly code.Disassembly) map[uint32]*CFGNode {
	// key: instrIdx of the block's first instruction
	cfg := make(map[uint32]*CFGNode)
	// key: instrIdx
	newBlockAt := make(map[uint32]bool)
	// key: start else node, value: end else node
	startElseNode := make(map[uint32]uint32)
	// key: then end instrIdx
	thenEnd := make(map[uint32]bool)

	//var successor uint32 = 0
	var instrIdx uint32 = 0
	var block *CFGNode = nil
	var blockCounter = 0

	for ; instrIdx < uint32(len(disassembly.DisassembledInstrs)); instrIdx++ {
		currentInstr := disassembly.DisassembledInstrs[instrIdx]

		switch currentInstr.Instruction.Name() {
		// handle all control instructions
		case "block", "loop", "if", "br", "br_if", "br_table", "call", "call_indirect", "return", "unreachable":
			{
				// create a new basic block after every WebAssembly block was closed
				if currentInstr.Instruction.Name() == "block" {
					if block, found := disassembly.Blocks[instrIdx]; found {
						newBlockAt[block.Continuation] = true
					}
				}
				// link last basic block with current control node
				if block != nil {
					if _, found := startElseNode[instrIdx]; !found {
						block.Successors = append(block.Successors, Edge{instrIdx, ""})
					}
					cfg[block.FuncOffset] = block
					block = nil
				}

				// handle return
				if currentInstr.Instruction.Name() == "return" {
					cfg[instrIdx] = &CFGNode{true, "return", instrIdx, currentInstr.Instruction, instrIdx, nil, nil}
				} else if currentInstr.Instruction.Name() == "unreachable" {
					cfg[instrIdx] = &CFGNode{true, "unreachable", instrIdx, currentInstr.Instruction, instrIdx, nil, nil}
				} else {
					// create current control node
					cfgControl := new(CFGNode)
					cfgControl.Control = true
					cfgControl.FuncOffset = instrIdx
					cfgControl.Name = currentInstr.Instruction.Name()
					cfgControl.Instruction = currentInstr.Instruction
					cfgControl.InstrIdx = instrIdx

					// handle if instruction
					if currentInstr.Instruction.Name() == "if" {
						if ifBlock, found := disassembly.IfBlocks[instrIdx]; found {
							if ifBlock.ThenEnd == ifBlock.ElseEnd {
								// only then block
								if ifBlock.ThenEnd+1 < uint32(len(disassembly.DisassembledInstrs)) {
									if _, found := startElseNode[ifBlock.ThenEnd+1]; found {
										if startElseNode[ifBlock.ThenEnd+1]+1 < uint32(len(disassembly.DisassembledInstrs)) {
											cfgControl.Successors = append(cfgControl.Successors, Edge{startElseNode[ifBlock.ThenEnd+1] + 1, ""})
										}
									} else {
										cfgControl.Successors = append(cfgControl.Successors, Edge{ifBlock.ThenEnd + 1, ""})
									}
								}
								if instrIdx+1 < uint32(len(disassembly.DisassembledInstrs)) {
									cfgControl.Successors = append(cfgControl.Successors, Edge{instrIdx + 1, "then"})
								}
								if ifBlock.ThenEnd+1 < uint32(len(disassembly.DisassembledInstrs)) {
									newBlockAt[ifBlock.ThenEnd+1] = true
								}
							} else {
								startElseNode[ifBlock.ThenEnd+1] = ifBlock.ElseEnd
								thenEnd[ifBlock.ThenEnd] = true
								newBlockAt[ifBlock.ElseEnd+1] = true
								if ifBlock.ThenEnd == instrIdx {
									// only else block
									if ifBlock.ElseEnd+1 < uint32(len(disassembly.DisassembledInstrs)) {
										if _, found := startElseNode[ifBlock.ElseEnd+1]; found {
											if startElseNode[ifBlock.ElseEnd+1]+1 < uint32(len(disassembly.DisassembledInstrs)) {
												cfgControl.Successors = append(cfgControl.Successors, Edge{startElseNode[ifBlock.ElseEnd+1] + 1, ""})
											}
										} else {
											cfgControl.Successors = append(cfgControl.Successors, Edge{ifBlock.ElseEnd + 1, ""})
										}
									}
									if instrIdx+1 < uint32(len(disassembly.DisassembledInstrs)) {
										cfgControl.Successors = append(cfgControl.Successors, Edge{instrIdx + 1, "else"})
									}
								} else {
									// then and else block
									if instrIdx+1 < uint32(len(disassembly.DisassembledInstrs)) {
										cfgControl.Successors = append(cfgControl.Successors, Edge{instrIdx + 1, "then"})
									}
									if ifBlock.ThenEnd+1 < uint32(len(disassembly.DisassembledInstrs)) {
										cfgControl.Successors = append(cfgControl.Successors, Edge{ifBlock.ThenEnd + 1, "else"})
									}
								}
							}
						}
					} else if currentInstr.Instruction.Name() == "br_table" {
						// This map is used to group the targets of a br_table instruction.
						continuations := make(map[uint32]bool)

						// set graph edges for branch instruction brTable
						if targetList, defaultTarget, err := disassembly.GetBrTableTargets(instrIdx); err == nil {
							// default case of br_table
							if continuation, err := disassembly.GetContinuation(defaultTarget); err == nil {
								if continuation < uint32(len(disassembly.DisassembledInstrs)) {
									edge := Edge{continuation, ""}
									cfgControl.Successors = append(cfgControl.Successors, edge)
									continuations[continuation] = true
								}
							}
							for _, target := range targetList {
								if continuation, err := disassembly.GetContinuation(target); err == nil {
									if _, found := continuations[continuation]; !found {
										if continuation < uint32(len(disassembly.DisassembledInstrs)) {
											edge := Edge{continuation, ""}
											cfgControl.Successors = append(cfgControl.Successors, edge)
											continuations[continuation] = true
										}
									}
								}
							}
						}
					} else if currentInstr.Instruction.Name() == "br" ||
						currentInstr.Instruction.Name() == "br_if" {

						// set graph edges for branch instructions br and brIf
						if targetLabelInstructionIdx, err := disassembly.GetBranchTarget(instrIdx); err == nil {
							if continuation, err := disassembly.GetContinuation(targetLabelInstructionIdx); err == nil {
								if continuation < uint32(len(disassembly.DisassembledInstrs)) {
									edge := Edge{continuation, ""}
									cfgControl.Successors = append(cfgControl.Successors, edge)
								}
							}
						}

						// Add an edge for the next instruction for the case that
						// the branch condition is not fulfilled.
						if currentInstr.Instruction.Name() == "br_if" {
							// link then branch to successor of if
							if successor, found := getIfSuccessor(instrIdx, thenEnd, startElseNode, disassembly); found {
								cfgControl.Successors = append(cfgControl.Successors, successor)
							} else if instrIdx+1 < uint32(len(disassembly.DisassembledInstrs)) { // Check if there is a next instruction.
								cfgControl.Successors = append(cfgControl.Successors, Edge{instrIdx + 1, ""})
							}
						}
					} else {
						// check if there is a next instruction
						if instrIdx+1 < uint32(len(disassembly.DisassembledInstrs)) {
							if _, found := startElseNode[instrIdx+1]; found {
								if startElseNode[instrIdx+1]+1 < uint32(len(disassembly.DisassembledInstrs)) {
									cfgControl.Successors = append(cfgControl.Successors, Edge{startElseNode[instrIdx+1] + 1, ""})
								}
							} else {
								cfgControl.Successors = append(cfgControl.Successors, Edge{instrIdx + 1, ""})
							}
						}

					}

					cfg[instrIdx] = cfgControl
				}
			}
		default:
			{
				if _, found := startElseNode[instrIdx]; found && block != nil {
					// if then block is finished store last block and go on
					cfg[block.FuncOffset] = block
					block = nil

				}

				// check if a new block has to be started
				if _, found := newBlockAt[instrIdx]; found {
					if block != nil {
						block.Successors = append(block.Successors, Edge{instrIdx, ""})
						cfg[block.FuncOffset] = block
						block = nil

					}
				}

				// basic blocks
				if block == nil {
					block = new(CFGNode)
					block.Control = false
					block.FuncOffset = instrIdx
					block.Name = fmt.Sprintf("Basic Block %v", blockCounter)
					block.InstrIdx = instrIdx
					block.Instruction = currentInstr.Instruction
					blockCounter++
					block.Block = make(map[uint32]instructions.Instruction)
					block.Block[instrIdx] = currentInstr.Instruction
				} else {
					block.Block[instrIdx] = currentInstr.Instruction
				}

				// link then branch to successor of if
				if successor, found := getIfSuccessor(instrIdx, thenEnd, startElseNode, disassembly); found {
					block.Successors = append(block.Successors, successor)
				}
			}
		}
	}

	if block != nil {
		cfg[block.InstrIdx] = block
	}

	return cfg
}

func getIfSuccessor(instrIdx uint32, thenEnd map[uint32]bool, startElseNode map[uint32]uint32, disassembly code.Disassembly) (Edge, bool) {
	// link then branch to successor of if
	if _, found := thenEnd[instrIdx]; found {
		if elseEnd, found := startElseNode[instrIdx+1]; found {
			for {
				// check if the current else branch is limited by an outer branch
				var outerElseEnd uint32
				if outerElseEnd, found = startElseNode[elseEnd+1]; !found && elseEnd+1 < uint32(len(disassembly.DisassembledInstrs)) {
					return Edge{elseEnd + 1, ""}, true
					break
				} else if !found && elseEnd+1 >= uint32(len(disassembly.DisassembledInstrs)) {
					break
				}
				elseEnd = outerElseEnd
			}

		}
	}
	return Edge{}, false
}

// ControlFlowGraph
type CFG struct {
	// key: instrIdx
	Tree        map[uint32]*CFGNode
	Disassembly code.Disassembly
}

func NewControlFlowGraph(module *modules.Module, complete bool, fIdx uint32) map[uint32]*CFG {
	// key: FuncIdx
	cfg := make(map[uint32]*CFG)
	if complete {
		if functionSection, err := module.GetFunctionSection(); err == nil {
			for funcIdx, _ := range functionSection.TypeIdxs {
				disassembly := code.DisassemblyFunction(funcIdx, module)
				cfg[funcIdx] = &CFG{NewTree(disassembly), disassembly}
			}
		}
	} else {
		disassembly := code.DisassemblyFunction(fIdx, module)
		cfg[fIdx] = &CFG{NewTree(disassembly), disassembly}
	}
	return cfg
}

func In(element uint32, list []uint32) bool {
	for _, listElement := range list {
		if element == listElement {
			return true
		}
	}
	return false
}

func (cfg *CFG) GetUnreachableCode() ([]uint32, error) {
	var visitedNodes []uint32
	var unreachable []uint32

	if firstCfgNode, found := cfg.Tree[0]; found {
		visitedNodes = append(visitedNodes, 0)
		var queue []Edge
		queue = append(queue, firstCfgNode.Successors...)

		for len(queue) > 0 {
			head, tail := queue[0], queue[1:]
			queue = tail
			visitedNodes = append(visitedNodes, head.TargetNode)

			if cfgNode, found := cfg.Tree[head.TargetNode]; found {
				queue = append(queue, cfgNode.Successors...)
			} else {
				return nil, errors.New(fmt.Sprintf("no CFG node with index %v found", head.TargetNode))
			}
		}
	} else {
		return nil, errors.New("no first CFG node with index 0 found")
	}

	for instrIdx, _ := range cfg.Tree {
		if !In(instrIdx, visitedNodes) {
			unreachable = append(unreachable, instrIdx)
		}
	}

	return unreachable, nil
}

// Saves the control flow graph for a given FuncIdx to a file in the dot format.
func (cfg *CFG) SaveDot(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("digraph G {\n")

	var i uint32 = 0
	for ; i < uint32(len(cfg.Disassembly.DisassembledInstrs)); i++ {
		if edge, found := cfg.Tree[i]; found {
			if edge.Control {
				if label, found := cfg.Disassembly.LabelsI[i]; found {
					file.WriteString(fmt.Sprintf("%v [label=\"#%v+%v: %v [L%v, next=%v]\"];\n", i, 1, i, edge.Name, label.LabelIdx, label.NextInstrIdx))
				} else {
					instr := cfg.Disassembly.DisassembledInstrs[i].Instruction
					if instr.Name() == "if" {
						file.WriteString(fmt.Sprintf("%v [shape=diamond, label=\"#%v+%v: %v\"];\n", i, 1, i, edge.Name))
					} else if instr.Name() == "call" {
						if funcIdx, err := instr.Funcidx(); err == nil {
							file.WriteString(fmt.Sprintf("%v [shape=cds, label=\"#%v+%v: %v [FuncIdx=%v]\"];\n", i, 1, i, edge.Name, funcIdx))
						} else {
							file.WriteString(fmt.Sprintf("%v [shape=cds, label=\"#%v+%v: %v\"];\n", i, 1, i, edge.Name))
						}
					} else if instr.Name() == "call_indirect" {
						if typeIdx, err := instr.Typeidx(); err == nil {
							file.WriteString(fmt.Sprintf("%v [shape=cds, label=\"#%v+%v: %v [typeIdx=%v]\"];\n", i, 1, i, edge.Name, typeIdx))
						} else {
							file.WriteString(fmt.Sprintf("%v [shape=cds, label=\"#%v+%v: %v\"];\n", i, 1, i, edge.Name))
						}
					} else {
						file.WriteString(fmt.Sprintf("%v [label=\"#%v+%v: %v\"];\n", i, 1, i, edge.Name))
					}
				}
			} else {
				var instructions string
				var j uint32 = 0
				for ; j < uint32(len(cfg.Disassembly.DisassembledInstrs)); j++ {
					if instr, found := edge.Block[j]; found {
						var value = ""

						switch instr.Name() {
						case "local.get", "local.set", "local.tee":
							{
								localidx, _ := instr.Localidx()
								value = fmt.Sprintf("%v", localidx)
							}
						case "global.get", "global.set":
							{
								globalidx, _ := instr.Globalidx()
								value = fmt.Sprintf("%v", globalidx)
							}
						case "i32.const":
							{
								i32, _ := instr.I32()
								value = fmt.Sprintf("%v", i32)
							}
						case "i64.const":
							{
								i64, _ := instr.I64()
								value = fmt.Sprintf("%v", i64)
							}
						case "f32.const":
							{
								f32, _ := instr.F32()
								value = fmt.Sprintf("%v", f32)
							}
						case "f64.const":
							{
								f64, _ := instr.F64()
								value = fmt.Sprintf("%v", f64)
							}
						}

						instructions = instructions + fmt.Sprintf("#%v+%v: %v %v\\l", 1, j, instr.Name(), value)
					}
				}
				file.WriteString(fmt.Sprintf("%v [shape=box, label=\"%v\n%v\"];\n", i, edge.Name, instructions))
			}
			for _, successor := range edge.Successors {
				file.WriteString(fmt.Sprintf("%v -> %v [label=\"%v\"];\n", i, successor.TargetNode, successor.Tag))
			}
		}
	}

	file.WriteString("}")
	return nil
}

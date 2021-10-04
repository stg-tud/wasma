package code

import (
	"errors"
	"fmt"
	"log"
	"sort"
	instructions2 "wasma/pkg/wasmp/instructions"
	"wasma/pkg/wasmp/modules"
	"wasma/pkg/wasmp/modules/sections"
	"wasma/pkg/wasmp/types"
)

type Loop struct {
	End          uint32
	Continuation uint32
}

type Block struct {
	End          uint32
	Continuation uint32
}

type IfBlock struct {
	ThenEnd      uint32
	ElseEnd      uint32
	Continuation uint32
}
type LabelL struct {
	InstrIdx     uint32
	NextInstrIdx uint32
}
type LabelI struct {
	LabelIdx     uint32
	NextInstrIdx uint32
}

type InstrDisassembly struct {
	Instruction instructions2.Instruction
	offset      uint32
}

type Disassembly struct {
	funcIdx uint32
	// key: instrIdx
	DisassembledInstrs map[uint32]InstrDisassembly
	// key: labelIdx
	LabelsL map[uint32]LabelL
	// key: instrIdx
	LabelsI map[uint32]LabelI
	// key: instrIdx of block instruction
	Blocks map[uint32]Block
	// key: instrIdx of loop instruction
	Loops map[uint32]Loop
	// key: instrIdx of if instruction
	IfBlocks map[uint32]IfBlock
	// key: instrIdx of branch instruction
	BranchJumps map[uint32][]uint32
}

func (disassembly *Disassembly) GetLabel(labelInstruction uint32) (uint32, bool) {
	label, found := disassembly.LabelsI[labelInstruction]
	return label.LabelIdx, found
}

func (disassembly *Disassembly) GetContinuation(labelInstruction uint32) (uint32, error) {
	if instrDisassembly, found := disassembly.DisassembledInstrs[labelInstruction]; found {
		if instrDisassembly.Instruction.Name() == "block" {
			if block, found := disassembly.Blocks[labelInstruction]; found {
				return block.Continuation, nil
			} else {
				return 0, errors.New(fmt.Sprintf("no block with instruction index=%v found", labelInstruction))
			}
		} else if instrDisassembly.Instruction.Name() == "loop" {
			if loop, found := disassembly.Loops[labelInstruction]; found {
				return loop.Continuation, nil
			} else {
				return 0, errors.New(fmt.Sprintf("no loop with instruction index=%v found", labelInstruction))
			}
		} else if instrDisassembly.Instruction.Name() == "if" {
			if ifBlock, found := disassembly.IfBlocks[labelInstruction]; found {
				return ifBlock.Continuation, nil
			} else {
				return 0, errors.New(fmt.Sprintf("no if block with instruction index=%v found", labelInstruction))
			}
		} else {
			return 0, errors.New(fmt.Sprintf("expected a block, loop or if instruction, but got: %v", instrDisassembly.Instruction.Name()))
		}
	}
	return 0, errors.New(fmt.Sprintf("%v is no valid instruction index", labelInstruction))
}

// Returns the instrIdx of a branch instruction's (br, brIf) target instruction
func (disassembly *Disassembly) GetBranchTarget(instrIdx uint32) (uint32, error) {
	if instrDisassembly, found := disassembly.DisassembledInstrs[instrIdx]; found &&
		(instrDisassembly.Instruction.Name() == "br" ||
			instrDisassembly.Instruction.Name() == "br_if") {
		if labelIdx, err := instrDisassembly.Instruction.Labelidx(); err == nil {
			return disassembly.getBranchTarget(instrIdx, labelIdx)
		} else {
			return 0, errors.New(fmt.Sprintf("instrIdx: %v is no branch instruction (br or brIf)", instrIdx))
		}
	} else {
		return 0, errors.New(fmt.Sprintf("no br or brIf instruction with given instrIdx: %v", instrIdx))
	}
}

func (disassembly *Disassembly) getBranchTarget(instrIdx uint32, labelIdx uint32) (uint32, error) {
	var labelInstruction uint32
	if stack, found := disassembly.BranchJumps[instrIdx]; found {
		labelIdx++ // plus one since the label index start with zero
		// labelIdx=2 => 0->1->2
		if uint32(len(stack)) >= labelIdx {
			for ; labelIdx > 0; labelIdx-- {
				labelInstruction = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		} else {
			return 0, errors.New(fmt.Sprintf("too few elements on the label stack for instrIdx: %v", instrIdx))
		}
	} else {
		return 0, errors.New(fmt.Sprintf("no label stack found for instrIdx: %v", instrIdx))
	}
	return labelInstruction, nil
}

// Returns the instrIdxs of a branch table instruction's target instructions
func (disassembly *Disassembly) GetBrTableTargets(instrIdx uint32) ([]uint32, uint32, error) {
	var defaultTarget uint32 = 0
	var targetList []uint32

	if instrDisassembly, found := disassembly.DisassembledInstrs[instrIdx]; found &&
		instrDisassembly.Instruction.Name() == "br_table" {
		// target list
		if vecLabelIdx, err := instrDisassembly.Instruction.VecLabelidx(); err == nil {
			for _, labelIdx := range vecLabelIdx {
				if target, err := disassembly.getBranchTarget(instrIdx, labelIdx); err == nil {
					targetList = append(targetList, target)
				} else {
					return nil, 0, err
				}
			}
		} else {
			return nil, 0, errors.New(fmt.Sprintf("labelIdx vector does not exist for instrIdx: %v", instrIdx))
		}
		// default branch
		if labelIdx, err := instrDisassembly.Instruction.Labelidx(); err == nil {
			if defTar, err := disassembly.getBranchTarget(instrIdx, labelIdx); err == nil {
				defaultTarget = defTar
			} else {
				return nil, 0, err
			}
		} else {
			return nil, 0, errors.New(fmt.Sprintf("instrIdx: %v is no branch table instruction", instrIdx))
		}
	} else {
		return nil, 0, errors.New(fmt.Sprintf("no br or brIf instruction with given instrIdx: %v", instrIdx))
	}
	return targetList, defaultTarget, nil
}

// Returns a text representation of a function's disassembly.
func (disassembly *Disassembly) GetTextRepresentation() []string {
	var textRepresentation []string
	structTags := make(map[int][]string)
	//structTagsBlock := make(map[int][]string)

	for i := 0; int(i) < len(disassembly.DisassembledInstrs); i++ {
		instr := disassembly.DisassembledInstrs[uint32(i)]
		if label, found := disassembly.LabelsI[uint32(i)]; found && !(instr.Instruction.Name() == "if") {
			if block, found := disassembly.Blocks[uint32(i)]; found {
				textRepresentation = append(textRepresentation,
					fmt.Sprintf("%v+%v: %v [label=L%v, BlockEnd=%v, Continuation=%v, file offset=%v]",
						disassembly.funcIdx,
						i,
						instr.Instruction.Name(),
						label.LabelIdx,
						block.End,
						block.Continuation,
						instr.offset))
				structTags[int(block.End)] = append(structTags[int(block.End)], fmt.Sprintf("%v->: end block", disassembly.funcIdx))
			} else if loop, found := disassembly.Loops[uint32(i)]; found {
				textRepresentation = append(textRepresentation,
					fmt.Sprintf("%v+%v: %v [label=L%v, LoopEnd=%v, Continuation=%v, file offset=%v]",
						disassembly.funcIdx,
						i,
						instr.Instruction.Name(),
						label.LabelIdx,
						loop.End,
						loop.Continuation,
						instr.offset))
				structTags[int(loop.End)] = append(structTags[int(loop.End)], fmt.Sprintf("%v->: end loop", disassembly.funcIdx))
			} else {
				textRepresentation = append(textRepresentation,
					fmt.Sprintf("%v+%v: %v [label=L%v, file offset=%v]",
						disassembly.funcIdx,
						i,
						instr.Instruction.Name(),
						label.LabelIdx,
						instr.offset))
			}

		} else if ifBlock, found := disassembly.IfBlocks[uint32(i)]; found {
			textRepresentation = append(textRepresentation,
				fmt.Sprintf("%v+%v: %v [label=L%v, ThenEnd=%v, ElseEnd=%v, file offset=%v]",
					disassembly.funcIdx,
					i,
					instr.Instruction.Name(),
					label.LabelIdx,
					ifBlock.ThenEnd,
					ifBlock.ElseEnd,
					instr.offset))
			if ifBlock.ThenEnd == ifBlock.ElseEnd {
				structTags[i] = append(structTags[i], fmt.Sprintf("%v->: then", disassembly.funcIdx))
				structTags[int(ifBlock.ThenEnd)] = append(structTags[int(ifBlock.ThenEnd)], fmt.Sprintf("%v->: end if", disassembly.funcIdx))

			} else {
				structTags[i] = append(structTags[i], fmt.Sprintf("%v->: then", disassembly.funcIdx))
				structTags[int(ifBlock.ThenEnd)] = append(structTags[int(ifBlock.ThenEnd)], fmt.Sprintf("%v->: else", disassembly.funcIdx))
				structTags[int(ifBlock.ElseEnd)] = append(structTags[int(ifBlock.ElseEnd)], fmt.Sprintf("%v->: end if", disassembly.funcIdx))
			}

		} else if instr.Instruction.Name() == "br" ||
			instr.Instruction.Name() == "br_if" ||
			instr.Instruction.Name() == "br_table" {
			labelIdx, _ := instr.Instruction.Labelidx()
			var targetLabel string = "?"

			if labelInstruction, err := disassembly.GetBranchTarget(uint32(i)); err == nil {
				targetLabel = fmt.Sprintf("%v", labelInstruction)

				if label, found := disassembly.GetLabel(labelInstruction); found {
					targetLabel = fmt.Sprintf("%v", label)
				}
			}

			textRepresentation = append(textRepresentation,
				fmt.Sprintf("%v+%v: %v [jumpup=%v, target label=L%v, file offset=%v]",
					disassembly.funcIdx,
					i,
					instr.Instruction.Name(),
					labelIdx,
					targetLabel,
					instr.offset))
		} else {
			textRepresentation = append(textRepresentation,
				fmt.Sprintf("%v+%v: %v [file offset=%v]",
					disassembly.funcIdx,
					i,
					instr.Instruction.Name(),
					instr.offset))
		}

		if tags, found := structTags[i]; found {
			start := len(tags) - 1

			for ; start >= 0; start-- {
				textRepresentation = append(textRepresentation, tags[start])
			}
		}
	}
	return textRepresentation
}

// Disassembles a function body
func DisassemblyFunction(funcIdx uint32, module *modules.Module) Disassembly {
	disassembledInstrs := make(map[uint32]InstrDisassembly)
	// key: labelIdx
	labelsL := make(map[uint32]LabelL)
	// key: instrIdx
	labelsI := make(map[uint32]LabelI)
	// key: instrIdx of if instruction
	ifBlocks := make(map[uint32]IfBlock)
	// key: instrIdx of block instruction
	blocks := make(map[uint32]Block)
	// key: instrIdx of loop instruction
	loops := make(map[uint32]Loop)
	// key: instrIdx of branch instruction
	branchJumps := make(map[uint32][]uint32)

	var stack []uint32
	if codeSection, err := module.GetCodeSection(); err == nil {
		if code, found := codeSection.Codes[funcIdx]; found {
			var currentInstrIdx uint32 = 0
			var currentLabelIdx uint32 = 1
			disassemblyInstrs(code.Function.Expr.Instructions, &currentInstrIdx, currentLabelIdx, &labelsL, &labelsI, &blocks, &loops, &ifBlocks, stack, &branchJumps, &disassembledInstrs)
		}
	}
	return Disassembly{funcIdx, disassembledInstrs, labelsL, labelsI, blocks, loops, ifBlocks, branchJumps}
}

// Disassembles an instruction
func disassemblyInstrs(instrs []instructions2.Instruction, currentInstrIdx *uint32, currentLabelIdx uint32, labelsL *map[uint32]LabelL, labelsI *map[uint32]LabelI, blocks *map[uint32]Block, loops *map[uint32]Loop, ifBlocks *map[uint32]IfBlock, stack []uint32, branchJumps *map[uint32][]uint32, disassembledInstr *map[uint32]InstrDisassembly) {
	for _, instruction := range instrs {
		originalInstrIdx := *currentInstrIdx
		(*disassembledInstr)[originalInstrIdx] = InstrDisassembly{instruction, instruction.Position()}
		*currentInstrIdx++

		var blockEnd uint32
		var blockContinuation uint32

		var loopEnd uint32
		var loopContinuation uint32

		var thenEnd uint32
		var elseEnd uint32
		var ifContinuation uint32

		// beginning of control instruction
		if instruction.Name() == "block" ||
			instruction.Name() == "loop" ||
			instruction.Name() == "if" {
			if _, err := instruction.Instr(); err != nil {
				if _, err := instruction.ElseInstr(); err != nil {
					stack = append(stack, originalInstrIdx)
				}
			}
		}

		if instructions, err := instruction.Instr(); err == nil {
			var newStack []uint32
			for _, element := range stack {
				newStack = append(newStack, element)
			}
			newStack = append(newStack, originalInstrIdx)
			disassemblyInstrs(instructions, currentInstrIdx, currentLabelIdx+1, labelsL, labelsI, blocks, loops, ifBlocks, newStack, branchJumps, disassembledInstr)
		}

		if instruction.Name() == "block" {
			blockEnd = *currentInstrIdx - 1
			blockContinuation = *currentInstrIdx // first instruction after the block
		}
		if instruction.Name() == "loop" {
			loopEnd = *currentInstrIdx - 1
			loopContinuation = originalInstrIdx + 1 // first instruction of the loop
		}
		if instruction.Name() == "if" {
			thenEnd = *currentInstrIdx - 1
			ifContinuation = *currentInstrIdx // first instruction after the if
		}

		if instructions, err := instruction.ElseInstr(); err == nil {
			disassemblyInstrs(instructions, currentInstrIdx, currentLabelIdx+1, labelsL, labelsI, blocks, loops, ifBlocks, append(stack, originalInstrIdx), branchJumps, disassembledInstr)
		}
		if instruction.Name() == "if" {
			elseEnd = *currentInstrIdx - 1
			ifContinuation = *currentInstrIdx
		}
		(*disassembledInstr)[originalInstrIdx] = InstrDisassembly{instruction, instruction.Position()}

		if instruction.Name() == "block" {
			(*blocks)[originalInstrIdx] = Block{blockEnd, blockContinuation}
		}
		if instruction.Name() == "loop" {
			(*loops)[originalInstrIdx] = Loop{loopEnd, loopContinuation}
		}
		if instruction.Name() == "if" {
			(*ifBlocks)[originalInstrIdx] = IfBlock{thenEnd, elseEnd, ifContinuation}
		}

		// after control instruction
		if instruction.Name() == "block" ||
			instruction.Name() == "loop" ||
			instruction.Name() == "if" {
			//(*labelsL)[originalInstrIdx+1] = LabelL{originalInstrIdx, uint32(*currentInstrIdx)}
			(*labelsL)[currentLabelIdx] = LabelL{originalInstrIdx, uint32(*currentInstrIdx)}
			//(*labelsI)[originalInstrIdx] = LabelI{originalInstrIdx + 1, uint32(*currentInstrIdx)}
			(*labelsI)[originalInstrIdx] = LabelI{currentLabelIdx, uint32(*currentInstrIdx)}
			//*currentLabelIdx++
		}

		if instruction.Name() == "br" ||
			instruction.Name() == "br_if" ||
			instruction.Name() == "br_table" {
			(*branchJumps)[originalInstrIdx] = stack
		}
	}
}

// Returns the function body of a given function
func GetFuncBody(funcIdx uint32, module *modules.Module) ([]instructions2.Instruction, error) {
	if codeSection, err := module.GetCodeSection(); err != nil {
		return codeSection.Codes[funcIdx].Function.Expr.Instructions, nil
	} else {
		return nil, errors.New(fmt.Sprintf("function body for funcIdx '%v' does not exist", funcIdx))
	}
}

func GetInstructionDistribution(module *modules.Module) map[string]uint32 {
	var instructionDistribution = make(map[string]uint32)

	if codeSection, err := module.GetCodeSection(); err == nil {
		for _, code := range codeSection.Codes {
			getInstructionDistribution(code.Function.Expr.Instructions, &instructionDistribution)
		}
	}
	return instructionDistribution
}

func getInstructionDistribution(instructions []instructions2.Instruction, instructionDistribution *map[string]uint32) {
	for _, instruction := range instructions {
		if subInstructions, err := instruction.Instr(); err == nil {
			getInstructionDistribution(subInstructions, instructionDistribution)
		}
		if subInstructions, err := instruction.ElseInstr(); err == nil {
			getInstructionDistribution(subInstructions, instructionDistribution)
		}
		if counter, found := (*instructionDistribution)[instruction.Name()]; found {
			counter++
			(*instructionDistribution)[instruction.Name()] = counter
		} else {
			(*instructionDistribution)[instruction.Name()] = 1
		}
	}
}

func GetFuncParams(funcIdx uint32, module *modules.Module) (*types.FunctionType, error) {
	if typeSection, err := module.GetTypeSection(); err == nil {
		if functionSection, err := module.GetFunctionSection(); err == nil {
			if typeIdx, found := functionSection.TypeIdxs[funcIdx]; found {
				if functionType, found := typeSection.FunctionTypes[typeIdx]; found {
					return &functionType, nil
				} else {
					return nil, errors.New(fmt.Sprintf("no function type found for funcIdx: %v and typeIdx: %v", funcIdx, typeIdx))
				}
			} else {
				return nil, errors.New(fmt.Sprintf("no typeIdx found for funcIdx: %v", funcIdx))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("function section does not exist, funcIdx: %v", funcIdx))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("type section does not exist, funcIdx: %v", funcIdx))
	}
}

func GetFuncLocals(funcIdx uint32, module *modules.Module) []string {
	var locals []string
	if codeSection, err := module.GetCodeSection(); err == nil {
		if code, found := codeSection.Codes[funcIdx]; found {
			for _, local := range code.Function.Locals {
				var i uint32 = 0
				for ; i < local.N; i++ {
					locals = append(locals, local.ValType)
				}
			}
		} else {
			log.Fatalf("function with index %v not found", funcIdx)
		}
	} else {
		log.Fatal("code section does not exist")
	}
	return locals
}

func GetGlobals(module *modules.Module) (map[uint32]sections.Global, error) {
	globalSection, err := module.GetGlobalSection()
	if err != nil {
		return nil, err
	}

	return globalSection.Globals, nil
}

func GetGlobalsList(module *modules.Module) ([]sections.Global, uint32, error) {
	var startIndex uint32
	var globalsList []sections.Global
	globalImports := false
	firstIndex := true

	importSection, err := module.GetImportSection()
	if err == nil {
		var globalImportIdxs []uint32
		for globalIdx, _ := range importSection.GlobalImports {
			globalImportIdxs = append(globalImportIdxs, globalIdx)
			if firstIndex {
				startIndex = globalIdx
				firstIndex = false
				globalImports = true
			}
			if globalIdx < startIndex {
				startIndex = globalIdx
			}
		}
		sort.Slice(globalImportIdxs, func(i, j int) bool {
			return globalImportIdxs[i] < globalImportIdxs[j]
		})

		for _, globalIdx := range globalImportIdxs {
			globalsList = append(globalsList, sections.Global{importSection.GlobalImports[globalIdx].ImportDesc.GlobalType, new(instructions2.Expr)})
		}
	}

	globalSection, err := module.GetGlobalSection()
	if err != nil {
		if globalImports {
			return globalsList, startIndex, nil
		} else {
			return nil, 0, err
		}
	}
	var globalIdxs []uint32
	for globalIdx, _ := range globalSection.Globals {
		globalIdxs = append(globalIdxs, globalIdx)
		if firstIndex {
			startIndex = globalIdx
			firstIndex = false
		}
		if globalIdx < startIndex {
			startIndex = globalIdx
		}
	}
	sort.Slice(globalIdxs, func(i, j int) bool {
		return globalIdxs[i] < globalIdxs[j]
	})

	for _, globalIdx := range globalIdxs {
		globalsList = append(globalsList, globalSection.Globals[globalIdx])
	}

	return globalsList, startIndex, nil
}

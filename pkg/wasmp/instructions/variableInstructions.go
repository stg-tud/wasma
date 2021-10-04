package instructions

import (
	"errors"
	"fmt"
	"io"
	"log"
	"wasma/pkg/wasmp/structures"
	values2 "wasma/pkg/wasmp/values"
)

type VariableInstructionLocal struct {
	opcodeValue   byte
	nameValue     string
	localidxValue uint32
	position      uint32
	executor      func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (variableInstructionLocal *VariableInstructionLocal) Opcode() byte {
	return variableInstructionLocal.opcodeValue
}
func (variableInstructionLocal *VariableInstructionLocal) Name() string {
	return variableInstructionLocal.nameValue
}
func (variableInstructionLocal *VariableInstructionLocal) Position() uint32 {
	return variableInstructionLocal.position
}
func (variableInstructionLocal *VariableInstructionLocal) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Localidx() (uint32, error) {
	return variableInstructionLocal.localidxValue, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionLocal *VariableInstructionLocal) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return variableInstructionLocal.executor
}
func (variableInstructionLocal *VariableInstructionLocal) ToString() string {
	return fmt.Sprintf("%v %v", variableInstructionLocal.nameValue, variableInstructionLocal.localidxValue)
}

func newVariableInstructionLocal(opcode byte, name string, reader io.Reader) (Instruction, error) {
	var err error
	variableInstructionLocal := new(VariableInstructionLocal)
	variableInstructionLocal.position = values2.GetPosition()
	variableInstructionLocal.opcodeValue = opcode
	variableInstructionLocal.nameValue = name
	variableInstructionLocal.localidxValue, err = values2.ReadU32(reader)

	if err != nil {
		return nil, err
	}

	switch opcode {
	case 0x20: // local.get
		variableInstructionLocal.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
			if originalLocal, currentLocal, err := environment.GetLocal(variableInstructionLocal.localidxValue); err == nil {
				environment.Push(currentLocal)

				environment.AddInput(instrIdx, originalLocal)
				environment.AddOutput(instrIdx, currentLocal)
			} else {
				log.Fatalf("no local with localIdx: %v", variableInstructionLocal.localidxValue)
			}
		}
	case 0x21: // local.set
		variableInstructionLocal.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
			variable := environment.Pop()
			_, currentLocal, err := environment.SetLocal(variableInstructionLocal.localidxValue, variable)
			if err != nil {
				log.Fatalf("not able to set local for localIdx: %v", variableInstructionLocal.localidxValue)
			}

			currentLocal.LocalGlobalIn = true
			environment.AddInput(instrIdx, variable)
			environment.AddOutput(instrIdx, currentLocal)
		}
	case 0x22: // local.tee
		variableInstructionLocal.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
			variable := environment.Pop()

			newVariable := environment.NewVariable(variable.VariableDataType)
			if variable.VariableType == "C" || variable.VariableType == "V" {
				if variable.Value != "unknown" {
					newVariable.Value = fmt.Sprintf("%v", variable.Value)
				}
			} else {
				newVariable.Value = fmt.Sprintf("value(%v)", variable.VariableName)
			}
			environment.Push(newVariable) // push copy as new variable for further processing
			environment.Push(variable)    // push original one

			variable = environment.Pop()
			_, currentLocal, err := environment.SetLocal(variableInstructionLocal.localidxValue, variable)
			if err != nil {
				log.Fatalf("not able to set local for localIdx: %v", variableInstructionLocal.localidxValue)
			}

			currentLocal.LocalGlobalIn = true
			environment.AddInput(instrIdx, variable)
			environment.AddOutput(instrIdx, currentLocal)
			environment.AddOutput(instrIdx, newVariable)
		}
	default:
		return nil, errors.New(fmt.Sprintf("invalid opcode for local instruction, got: %v", opcode))
	}

	return variableInstructionLocal, nil
}

func newLocalGet(reader io.Reader) (Instruction, error) {
	return newVariableInstructionLocal(0x20, "local.get", reader)
}

func newLocalSet(reader io.Reader) (Instruction, error) {
	return newVariableInstructionLocal(0x21, "local.set", reader)
}

func newLocalTee(reader io.Reader) (Instruction, error) {
	return newVariableInstructionLocal(0x22, "local.tee", reader)
}

type VariableInstructionGlobal struct {
	opcodeValue    byte
	nameValue      string
	globalidxValue uint32
	position       uint32
	executor       func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (variableInstructionGlobal *VariableInstructionGlobal) Opcode() byte {
	return variableInstructionGlobal.opcodeValue
}
func (variableInstructionGlobal *VariableInstructionGlobal) Name() string {
	return variableInstructionGlobal.nameValue
}
func (variableInstructionGlobal *VariableInstructionGlobal) Position() uint32 {
	return variableInstructionGlobal.position
}
func (variableInstructionGlobal *VariableInstructionGlobal) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) Globalidx() (uint32, error) {
	return variableInstructionGlobal.globalidxValue, nil
}
func (variableInstructionGlobal *VariableInstructionGlobal) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (variableInstructionGlobal *VariableInstructionGlobal) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return variableInstructionGlobal.executor
}
func (variableInstructionGlobal *VariableInstructionGlobal) ToString() string {
	return fmt.Sprintf("%v %v", variableInstructionGlobal.nameValue, variableInstructionGlobal.globalidxValue)
}

func newVariableInstructionGlobal(opcode byte, name string, reader io.Reader) (Instruction, error) {
	var err error
	variableInstructionGlobal := new(VariableInstructionGlobal)
	variableInstructionGlobal.position = values2.GetPosition()
	variableInstructionGlobal.opcodeValue = opcode
	variableInstructionGlobal.nameValue = name
	variableInstructionGlobal.globalidxValue, err = values2.ReadU32(reader)

	if err != nil {
		return nil, err
	}

	switch opcode {
	case 0x23: // global.get
		variableInstructionGlobal.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
			if originalGlobal, currentGlobal, err := environment.GetGlobal(variableInstructionGlobal.globalidxValue); err == nil {
				environment.Push(currentGlobal)

				environment.AddInput(instrIdx, originalGlobal)
				environment.AddOutput(instrIdx, currentGlobal)
			} else {
				log.Fatalf("no global with globalIdx: %v", variableInstructionGlobal.globalidxValue)
			}
		}
	case 0x24: // global.set
		variableInstructionGlobal.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
			variable := environment.Pop()
			_, currentGlobal, err := environment.SetGlobal(variableInstructionGlobal.globalidxValue, variable)
			if err != nil {
				log.Fatalf("not able to set global for globalIdx: %v", variableInstructionGlobal.globalidxValue)
			}

			currentGlobal.LocalGlobalIn = true
			environment.AddInput(instrIdx, variable)
			environment.AddOutput(instrIdx, currentGlobal)
		}
	}
	return variableInstructionGlobal, nil
}

func newGlobalGet(reader io.Reader) (Instruction, error) {
	return newVariableInstructionGlobal(0x23, "global.get", reader)
}

func newGlobalSet(reader io.Reader) (Instruction, error) {
	return newVariableInstructionGlobal(0x24, "global.set", reader)
}

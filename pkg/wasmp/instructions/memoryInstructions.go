package instructions

import (
	"errors"
	"fmt"
	"io"
	"wasma/pkg/wasmp/structures"
	values2 "wasma/pkg/wasmp/values"
)

type MemoryInstruction struct {
	opcodeValue byte
	nameValue   string
	alignValue  uint32
	offsetValue uint32
	position    uint32
	executor    func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (memoryInstruction *MemoryInstruction) Opcode() byte {
	return memoryInstruction.opcodeValue
}
func (memoryInstruction *MemoryInstruction) Name() string {
	return memoryInstruction.nameValue
}
func (memoryInstruction *MemoryInstruction) Position() uint32 {
	return memoryInstruction.position
}
func (memoryInstruction *MemoryInstruction) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Align() (uint32, error) {
	return memoryInstruction.alignValue, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Offset() (uint32, error) {
	return memoryInstruction.offsetValue, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (memoryInstruction *MemoryInstruction) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return memoryInstruction.executor
}
func (memoryInstruction *MemoryInstruction) ToString() string {
	return fmt.Sprintf("%v %v, %v", memoryInstruction.nameValue, memoryInstruction.alignValue, memoryInstruction.offsetValue)
}

var memoryLoadI32 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newConstant := environment.NewConstant("unknown", "i32")
	environment.Push(newConstant)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newConstant)
}

var memoryLoadI64 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newConstant := environment.NewConstant("unknown", "i64")
	environment.Push(newConstant)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newConstant)
}

var memoryLoadF32 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newConstant := environment.NewConstant("unknown", "f32")
	environment.Push(newConstant)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newConstant)
}

var memoryLoadF64 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newConstant := environment.NewConstant("unknown", "f64")
	environment.Push(newConstant)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newConstant)
}

var memoryStore = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable1 := environment.Pop()
	variable2 := environment.Pop()

	environment.AddInput(instrIdx, variable1)
	environment.AddInput(instrIdx, variable2)
}

func newMemoryInstruction(
	opcode byte,
	name string,
	executor func(instrIdx uint32, instruction Instruction, environment structures.Environment),
	reader io.Reader) (Instruction, error) {
	memoryInstruction := new(MemoryInstruction)
	memoryInstruction.position = values2.GetPosition()
	memoryInstruction.opcodeValue = opcode
	memoryInstruction.nameValue = name
	memoryInstruction.executor = executor

	align, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	offset, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	memoryInstruction.alignValue = align
	memoryInstruction.offsetValue = offset

	return memoryInstruction, nil
}

func newI32Load(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x28, "i32.load", memoryLoadI32, reader)
}

func newI64Load(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x29, "i64.load", memoryLoadI64, reader)
}

func newF32Load(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x2A, "f32.load", memoryLoadF32, reader)
}

func newF64Load(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x2B, "f64.load", memoryLoadF64, reader)
}

func newI32Load8s(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x2C, "i32.load8_s", memoryLoadI32, reader)
}

func newI32Load8u(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x2D, "i32.load8_u", memoryLoadI32, reader)
}

func newI32Load16s(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x2E, "i32.load16_s", memoryLoadI32, reader)
}

func newI32Load16u(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x2F, "i32.load16_u", memoryLoadI32, reader)
}

func newI64Load8s(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x30, "i64.load8_s", memoryLoadI64, reader)
}

func newI64Load8u(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x31, "i64.load8_u", memoryLoadI64, reader)
}

func newI64Load16s(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x32, "i64.load16_s", memoryLoadI64, reader)
}

func newI64Load16u(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x33, "i64.load16_u", memoryLoadI64, reader)
}

func newI64Load32s(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x34, "i64.load32_s", memoryLoadI64, reader)
}

func newI64Load32u(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x35, "i64.load32_u", memoryLoadI64, reader)
}

func newI32Store(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x36, "i32.store", memoryStore, reader)
}

func newI64Store(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x37, "i64.store", memoryStore, reader)
}

func newF32Store(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x38, "f32.store", memoryStore, reader)
}

func newF64Store(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x39, "f64.store", memoryStore, reader)
}

func newI32Store8(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x3A, "i32.store8", memoryStore, reader)
}

func newI32Store16(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x3B, "i32.store16", memoryStore, reader)
}

func newI64Store8(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x3C, "i64.store8", memoryStore, reader)
}

func newI64Store16(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x3D, "i64.store16", memoryStore, reader)
}

func newI64Store32(reader io.Reader) (Instruction, error) {
	return newMemoryInstruction(0x3E, "i64.store32", memoryStore, reader)
}

var memorySize = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	newConstant := environment.NewConstant("memory size", "i32")
	environment.Push(newConstant)

	environment.AddOutput(instrIdx, newConstant)
}

var memoryGrow = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newConstant := environment.NewConstant("unknown", "i32")
	environment.Push(newConstant)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newConstant)
}

func newAdditionalMemoryInstruction(
	opcode byte,
	name string,
	executor func(instrIdx uint32, instruction Instruction, environment structures.Environment),
	reader io.Reader) (Instruction, error) {
	additionalMemoryInstruction := new(SimpleInstruction)
	additionalMemoryInstruction.position = values2.GetPosition()
	additionalMemoryInstruction.opcodeValue = opcode
	additionalMemoryInstruction.nameValue = name
	additionalMemoryInstruction.executor = executor

	nextByte, err := values2.ReadNextByte(reader)
	if err != nil {
		return nil, err
	}

	if nextByte != 0x00 {
		return nil, errors.New(fmt.Sprintf("Error while reading %s. Expected 0x00 but got: 0x%x", name, nextByte))
	}

	return additionalMemoryInstruction, nil
}

func newMemorySize(reader io.Reader) (Instruction, error) {
	return newAdditionalMemoryInstruction(0x3F, "memory.size", memorySize, reader)
}

func newMemoryGrow(reader io.Reader) (Instruction, error) {
	return newAdditionalMemoryInstruction(0x40, "memory.grow", memoryGrow, reader)
}

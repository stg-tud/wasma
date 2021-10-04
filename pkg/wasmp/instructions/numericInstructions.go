package instructions

import (
	"errors"
	"fmt"
	"io"
	"wasma/pkg/wasmp/structures"
	values2 "wasma/pkg/wasmp/values"
)

type NumericInstructionI32 struct {
	OpcodeValue   byte
	NameValue     string
	I32Value      int32
	PositionValue uint32
	executor      func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (numericInstructionI32 *NumericInstructionI32) Opcode() byte {
	return numericInstructionI32.OpcodeValue
}
func (numericInstructionI32 *NumericInstructionI32) Name() string {
	return numericInstructionI32.NameValue
}
func (numericInstructionI32 *NumericInstructionI32) Position() uint32 {
	return numericInstructionI32.PositionValue
}
func (numericInstructionI32 *NumericInstructionI32) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) I32() (int32, error) {
	return numericInstructionI32.I32Value, nil
}
func (numericInstructionI32 *NumericInstructionI32) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI32 *NumericInstructionI32) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return numericInstructionI32.executor
}
func (numericInstructionI32 *NumericInstructionI32) ToString() string {
	return fmt.Sprintf("%v %v (0x%x)", numericInstructionI32.NameValue, uint32(numericInstructionI32.I32Value), uint32(numericInstructionI32.I32Value))
}

func NewI32(reader io.Reader) (Instruction, error) {
	var err error
	i32 := new(NumericInstructionI32)
	i32.PositionValue = values2.GetPosition()
	i32.OpcodeValue = 0x41
	i32.NameValue = "i32.const"
	i32.I32Value, err = values2.ReadS32(reader)

	i32.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
		newConstant := environment.NewConstant(fmt.Sprintf("%v", i32.I32Value), "i32")
		environment.Push(newConstant)

		environment.AddOutput(instrIdx, newConstant)
	}

	return i32, err
}

type NumericInstructionI64 struct {
	opcodeValue byte
	nameValue   string
	i64Value    int64
	position    uint32
	executor    func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (numericInstructionI64 *NumericInstructionI64) Opcode() byte {
	return numericInstructionI64.opcodeValue
}
func (numericInstructionI64 *NumericInstructionI64) Name() string {
	return numericInstructionI64.nameValue
}
func (numericInstructionI64 *NumericInstructionI64) Position() uint32 {
	return numericInstructionI64.position
}
func (numericInstructionI64 *NumericInstructionI64) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) I64() (int64, error) {
	return numericInstructionI64.i64Value, nil
}
func (numericInstructionI64 *NumericInstructionI64) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionI64 *NumericInstructionI64) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return numericInstructionI64.executor
}
func (numericInstructionI64 *NumericInstructionI64) ToString() string {
	return fmt.Sprintf("%v %v (0x%x)", numericInstructionI64.nameValue, uint64(numericInstructionI64.i64Value), uint64(numericInstructionI64.i64Value))
}

func newI64(reader io.Reader) (Instruction, error) {
	var err error
	i64 := new(NumericInstructionI64)
	i64.position = values2.GetPosition()
	i64.opcodeValue = 0x42
	i64.nameValue = "i64.const"
	i64.i64Value, err = values2.ReadS64(reader)

	i64.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
		newConstant := environment.NewConstant(fmt.Sprintf("%v", i64.i64Value), "i64")
		environment.Push(newConstant)

		environment.AddOutput(instrIdx, newConstant)
	}

	return i64, err
}

type NumericInstructionF32 struct {
	opcodeValue byte
	nameValue   string
	f32Value    float32
	position    uint32
	executor    func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (numericInstructionF32 *NumericInstructionF32) Opcode() byte {
	return numericInstructionF32.opcodeValue
}
func (numericInstructionF32 *NumericInstructionF32) Name() string {
	return numericInstructionF32.nameValue
}
func (numericInstructionF32 *NumericInstructionF32) Position() uint32 {
	return numericInstructionF32.position
}
func (numericInstructionF32 *NumericInstructionF32) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) F32() (float32, error) {
	return numericInstructionF32.f32Value, nil
}
func (numericInstructionF32 *NumericInstructionF32) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF32 *NumericInstructionF32) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return numericInstructionF32.executor
}
func (numericInstructionF32 *NumericInstructionF32) ToString() string {
	return fmt.Sprintf("%v %v (%x)", numericInstructionF32.nameValue, numericInstructionF32.f32Value, numericInstructionF32.f32Value)
}

func newF32(reader io.Reader) (Instruction, error) {
	var err error
	f32 := new(NumericInstructionF32)
	f32.position = values2.GetPosition()
	f32.opcodeValue = 0x43
	f32.nameValue = "f32.const"
	f32.f32Value, err = values2.ReadF32(reader)

	f32.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
		newConstant := environment.NewConstant(fmt.Sprintf("%v", f32.f32Value), "f32")
		environment.Push(newConstant)

		environment.AddOutput(instrIdx, newConstant)
	}

	return f32, err
}

type NumericInstructionF64 struct {
	opcodeValue byte
	nameValue   string
	f64Value    float64
	position    uint32
	executor    func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (numericInstructionF64 *NumericInstructionF64) Opcode() byte {
	return numericInstructionF64.opcodeValue
}
func (numericInstructionF64 *NumericInstructionF64) Name() string {
	return numericInstructionF64.nameValue
}
func (numericInstructionF64 *NumericInstructionF64) Position() uint32 {
	return numericInstructionF64.position
}
func (numericInstructionF64 *NumericInstructionF64) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) F64() (float64, error) {
	return numericInstructionF64.f64Value, nil
}
func (numericInstructionF64 *NumericInstructionF64) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionF64 *NumericInstructionF64) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return numericInstructionF64.executor
}
func (numericInstructionF64 *NumericInstructionF64) ToString() string {
	return fmt.Sprintf("%v %v (%x)", numericInstructionF64.nameValue, numericInstructionF64.f64Value, numericInstructionF64.f64Value)
}

func newF64(reader io.Reader) (Instruction, error) {
	var err error
	f64 := new(NumericInstructionF64)
	f64.position = values2.GetPosition()
	f64.opcodeValue = 0x44
	f64.nameValue = "f64.const"
	f64.f64Value, err = values2.ReadF64(reader)

	f64.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
		newConstant := environment.NewConstant(fmt.Sprintf("%v", f64.f64Value), "f64")
		environment.Push(newConstant)

		environment.AddOutput(instrIdx, newConstant)
	}

	return f64, err
}

var iunopExecutorI32 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("i32")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var iunopExecutorI64 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("i64")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var ibinopExecutorI32 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable1 := environment.Pop()
	variable2 := environment.Pop()
	newVariable := environment.NewVariable("i32")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable1)
	environment.AddInput(instrIdx, variable2)
	environment.AddOutput(instrIdx, newVariable)
}
var ibinopExecutorI64 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable1 := environment.Pop()
	variable2 := environment.Pop()
	newVariable := environment.NewVariable("i64")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable1)
	environment.AddInput(instrIdx, variable2)
	environment.AddOutput(instrIdx, newVariable)
}
var funopExecutorF32 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("f32")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var funopExecutorF64 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("f64")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var fbinopExecutorF32 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable1 := environment.Pop()
	variable2 := environment.Pop()
	newVariable := environment.NewVariable("f32")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable1)
	environment.AddInput(instrIdx, variable2)
	environment.AddOutput(instrIdx, newVariable)
}
var fbinopExecutorF64 = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable1 := environment.Pop()
	variable2 := environment.Pop()
	newVariable := environment.NewVariable("f64")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable1)
	environment.AddInput(instrIdx, variable2)
	environment.AddOutput(instrIdx, newVariable)
}
var itestopExecutorBool = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("i32") // boolean integer result
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var irelopExecutorBool = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable1 := environment.Pop()
	variable2 := environment.Pop()
	newVariable := environment.NewVariable("i32") // boolean integer result
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable1)
	environment.AddInput(instrIdx, variable2)
	environment.AddOutput(instrIdx, newVariable)
}
var frelopExecutorBool = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable1 := environment.Pop()
	variable2 := environment.Pop()
	newVariable := environment.NewVariable("i32") // boolean integer result
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable1)
	environment.AddInput(instrIdx, variable2)
	environment.AddOutput(instrIdx, newVariable)
}
var convertToI32Executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("i32")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var convertToI64Executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("i64")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var convertToF32Executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("f32")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var convertToF64Executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	variable := environment.Pop()
	newVariable := environment.NewVariable("f64")
	environment.Push(newVariable)

	environment.AddInput(instrIdx, variable)
	environment.AddOutput(instrIdx, newVariable)
}
var extendExecutor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	// WebAssembly 1.1
}
var truncExecutor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	// WebAssembly 1.1
}

func newI32Eqz(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x45, "i32.eqz", values2.GetPosition(), itestopExecutorBool}, nil
}

func newI32Eq(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x46, "i32.eq", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Ne(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x47, "i32.ne", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Lts(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x48, "i32.lt_s", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Ltu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x49, "i32.lt_u", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Gts(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x4A, "i32.gt_s", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Gtu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x4B, "i32.gt_u", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Les(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x4C, "i32.le_s", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Leu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x4D, "i32.le_u", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Ges(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x4E, "i32.ge_s", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI32Geu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x4F, "i32.ge_u", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Eqz(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x50, "i64.eqz", values2.GetPosition(), itestopExecutorBool}, nil
}

func newI64Eq(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x51, "i64.eq", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Ne(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x52, "i64.ne", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Lts(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x53, "i64.lt_s", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Ltu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x54, "i64.lt_u", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Gts(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x55, "i64.gt_s", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Gtu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x56, "i64.gt_u", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Les(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x57, "i64.le_s", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Leu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x58, "i64.le_u", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Ges(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x59, "i64.ge_s", values2.GetPosition(), irelopExecutorBool}, nil
}

func newI64Geu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x5A, "i64.ge_u", values2.GetPosition(), irelopExecutorBool}, nil
}

func newF32Eq(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x5B, "f32.eq", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF32Ne(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x5C, "f32.ne", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF32Lt(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x5D, "f32.lt", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF32Gt(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x5E, "f32.gt", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF32Le(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x5F, "f32.le", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF32Ge(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x60, "f32.ge", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF64Eq(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x61, "f64.eq", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF64Ne(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x62, "f64.ne", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF64Lt(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x63, "f64.lt", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF64Gt(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x64, "f64.gt", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF64Le(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x65, "f64.le", values2.GetPosition(), frelopExecutorBool}, nil
}

func newF64Ge(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x66, "f64.ge", values2.GetPosition(), frelopExecutorBool}, nil
}

func newI32Clz(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x67, "i32.clz", values2.GetPosition(), iunopExecutorI32}, nil
}

func newI32Ctz(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x68, "i32.ctz", values2.GetPosition(), iunopExecutorI32}, nil
}

func newI32Popcnt(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x69, "i32.popcnt", values2.GetPosition(), iunopExecutorI32}, nil
}

func newI32Add(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x6A, "i32.add", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Sub(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x6B, "i32.sub", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Mul(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x6C, "i32.mul", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Divs(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x6D, "i32.div_s", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Divu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x6E, "i32.div_u", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Rems(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x6F, "i32.rem_s", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Remu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x70, "i32.rem_u", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32And(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x71, "i32.and", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Or(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x72, "i32.or", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Xor(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x73, "i32.xor", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Shl(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x74, "i32.shl", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Shrs(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x75, "i32.shr_s", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Shru(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x76, "i32.shr_u", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Rotl(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x77, "i32.rotl", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI32Rotr(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x78, "i32.rotr", values2.GetPosition(), ibinopExecutorI32}, nil
}

func newI64Clz(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x79, "i64.clz", values2.GetPosition(), iunopExecutorI64}, nil
}

func newI64Ctz(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x7A, "i64.ctz", values2.GetPosition(), iunopExecutorI64}, nil
}

func newI64Popcnt(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x7B, "i64.popcnt", values2.GetPosition(), iunopExecutorI64}, nil
}

func newI64Add(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x7C, "i64.add", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Sub(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x7D, "i64.sub", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Mul(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x7E, "i64.mul", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Divs(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x7F, "i64.div_s", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Divu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x80, "i64.div_u", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Rems(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x81, "i64.rem_s", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Remu(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x82, "i64.rem_u", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64And(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x83, "i64.and", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Or(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x84, "i64.or", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Xor(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x85, "i64.xor", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Shl(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x86, "i64.shl", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Shrs(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x87, "i64.shr_s", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Shru(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x88, "i64.shr_u", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Rotl(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x89, "i64.rotl", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newI64Rotr(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x8A, "i64.rotr", values2.GetPosition(), ibinopExecutorI64}, nil
}

func newF32Abs(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x8B, "f32.abs", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Neg(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x8C, "f32.neg", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Ceil(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x8D, "f32.ceil", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Floor(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x8E, "f32.floor", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Trunc(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x8F, "f32.trunc", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Nearest(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x90, "f32.nearest", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Sqrt(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x91, "f32.sqrt", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Add(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x92, "f32.add", values2.GetPosition(), fbinopExecutorF32}, nil
}

func newF32Sub(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x93, "f32.sub", values2.GetPosition(), fbinopExecutorF32}, nil
}

func newF32Mul(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x94, "f32.mul", values2.GetPosition(), fbinopExecutorF32}, nil
}

func newF32Div(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x95, "f32.div", values2.GetPosition(), fbinopExecutorF32}, nil
}

func newF32Min(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x96, "f32.min", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Max(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x97, "f32.max", values2.GetPosition(), funopExecutorF32}, nil
}

func newF32Copysign(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x98, "f32.copysign", values2.GetPosition(), fbinopExecutorF32}, nil
}

func newF64Abs(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x99, "f64.abs", values2.GetPosition(), funopExecutorF64}, nil
}

func newF64Neg(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x9A, "f64.neg", values2.GetPosition(), funopExecutorF64}, nil
}

func newF64Ceil(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x9B, "f64.ceil", values2.GetPosition(), funopExecutorF64}, nil
}

func newF64Floor(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x9C, "f64.floor", values2.GetPosition(), funopExecutorF64}, nil
}

func newF64Trunc(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x9D, "f64.trunc", values2.GetPosition(), funopExecutorF64}, nil
}

func newF64Nearest(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x9E, "f64.nearest", values2.GetPosition(), funopExecutorF64}, nil
}

func newF64Sqrt(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0x9F, "f64.sqrt", values2.GetPosition(), funopExecutorF64}, nil
}

func newF64Add(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA0, "f64.add", values2.GetPosition(), fbinopExecutorF64}, nil
}

func newF64Sub(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA1, "f64.sub", values2.GetPosition(), fbinopExecutorF64}, nil
}

func newF64Mul(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA2, "f64.mul", values2.GetPosition(), fbinopExecutorF64}, nil
}

func newF64Div(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA3, "f64.div", values2.GetPosition(), fbinopExecutorF64}, nil
}

func newF64Min(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA4, "f64.min", values2.GetPosition(), fbinopExecutorF64}, nil
}

func newF64Max(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA5, "f64.max", values2.GetPosition(), fbinopExecutorF64}, nil
}

func newF64Copysign(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA6, "f64.copysign", values2.GetPosition(), fbinopExecutorF64}, nil
}

func newI32WrapI64(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA7, "i32.wrap_i64", values2.GetPosition(), convertToI32Executor}, nil
}

func newI32TruncF32s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA8, "i32.trunc_f32_s", values2.GetPosition(), convertToI32Executor}, nil
}

func newI32TruncF32u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xA9, "i32.trunc_f32_u", values2.GetPosition(), convertToI32Executor}, nil
}

func newI32TruncF64s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xAA, "i32.trunc_f64_s", values2.GetPosition(), convertToI32Executor}, nil
}

func newI32TruncF64u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xAB, "i32.trunc_f64_u", values2.GetPosition(), convertToI32Executor}, nil
}

func newI64ExtendI32s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xAC, "i64.extend_i32_s", values2.GetPosition(), convertToI64Executor}, nil
}

func newI64ExtendI32u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xAD, "i64.extend_i32_u", values2.GetPosition(), convertToI64Executor}, nil
}

func newI64TruncF32s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xAE, "i64.trunc_f32_s", values2.GetPosition(), convertToI64Executor}, nil
}

func newI64TruncF32u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xAF, "i64.trunc_f32_u", values2.GetPosition(), convertToI64Executor}, nil
}

func newI64TruncF64s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB0, "i64.trunc_f64_s", values2.GetPosition(), convertToI64Executor}, nil
}

func newI64TruncF64u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB1, "i64.trunc_f64_u", values2.GetPosition(), convertToI64Executor}, nil
}

func newF32ConvertI32s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB2, "f32.convert_i32_s", values2.GetPosition(), convertToF32Executor}, nil
}

func newF32ConvertI32u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB3, "f32.convert_i32_u", values2.GetPosition(), convertToF32Executor}, nil
}

func newF32ConvertI64s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB4, "f32.convert_i64_s", values2.GetPosition(), convertToF32Executor}, nil
}

func newF32ConvertI64u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB5, "f32.convert_i64_u", values2.GetPosition(), convertToF32Executor}, nil
}

func newF32DemoteF64(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB6, "f32.demote_f64", values2.GetPosition(), convertToF32Executor}, nil
}

func newF64ConvertI32s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB7, "f64.convert_i32_s", values2.GetPosition(), convertToF64Executor}, nil
}

func newF64ConvertI32u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB8, "f64.convert_i32_u", values2.GetPosition(), convertToF64Executor}, nil
}

func newF64ConvertI64s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xB9, "f64.convert_i64_s", values2.GetPosition(), convertToF64Executor}, nil
}

func newF64ConvertI64u(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xBA, "f64.convert_i64_u", values2.GetPosition(), convertToF64Executor}, nil
}

func newF64PromoteF32(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xBB, "f64.promote_f32", values2.GetPosition(), convertToF64Executor}, nil
}

func newI32ReinterpretF32(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xBC, "i32.reinterpret_f32", values2.GetPosition(), convertToI32Executor}, nil
}

func newI64ReinterpretF64(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xBD, "i64.reinterpret_f64", values2.GetPosition(), convertToI64Executor}, nil
}

func newF32ReinterpretI32(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xBE, "f32.reinterpret_i32", values2.GetPosition(), convertToF32Executor}, nil
}

func newF64ReinterpretI64(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xBF, "f64.reinterpret_i64", values2.GetPosition(), convertToF64Executor}, nil
}

func newI32Extend8s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xC0, "i32.extend8_s", values2.GetPosition(), extendExecutor}, nil
}

func newI32Extend16s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xC1, "i32.extend16_s", values2.GetPosition(), extendExecutor}, nil
}

func newI64Extend8s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xC2, "i64.extend8_s", values2.GetPosition(), extendExecutor}, nil
}

func newI64Extend16s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xC3, "i64.extend16_s", values2.GetPosition(), extendExecutor}, nil
}

func newI64Extend32s(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{0xC4, "i64.extend32_s", values2.GetPosition(), extendExecutor}, nil
}

type NumericInstructionTrunc struct {
	opcodeValue      byte
	nameValue        string
	truncOpcodeValue uint32
	position         uint32
	executor         func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (numericInstructionTrunc *NumericInstructionTrunc) Opcode() byte {
	return numericInstructionTrunc.opcodeValue
}
func (numericInstructionTrunc *NumericInstructionTrunc) Name() string {
	return numericInstructionTrunc.nameValue
}
func (numericInstructionTrunc *NumericInstructionTrunc) Position() uint32 {
	return numericInstructionTrunc.position
}
func (numericInstructionTrunc *NumericInstructionTrunc) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (numericInstructionTrunc *NumericInstructionTrunc) TruncOpcode() (uint32, error) {
	return numericInstructionTrunc.truncOpcodeValue, nil
}
func (numericInstructionTrunc *NumericInstructionTrunc) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return numericInstructionTrunc.executor
}
func (numericInstructionTrunc *NumericInstructionTrunc) ToString() string {
	return fmt.Sprintf("%v %v", numericInstructionTrunc.nameValue, numericInstructionTrunc.truncOpcodeValue)
}

func newTrunc(reader io.Reader) (Instruction, error) {
	opcode, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	numericInstructionTrunc := new(NumericInstructionTrunc)
	numericInstructionTrunc.position = values2.GetPosition() // PositionValue of trunc Opcode
	numericInstructionTrunc.opcodeValue = 0xFC
	numericInstructionTrunc.truncOpcodeValue = opcode
	numericInstructionTrunc.executor = truncExecutor
	switch opcode {
	case 0x00:
		numericInstructionTrunc.nameValue = "i32.trunc_sat_f32_s"
	case 0x01:
		numericInstructionTrunc.nameValue = "i32.trunc_sat_f32_u"
	case 0x02:
		numericInstructionTrunc.nameValue = "i32.trunc_sat_f64_s"
	case 0x03:
		numericInstructionTrunc.nameValue = "i32.trunc_sat_f64_u"
	case 0x04:
		numericInstructionTrunc.nameValue = "i64.trunc_sat_f32_s"
	case 0x05:
		numericInstructionTrunc.nameValue = "i64.trunc_sat_f32_u"
	case 0x06:
		numericInstructionTrunc.nameValue = "i64.trunc_sat_f64_s"
	case 0x07:
		numericInstructionTrunc.nameValue = "i64.trunc_sat_f64_u"
	default:
		return nil, errors.New(fmt.Sprintf("Error while parsing trunc sat instraction. Expected Opcode between 0x00 and 0x07 but got: %x", opcode))
	}

	return numericInstructionTrunc, nil
}

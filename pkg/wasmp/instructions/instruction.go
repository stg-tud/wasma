package instructions

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"wasma/pkg/wasmp/structures"
	"wasma/pkg/wasmp/values"
)

var attributeNotAvailable = errors.New("attribute not available")

type Instruction interface {
	Opcode() byte
	Name() string
	Position() uint32
	Blocktype() (string, error)
	Instr() ([]Instruction, error)
	ElseInstr() ([]Instruction, error)
	Labelidx() (uint32, error)
	Funcidx() (uint32, error)
	Typeidx() (uint32, error)
	Localidx() (uint32, error)
	Globalidx() (uint32, error)
	VecLabelidx() ([]uint32, error)
	Align() (uint32, error)
	Offset() (uint32, error)
	I32() (int32, error)
	I64() (int64, error)
	F32() (float32, error)
	F64() (float64, error)
	TruncOpcode() (uint32, error)
	Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment)
	ToString() string
}

type SimpleInstruction struct {
	opcodeValue byte
	nameValue   string
	position    uint32
	executor    func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (simpleInstruction *SimpleInstruction) Opcode() byte {
	return simpleInstruction.opcodeValue
}
func (simpleInstruction *SimpleInstruction) Name() string {
	return simpleInstruction.nameValue
}
func (simpleInstruction *SimpleInstruction) Position() uint32 {
	return simpleInstruction.position
}
func (simpleInstruction *SimpleInstruction) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (simpleInstruction *SimpleInstruction) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return simpleInstruction.executor
}

func (simpleInstruction *SimpleInstruction) ToString() string {
	if simpleInstruction.nameValue == "memory.grow" || simpleInstruction.nameValue == "memory.size" {
		return fmt.Sprintf("%v 0 (0x0)", simpleInstruction.nameValue)
	}
	return simpleInstruction.nameValue
}

var opcodes map[byte]func(io.Reader) (Instruction, error)
var initOpcodes = false

func getValtype(byteValue byte) (string, error) {
	switch byteValue {
	case 0x7F:
		return "i32", nil
	case 0x7E:
		return "i64", nil
	case 0x7D:
		return "f32", nil
	case 0x7C:
		return "f64", nil
	default:
		return "no valid value type", errors.New("no valid value type")
	}
}

func mapOpcode(opcode byte, reader io.Reader) (Instruction, error) {
	if !initOpcodes {
		initOpcodes = true
		opcodes = map[byte]func(io.Reader) (Instruction, error){
			0x00: newUnreachable,
			0x01: newNop,
			0x02: newBlock,
			0x03: newLoop,
			0x04: newIfElse,
			0x0C: newBr,
			0x0D: newBrIf,
			0x0E: newBrTable,
			0x0F: newReturn,
			0x10: newCall,
			0x11: newCallIndirect,
			0x1A: newDrop,
			0x1B: newSelect,
			0x20: newLocalGet,
			0x21: newLocalSet,
			0x22: newLocalTee,
			0x23: newGlobalGet,
			0x24: newGlobalSet,
			0x28: newI32Load,
			0x29: newI64Load,
			0x2A: newF32Load,
			0x2B: newF64Load,
			0x2C: newI32Load8s,
			0x2D: newI32Load8u,
			0x2E: newI32Load16s,
			0x2F: newI32Load16u,
			0x30: newI64Load8s,
			0x31: newI64Load8u,
			0x32: newI64Load16s,
			0x33: newI64Load16u,
			0x34: newI64Load32s,
			0x35: newI64Load32u,
			0x36: newI32Store,
			0x37: newI64Store,
			0x38: newF32Store,
			0x39: newF64Store,
			0x3A: newI32Store8,
			0x3B: newI32Store16,
			0x3C: newI64Store8,
			0x3D: newI64Store16,
			0x3E: newI64Store32,
			0x3F: newMemorySize,
			0x40: newMemoryGrow,
			0x41: NewI32,
			0X42: newI64,
			0x43: newF32,
			0x44: newF64,
			0x45: newI32Eqz,
			0x46: newI32Eq,
			0x47: newI32Ne,
			0x48: newI32Lts,
			0x49: newI32Ltu,
			0x4A: newI32Gts,
			0x4B: newI32Gtu,
			0x4C: newI32Les,
			0x4D: newI32Leu,
			0x4E: newI32Ges,
			0x4F: newI32Geu,
			0x50: newI64Eqz,
			0x51: newI64Eq,
			0x52: newI64Ne,
			0x53: newI64Lts,
			0x54: newI64Ltu,
			0x55: newI64Gts,
			0x56: newI64Gtu,
			0x57: newI64Les,
			0x58: newI64Leu,
			0x59: newI64Ges,
			0x5A: newI64Geu,
			0x5B: newF32Eq,
			0x5C: newF32Ne,
			0x5D: newF32Lt,
			0x5E: newF32Gt,
			0x5F: newF32Le,
			0x60: newF32Ge,
			0x61: newF64Eq,
			0x62: newF64Ne,
			0x63: newF64Lt,
			0x64: newF64Gt,
			0x65: newF64Le,
			0x66: newF64Ge,
			0x67: newI32Clz,
			0x68: newI32Ctz,
			0x69: newI32Popcnt,
			0x6A: newI32Add,
			0x6B: newI32Sub,
			0x6C: newI32Mul,
			0x6D: newI32Divs,
			0x6E: newI32Divu,
			0x6F: newI32Rems,
			0x70: newI32Remu,
			0x71: newI32And,
			0x72: newI32Or,
			0x73: newI32Xor,
			0x74: newI32Shl,
			0x75: newI32Shrs,
			0x76: newI32Shru,
			0x77: newI32Rotl,
			0x78: newI32Rotr,
			0x79: newI64Clz,
			0x7A: newI64Ctz,
			0x7B: newI64Popcnt,
			0x7C: newI64Add,
			0x7D: newI64Sub,
			0x7E: newI64Mul,
			0x7F: newI64Divs,
			0x80: newI64Divu,
			0x81: newI64Rems,
			0x82: newI64Remu,
			0x83: newI64And,
			0x84: newI64Or,
			0x85: newI64Xor,
			0x86: newI64Shl,
			0x87: newI64Shrs,
			0x88: newI64Shru,
			0x89: newI64Rotl,
			0x8A: newI64Rotr,
			0x8B: newF32Abs,
			0x8C: newF32Neg,
			0x8D: newF32Ceil,
			0x8E: newF32Floor,
			0x8F: newF32Trunc,
			0x90: newF32Nearest,
			0x91: newF32Sqrt,
			0x92: newF32Add,
			0x93: newF32Sub,
			0x94: newF32Mul,
			0x95: newF32Div,
			0x96: newF32Min,
			0x97: newF32Max,
			0x98: newF32Copysign,
			0x99: newF64Abs,
			0x9A: newF64Neg,
			0x9B: newF64Ceil,
			0x9C: newF64Floor,
			0x9D: newF64Trunc,
			0x9E: newF64Nearest,
			0x9F: newF64Sqrt,
			0xA0: newF64Add,
			0xA1: newF64Sub,
			0xA2: newF64Mul,
			0xA3: newF64Div,
			0xA4: newF64Min,
			0xA5: newF64Max,
			0xA6: newF64Copysign,
			0xA7: newI32WrapI64,
			0xA8: newI32TruncF32s,
			0xA9: newI32TruncF32u,
			0xAA: newI32TruncF64s,
			0xAB: newI32TruncF64u,
			0xAC: newI64ExtendI32s,
			0xAD: newI64ExtendI32u,
			0xAE: newI64TruncF32s,
			0xAF: newI64TruncF32u,
			0xB0: newI64TruncF64s,
			0xB1: newI64TruncF64u,
			0xB2: newF32ConvertI32s,
			0xB3: newF32ConvertI32u,
			0xB4: newF32ConvertI64s,
			0xB5: newF32ConvertI64u,
			0xB6: newF32DemoteF64,
			0xB7: newF64ConvertI32s,
			0xB8: newF64ConvertI32u,
			0xB9: newF64ConvertI64s,
			0xBA: newF64ConvertI64u,
			0xBB: newF64PromoteF32,
			0xBC: newI32ReinterpretF32,
			0xBD: newI64ReinterpretF64,
			0xBE: newF32ReinterpretI32,
			0xBF: newF64ReinterpretI64,
			0xC0: newI32Extend8s,
			0xC1: newI32Extend16s,
			0xC2: newI64Extend8s,
			0xC3: newI64Extend16s,
			0xC4: newI64Extend32s,
			0xFC: newTrunc,
		}
	}
	constructor, exist := opcodes[opcode]

	if !exist {
		bytes := []byte{opcode}
		nextBytes, _ := values.ReadNextBytes(reader, 14)
		return nil, errors.New(fmt.Sprintf("Expected a valid Opcode but got: 0x%s.\nNext bytes: %x", hex.EncodeToString(bytes), nextBytes))
	}

	return constructor(reader)
}

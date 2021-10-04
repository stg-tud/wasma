package instructions

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"wasma/pkg/wasmp/structures"
	values2 "wasma/pkg/wasmp/values"
)

type ControlInstructionBlock struct {
	opcodeValue    byte
	nameValue      string
	blocktypeValue string
	instrValues    []Instruction
	position       uint32
	executor       func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (controlInstructionBlock *ControlInstructionBlock) Opcode() byte {
	return controlInstructionBlock.opcodeValue
}
func (controlInstructionBlock *ControlInstructionBlock) Name() string {
	return controlInstructionBlock.nameValue
}
func (controlInstructionBlock *ControlInstructionBlock) Position() uint32 {
	return controlInstructionBlock.position
}
func (controlInstructionBlock *ControlInstructionBlock) Blocktype() (string, error) {
	return controlInstructionBlock.blocktypeValue, nil
}
func (controlInstructionBlock *ControlInstructionBlock) Instr() ([]Instruction, error) {
	return controlInstructionBlock.instrValues, nil
}
func (controlInstructionBlock *ControlInstructionBlock) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBlock *ControlInstructionBlock) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return controlInstructionBlock.executor
}

func (controlInstructionBlock *ControlInstructionBlock) ToString() string {
	if controlInstructionBlock.blocktypeValue == "empty" {
		return fmt.Sprintf("%v", controlInstructionBlock.nameValue)
	}
	return fmt.Sprintf("%v %v", controlInstructionBlock.nameValue, controlInstructionBlock.blocktypeValue)
}

type ControlInstructionIfElse struct {
	opcodeValue     byte
	nameValue       string
	blocktypeValue  string
	instrValues     []Instruction
	elseInstrValues []Instruction
	position        uint32
	executor        func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (controlInstructionIfElse *ControlInstructionIfElse) Opcode() byte {
	return controlInstructionIfElse.opcodeValue
}
func (controlInstructionIfElse *ControlInstructionIfElse) Name() string {
	return controlInstructionIfElse.nameValue
}
func (controlInstructionIfElse *ControlInstructionIfElse) Position() uint32 {
	return controlInstructionIfElse.position
}
func (controlInstructionIfElse *ControlInstructionIfElse) Blocktype() (string, error) {
	return controlInstructionIfElse.blocktypeValue, nil
}
func (controlInstructionIfElse *ControlInstructionIfElse) Instr() ([]Instruction, error) {
	return controlInstructionIfElse.instrValues, nil
}
func (controlInstructionIfElse *ControlInstructionIfElse) ElseInstr() ([]Instruction, error) {
	return controlInstructionIfElse.elseInstrValues, nil
}
func (controlInstructionIfElse *ControlInstructionIfElse) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionIfElse *ControlInstructionIfElse) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return controlInstructionIfElse.executor
}

func (controlInstructionIfElse *ControlInstructionIfElse) ToString() string {
	if controlInstructionIfElse.blocktypeValue == "empty" {
		return fmt.Sprintf("%v", controlInstructionIfElse.nameValue)
	}
	return fmt.Sprintf("%v %v", controlInstructionIfElse.nameValue, controlInstructionIfElse.blocktypeValue)
}

type ControlInstructionBr struct {
	opcodeValue   byte
	nameValue     string
	labelidxValue uint32
	position      uint32
	executor      func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (controlInstructionBr *ControlInstructionBr) Opcode() byte {
	return controlInstructionBr.opcodeValue
}
func (controlInstructionBr *ControlInstructionBr) Name() string {
	return controlInstructionBr.nameValue
}
func (controlInstructionBr *ControlInstructionBr) Position() uint32 {
	return controlInstructionBr.position
}
func (controlInstructionBr *ControlInstructionBr) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) Labelidx() (uint32, error) {
	return controlInstructionBr.labelidxValue, nil
}
func (controlInstructionBr *ControlInstructionBr) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBr *ControlInstructionBr) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return controlInstructionBr.executor
}

func (controlInstructionBr *ControlInstructionBr) ToString() string {
	return fmt.Sprintf("%v %v", controlInstructionBr.nameValue, controlInstructionBr.labelidxValue)
}

type ControlInstructionBrTable struct {
	opcodeValue       byte
	nameValue         string
	vecLabelidxValues []uint32
	labelidxValue     uint32
	position          uint32
	executor          func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (controlInstructionBrTable *ControlInstructionBrTable) Opcode() byte {
	return controlInstructionBrTable.opcodeValue
}
func (controlInstructionBrTable *ControlInstructionBrTable) Name() string {
	return controlInstructionBrTable.nameValue
}
func (controlInstructionBrTable *ControlInstructionBrTable) Position() uint32 {
	return controlInstructionBrTable.position
}
func (controlInstructionBrTable *ControlInstructionBrTable) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) VecLabelidx() ([]uint32, error) {
	return controlInstructionBrTable.vecLabelidxValues, nil
}
func (controlInstructionBrTable *ControlInstructionBrTable) Labelidx() (uint32, error) {
	return controlInstructionBrTable.labelidxValue, nil
}
func (controlInstructionBrTable *ControlInstructionBrTable) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionBrTable *ControlInstructionBrTable) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return controlInstructionBrTable.executor
}
func (controlInstructionBrTable *ControlInstructionBrTable) ToString() string {
	var labelIndices = ""
	for _, labelIdx := range controlInstructionBrTable.vecLabelidxValues {
		labelIndices = labelIndices + fmt.Sprintf("%v, ", labelIdx)
	}
	return fmt.Sprintf("%v %v%v", controlInstructionBrTable.nameValue, labelIndices, controlInstructionBrTable.labelidxValue)
}

type ControlInstructionCall struct {
	opcodeValue  byte
	nameValue    string
	funcidxValue uint32
	position     uint32
	executor     func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (controlInstructionCall *ControlInstructionCall) Opcode() byte {
	return controlInstructionCall.opcodeValue
}
func (controlInstructionCall *ControlInstructionCall) Name() string {
	return controlInstructionCall.nameValue
}
func (controlInstructionCall *ControlInstructionCall) Position() uint32 {
	return controlInstructionCall.position
}
func (controlInstructionCall *ControlInstructionCall) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) Funcidx() (uint32, error) {
	return controlInstructionCall.funcidxValue, nil
}
func (controlInstructionCall *ControlInstructionCall) Typeidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCall *ControlInstructionCall) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return controlInstructionCall.executor
}
func (controlInstructionCall *ControlInstructionCall) ToString() string {
	return fmt.Sprintf("%v %v", controlInstructionCall.nameValue, controlInstructionCall.funcidxValue)
}

type ControlInstructionCallIndirect struct {
	opcodeValue  byte
	nameValue    string
	typeidxValue uint32
	position     uint32
	executor     func(instrIdx uint32, instruction Instruction, environment structures.Environment)
}

func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Opcode() byte {
	return controlInstructionCallIndirect.opcodeValue
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Name() string {
	return controlInstructionCallIndirect.nameValue
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Position() uint32 {
	return controlInstructionCallIndirect.position
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Blocktype() (string, error) {
	return "", attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Instr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) ElseInstr() ([]Instruction, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) VecLabelidx() ([]uint32, error) {
	return nil, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Labelidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Funcidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Typeidx() (uint32, error) {
	return controlInstructionCallIndirect.typeidxValue, nil
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Localidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Globalidx() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Align() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Offset() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) I32() (int32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) I64() (int64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) F32() (float32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) F64() (float64, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) TruncOpcode() (uint32, error) {
	return 0, attributeNotAvailable
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) Executor() func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
	return controlInstructionCallIndirect.executor
}
func (controlInstructionCallIndirect *ControlInstructionCallIndirect) ToString() string {
	return fmt.Sprintf("%v 0, %v", controlInstructionCallIndirect.nameValue, controlInstructionCallIndirect.typeidxValue)
}

func newUnreachable(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{
		0x00,
		"unreachable",
		values2.GetPosition(),
		func(instrIdx uint32, instruction Instruction, environment structures.Environment) {},
	}, nil
}

func newNop(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{
		0x01,
		"nop",
		values2.GetPosition(),
		func(instrIdx uint32, instruction Instruction, environment structures.Environment) {},
	}, nil
}

func readBlocktype(reader io.Reader) (string, error) {
	nextByte, err := values2.ReadNextByte(reader)
	if err != nil {
		return "", err
	}
	if nextByte == 0x40 {
		return "empty", nil
	} else {
		valtype, err := getValtype(nextByte)
		if err == nil {
			return valtype, nil
		} else {
			valueType, err := values2.ReadS33(nextByte, reader)
			if err != nil {
				return "", err
			}
			return strconv.FormatInt(int64(valueType), 10), nil
		}
	}
}

func newControlInstructionBlock(opcode byte, name string, reader io.Reader) (Instruction, error) {
	controlInstructionBlock := new(ControlInstructionBlock)
	controlInstructionBlock.position = values2.GetPosition()
	controlInstructionBlock.opcodeValue = opcode
	controlInstructionBlock.nameValue = name
	controlInstructionBlock.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {}

	blocktype, err := readBlocktype(reader)

	if err != nil {
		return nil, err
	}

	controlInstructionBlock.blocktypeValue = blocktype
	for {
		nextByte, err := values2.ReadNextByte(reader)

		if err != nil {
			return nil, err
		}

		if nextByte == 0x0B {
			values2.IncrementEndCounter()
			break
		}

		instruction, err := mapOpcode(nextByte, reader)
		if err != nil {
			return nil, err
		}
		controlInstructionBlock.instrValues = append(controlInstructionBlock.instrValues, instruction)
	}
	return controlInstructionBlock, nil
}

func newBlock(reader io.Reader) (Instruction, error) {
	return newControlInstructionBlock(0x02, "block", reader)
}

func newLoop(reader io.Reader) (Instruction, error) {
	return newControlInstructionBlock(0x03, "loop", reader)
}

func newIfElse(reader io.Reader) (Instruction, error) {
	ifInstr := new(ControlInstructionIfElse)
	ifInstr.position = values2.GetPosition()
	ifInstr.opcodeValue = 0x04
	ifInstr.nameValue = "if"
	ifInstr.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
		variable := environment.Pop()
		environment.AddInput(instrIdx, variable)
	}

	var elseBranch bool = false

	blocktype, err := readBlocktype(reader)
	if err != nil {
		return nil, err
	}
	ifInstr.blocktypeValue = blocktype

	for {
		nextByte, err := values2.ReadNextByte(reader)
		if err != nil {
			return nil, err
		}

		if nextByte == 0x0B {
			values2.IncrementEndCounter()
			break
		} else if nextByte == 0x05 {
			values2.IncrementElseCounter()
			elseBranch = true
			nextByte, err = values2.ReadNextByte(reader)
			if err != nil {
				return nil, err
			}
			if nextByte == 0x0B && blocktype == "empty" {
				values2.IncrementEndCounter()
				break
			}
		}

		instruction, err := mapOpcode(nextByte, reader)
		if err != nil {
			return nil, err
		}

		if elseBranch {
			ifInstr.elseInstrValues = append(ifInstr.elseInstrValues, instruction)
		} else {
			ifInstr.instrValues = append(ifInstr.instrValues, instruction)
		}
	}

	return ifInstr, nil
}

func newControlInstructionBr(opcode byte, name string, reader io.Reader) (Instruction, error) {
	br := new(ControlInstructionBr)
	br.position = values2.GetPosition()
	br.opcodeValue = opcode
	br.nameValue = name
	if name == "br_if" {
		br.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
			variable := environment.Pop()
			environment.AddInput(instrIdx, variable)
		}
	} else {
		br.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {}
	}
	intValue, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}
	br.labelidxValue = intValue

	return br, nil
}

func newBr(reader io.Reader) (Instruction, error) {
	return newControlInstructionBr(0x0C, "br", reader)
}

func newBrIf(reader io.Reader) (Instruction, error) {
	return newControlInstructionBr(0x0D, "br_if", reader)
}

func newBrTable(reader io.Reader) (Instruction, error) {
	brTable := new(ControlInstructionBrTable)
	brTable.position = values2.GetPosition()
	brTable.opcodeValue = 0x0E
	brTable.nameValue = "br_table"
	brTable.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
		variable := environment.Pop()
		environment.AddInput(instrIdx, variable)
	}

	vecLen, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		lidx, err := values2.ReadU32(reader)
		if err != nil {
			return nil, err
		}
		brTable.vecLabelidxValues = append(brTable.vecLabelidxValues, lidx)
	}
	labelidx, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}
	brTable.labelidxValue = labelidx

	return brTable, nil
}

func newReturn(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{
		0x0F,
		"return",
		values2.GetPosition(),
		func(instrIdx uint32, instruction Instruction, environment structures.Environment) {},
	}, nil
}

func newCall(reader io.Reader) (Instruction, error) {
	call := new(ControlInstructionCall)
	call.position = values2.GetPosition()
	call.opcodeValue = 0x10
	call.nameValue = "call"
	call.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
		funcIdx, err := instruction.Funcidx()
		if err != nil {
			log.Fatalf(err.Error())
		}

		functionType, err := environment.GetNumOfFuncParameterFuncIdx(funcIdx)
		if err != nil {
			log.Fatalf(err.Error())
		}

		parameterTypes := functionType.ParameterTypes

		for i := len(parameterTypes) - 1; i >= 0; i-- {
			paramType := parameterTypes[i]
			variable := environment.Pop()
			if variable.VariableDataType != paramType {
				log.Fatalf(fmt.Sprintf("function call error: expected %v but got %v, instrIdx: %v", paramType, variable.VariableDataType, instrIdx))
			}
			environment.AddInput(instrIdx, variable)
		}

		for _, resultType := range functionType.ResultTypes {
			newVariable := environment.NewVariable(resultType)
			environment.Push(newVariable)
			environment.AddOutput(instrIdx, newVariable)
		}
	}
	funcidx, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}
	call.funcidxValue = funcidx

	return call, nil
}

func newCallIndirect(reader io.Reader) (Instruction, error) {
	callIndirect := new(ControlInstructionCallIndirect)
	callIndirect.position = values2.GetPosition()
	callIndirect.opcodeValue = 0x11
	callIndirect.nameValue = "call_indirect"
	callIndirect.executor = func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
		typeIdx, err := instruction.Typeidx()
		if err != nil {
			log.Fatalf(err.Error())
		}

		functionType, err := environment.GetNumOfFuncParameterTypeIdx(typeIdx)
		if err != nil {
			log.Fatalf(err.Error())
		}

		// pop funcref
		funcref := environment.Pop()
		if funcref.VariableDataType != "i32" {
			log.Fatalf(fmt.Sprintf("indirect function call error: expected funcref of type i32 but got %v, instrIdx: %v", funcref.VariableDataType, instrIdx))
		}
		environment.AddInput(instrIdx, funcref)

		parameterTypes := functionType.ParameterTypes

		for i := len(parameterTypes) - 1; i >= 0; i-- {
			paramType := parameterTypes[i]
			variable := environment.Pop()
			if variable.VariableDataType != paramType {
				log.Fatalf(fmt.Sprintf("indirect function call error: expected %v but got %v, instrIdx: %v", paramType, variable.VariableDataType, instrIdx))
			}
			environment.AddInput(instrIdx, variable)
		}

		for _, resultType := range functionType.ResultTypes {
			newVariable := environment.NewVariable(resultType)
			environment.Push(newVariable)
			environment.AddOutput(instrIdx, newVariable)
		}
	}
	typeidx, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}
	callIndirect.typeidxValue = typeidx

	nextByte, err := values2.ReadNextByte(reader)
	if err != nil {
		return nil, err
	}
	if nextByte != 0x00 {
		return nil, errors.New(fmt.Sprintf("Expected 0x00 while reading call_indirect instraction but got: %x", nextByte))
	}

	return callIndirect, nil
}

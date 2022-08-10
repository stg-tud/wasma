package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	instructions2 "wasma/pkg/wasmp/instructions"
	"wasma/pkg/wasmp/types"
	"wasma/pkg/wasmp/values"
)

type TestCaseGlobalSection struct {
	reader        io.Reader
	expectedValue GlobalSection
}

func TestNewGlobalSection(t *testing.T) {
	numericInstructionI32 := new(instructions2.NumericInstructionI32)
	numericInstructionI32.OpcodeValue = 0x41
	numericInstructionI32.NameValue = "i32.const"
	numericInstructionI32.I32Value = 1
	numericInstructionI32.PositionValue = 5

	// Positive test cases
	positiveTestCases := []TestCaseGlobalSection{{bytes.NewReader([]byte{0x06, 0x06, 0x01, 0x7F, 0x00, 0x41, 0x01, 0x0B}), GlobalSection{0x06, 6, map[uint32]Global{0: {&types.GlobalType{ValType: "i32", Mut: 0x00}, &instructions2.Expr{Instructions: []instructions2.Instruction{numericInstructionI32}}}}, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x06, sectionId, t)
		globalSection, _ := NewGlobalSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, globalSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, globalSection.Size, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.Globals), len(globalSection.Globals), t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, globalSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, globalSection.StartContent, t)
		for i, global := range testCase.expectedValue.Globals {
			test_utilities2.CompareString(global.GlobalType.ValType, globalSection.Globals[i].GlobalType.ValType, t)
			test_utilities2.CompareByte(global.GlobalType.Mut, globalSection.Globals[i].GlobalType.Mut, t)
			test_utilities2.CompareInt(len(global.Expr.Instructions), len(globalSection.Globals[i].Expr.Instructions), t)
			for j, instruction := range global.Expr.Instructions {
				test_utilities2.CompareString(instruction.Name(), globalSection.Globals[i].Expr.Instructions[j].Name(), t)
			}
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewGlobalSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

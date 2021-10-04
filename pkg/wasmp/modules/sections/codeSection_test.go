package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	instructions2 "wasma/pkg/wasmp/instructions"
	"wasma/pkg/wasmp/values"
)

type TestCaseCodeSection struct {
	reader        io.Reader
	expectedValue CodeSection
}

func TestNewCodeSection(t *testing.T) {
	numericInstructionI32 := new(instructions2.NumericInstructionI32)
	numericInstructionI32.OpcodeValue = 0x41
	numericInstructionI32.NameValue = "i32.const"
	numericInstructionI32.I32Value = 42
	numericInstructionI32.PositionValue = 5

	// Positive test cases
	positiveTestCases := []TestCaseCodeSection{{bytes.NewReader([]byte{0x0A, 0x06, 0x01, 0x04, 0x00, 0x41, 0x2A, 0x0B}),
		CodeSection{0x0A, 6, map[uint32]*Code{0: {4, &Function{[]*Local{}, &instructions2.Expr{[]instructions2.Instruction{numericInstructionI32}}}}}, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x0A, sectionId, t)
		codeSection, _ := NewCodeSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, codeSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, codeSection.Size, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, codeSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, codeSection.StartContent, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.Codes), len(codeSection.Codes), t)
		for i, code := range testCase.expectedValue.Codes {
			test_utilities2.CompareUInt32(code.Size, codeSection.Codes[i].Size, t)
			test_utilities2.CompareInt(len(code.Function.Locals), len(codeSection.Codes[i].Function.Locals), t)
			for j, local := range code.Function.Locals {
				test_utilities2.CompareUInt32(local.N, codeSection.Codes[i].Function.Locals[j].N, t)
				test_utilities2.CompareString(local.ValType, codeSection.Codes[i].Function.Locals[j].ValType, t)
			}
			test_utilities2.CompareInt(len(code.Function.Expr.Instructions), len(codeSection.Codes[i].Function.Expr.Instructions), t)
			for j, instruction := range code.Function.Expr.Instructions {
				test_utilities2.CompareString(instruction.Name(), codeSection.Codes[i].Function.Expr.Instructions[j].Name(), t)
			}
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewCodeSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

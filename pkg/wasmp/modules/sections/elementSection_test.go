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

type TestCaseElementSection struct {
	reader        io.Reader
	expectedValue ElementSection
}

func TestNewElementSection(t *testing.T) {
	numericInstructionI32 := new(instructions2.NumericInstructionI32)
	numericInstructionI32.OpcodeValue = 0x41
	numericInstructionI32.NameValue = "i32.const"
	numericInstructionI32.I32Value = 1
	numericInstructionI32.PositionValue = 4

	// Positive test cases
	positiveTestCases := []TestCaseElementSection{{bytes.NewReader([]byte{0x09, 0x06, 0x01, 0x00, 0x41, 0x01, 0x0B, 0x00}),
		ElementSection{
			0x09,
			6,
			[]*Element{{0, &instructions2.Expr{Instructions: []instructions2.Instruction{numericInstructionI32}}, []uint32{}}},
			0,
			2,
			make(map[uint32]bool)}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x09, sectionId, t)
		elementSection, _ := NewElementSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, elementSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, elementSection.Size, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.Elements), len(elementSection.Elements), t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, elementSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, elementSection.StartContent, t)
		for i, element := range testCase.expectedValue.Elements {
			test_utilities2.CompareUInt32(element.TableIdx, elementSection.Elements[i].TableIdx, t)
			test_utilities2.CompareUInt32s(element.FuncIdxs, elementSection.Elements[i].FuncIdxs, t)
			test_utilities2.CompareInt(len(element.Expr.Instructions), len(elementSection.Elements[i].Expr.Instructions), t)
			for j, instruction := range element.Expr.Instructions {
				test_utilities2.CompareString(instruction.Name(), elementSection.Elements[i].Expr.Instructions[j].Name(), t)
			}
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewElementSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

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

type TestCaseDataSection struct {
	reader        io.Reader
	expectedValue DataSection
}

func TestNewDataSection(t *testing.T) {
	numericInstructionI32 := new(instructions2.NumericInstructionI32)
	numericInstructionI32.OpcodeValue = 0x41
	numericInstructionI32.NameValue = "i32.const"
	numericInstructionI32.I32Value = 1
	numericInstructionI32.PositionValue = 4

	// Positive test cases
	positiveTestCases := []TestCaseDataSection{{bytes.NewReader([]byte{0x0B, 0x06, 0x01, 0x00, 0x41, 0x01, 0x0B, 0x00}), DataSection{0x0B, 6, []*Data{{0, &instructions2.Expr{Instructions: []instructions2.Instruction{numericInstructionI32}}, []byte{}}}, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x0B, sectionId, t)
		dataSection, _ := NewDataSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, dataSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, dataSection.Size, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.Datas), len(dataSection.Datas), t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, dataSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, dataSection.StartContent, t)
		for i, data := range testCase.expectedValue.Datas {
			test_utilities2.CompareUInt32(data.MemIdx, dataSection.Datas[i].MemIdx, t)
			test_utilities2.CompareBytes(data.Bytes, dataSection.Datas[i].Bytes, t)
			test_utilities2.CompareInt(len(data.Expr.Instructions), len(dataSection.Datas[i].Expr.Instructions), t)
			for j, instruction := range data.Expr.Instructions {
				test_utilities2.CompareString(instruction.Name(), dataSection.Datas[i].Expr.Instructions[j].Name(), t)
			}
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewDataSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

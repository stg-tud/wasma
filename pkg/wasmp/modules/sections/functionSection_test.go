package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	"wasma/pkg/wasmp/values"
)

type TestCaseFunctionSection struct {
	reader        io.Reader
	expectedValue FunctionSection
}

func TestNewFunctionSection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseFunctionSection{{bytes.NewReader([]byte{0x03, 0x02, 0x01, 0x02}), FunctionSection{0x03, 2, map[uint32]uint32{0: 2}, make(map[uint32][]uint32), 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x03, sectionId, t)
		functionSection, _ := NewFunctionSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, functionSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, functionSection.Size, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, functionSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, functionSection.StartContent, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.TypeIdxs), len(functionSection.TypeIdxs), t)
		for i, typeIdx := range testCase.expectedValue.TypeIdxs {
			test_utilities2.CompareUInt32(typeIdx, functionSection.TypeIdxs[i], t)
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewFunctionSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

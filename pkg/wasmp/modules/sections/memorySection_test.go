package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	"wasma/pkg/wasmp/types"
	"wasma/pkg/wasmp/values"
)

type TestCaseMemorySection struct {
	reader        io.Reader
	expectedValue MemorySection
}

func TestNewMemorySection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseMemorySection{{bytes.NewReader([]byte{0x05, 0x04, 0x01, 0x01, 0x00, 0x00}), MemorySection{0x05, 4, map[uint32]*types.Limit{0: {Min: 0, Max: 0, Type: 0x01}}, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x05, sectionId, t)
		memorySection, _ := NewMemorySection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, memorySection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, memorySection.Size, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.MemTypes), len(memorySection.MemTypes), t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, memorySection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, memorySection.StartContent, t)
		for i, memTyp := range testCase.expectedValue.MemTypes {
			test_utilities2.CompareUInt32(memTyp.Min, memorySection.MemTypes[i].Min, t)
			test_utilities2.CompareUInt32(memTyp.Max, memorySection.MemTypes[i].Max, t)
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewMemorySection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

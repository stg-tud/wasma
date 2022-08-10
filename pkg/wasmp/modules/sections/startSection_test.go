package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	"wasma/pkg/wasmp/values"
)

type TestCaseStartSection struct {
	reader        io.Reader
	expectedValue StartSection
}

func TestNewStartSection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseStartSection{{bytes.NewReader([]byte{0x08, 0x01, 0x00}), StartSection{0x08, 1, 0x00, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x08, sectionId, t)
		startSection, _ := NewStartSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, startSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, startSection.Size, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.FuncIdx, startSection.FuncIdx, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, startSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, startSection.StartContent, t)
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewStartSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

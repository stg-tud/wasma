package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	"wasma/pkg/wasmp/values"
)

type TestCaseCustomSection struct {
	reader        io.Reader
	expectedValue CustomSection
}

func TestNewCustomSection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseCustomSection{{bytes.NewReader([]byte{0x00, 0x0A, 0x05, 0x63, 0x54, 0x65, 0x73, 0x74, 0x00, 0x01, 0x02, 0x03}), CustomSection{0x00, 10, "cTest", []byte{0x00, 0x01, 0x02, 0x03}, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x00, sectionId, t)
		customSection, _ := NewCustomSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, customSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, customSection.Size, t)
		test_utilities2.CompareString(testCase.expectedValue.Name, customSection.Name, t)
		test_utilities2.CompareBytes(testCase.expectedValue.CustomBytes, customSection.CustomBytes, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, customSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, customSection.StartContent, t)
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewCustomSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

package sections

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseCompletelyRead struct {
	section       string
	size          uint32
	startByte     uint32
	currentByte   uint32
	expectedValue error
}

func TestContentCompletelyRead(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseCompletelyRead{{"test", 10, 1, 10, nil}}

	for _, testCase := range positiveTestCases {
		test_utilities2.ErrorNil(contentCompletelyRead(testCase.section, testCase.size, testCase.startByte, testCase.currentByte), t)
	}

	// Negative test cases
	negativeTestCases := []TestCaseCompletelyRead{
		{"test", 10, 1, 9, errors.New("not all Bytes of the test section were read")},
		{"test", 10, 1, 12, errors.New("more Bytes were read than the current test section contains")}}

	for _, testCase := range negativeTestCases {
		test_utilities2.CompareErrorMessage(testCase.expectedValue, contentCompletelyRead(testCase.section, testCase.size, testCase.startByte, testCase.currentByte), t)
	}
}

type TestCaseReadNextSection struct {
	reader         io.Reader
	n              uint32
	expectedValues []byte
}

func TestReadNextSection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseReadNextSection{
		{bytes.NewReader([]byte{0x01, 0x02, 0x03}), 3, []byte{0x01, 0x02, 0x03}},
		{bytes.NewReader([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}), 7, []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}},
	}

	for _, testCase := range positiveTestCases {
		sectionReader, _ := readNextSection(testCase.reader, testCase.n)
		sectionBytes, _ := ioutil.ReadAll(sectionReader)
		test_utilities2.CompareBytes(testCase.expectedValues, sectionBytes, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, err := readNextSection(testCase.Reader, 1)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

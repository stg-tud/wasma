package values

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseBytes struct {
	reader         io.Reader
	n              uint32
	expectedValues []byte
}

type TestCaseByte struct {
	reader        io.Reader
	expectedValue byte
}

func TestReadNextBytes(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseBytes{
		{bytes.NewReader([]byte{0x00}), 1, []byte{0x00}},
		{bytes.NewReader([]byte{0x01, 0x02, 0x03}), 3, []byte{0x01, 0x02, 0x03}},
	}

	for _, testCase := range positiveTestCases {
		nextBytes, _ := ReadNextBytes(testCase.reader, testCase.n)
		test_utilities2.CompareBytes(testCase.expectedValues, nextBytes, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, err := ReadNextBytes(testCase.Reader, 1)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

func TestReadNextByte(t *testing.T) {
	// Positive test case
	positiveTestCases := []TestCaseByte{
		{bytes.NewReader([]byte{0x00}), 0x00},
		{bytes.NewReader([]byte{0x01}), 0x01},
		{bytes.NewReader([]byte{0x02}), 0x02},
		{bytes.NewReader([]byte{0x03}), 0x03}}

	for _, testCase := range positiveTestCases {
		nextByte, _ := ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue, nextByte, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, err := ReadNextByte(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

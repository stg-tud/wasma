package types

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseMemoryType struct {
	reader        io.Reader
	expectedValue Limit
}

func TestNewMemoryType(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseMemoryType{
		{bytes.NewReader([]byte{0x00, 0x01}), Limit{1, 0, 0x00}},
		{bytes.NewReader([]byte{0x01, 0x01, 0x64}), Limit{1, 100, 0x01}}}

	for _, testCase := range positiveTestCases {
		limit, _ := NewMemoryType(testCase.reader)
		test_utilities2.CompareUInt32(testCase.expectedValue.Min, limit.Min, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Max, limit.Max, t)
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF")},
		{Reader: bytes.NewReader([]byte{0x03, 0x01}), Err: errors.New("Error while reading memory type. Expected 0x00 or 0x01 but got: 3")},
		{Reader: bytes.NewReader([]byte{0x01, 0x01}), Err: errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, err := NewMemoryType(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

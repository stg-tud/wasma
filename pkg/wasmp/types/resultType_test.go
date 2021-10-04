package types

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseResultType struct {
	reader         io.Reader
	expectedValues []string
}

func TestNewResultType(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseResultType{
		{bytes.NewReader([]byte{0x04, 0x7F, 0x7E, 0x7D, 0x7C}), []string{"i32", "i64", "f32", "f64"}}}

	for _, testCase := range positiveTestCases {
		resultType, _ := NewResultType(testCase.reader)
		test_utilities2.CompareStrings(testCase.expectedValues, resultType.ValType, t)
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF")},
		{bytes.NewReader([]byte{0x05, 0x7F, 0x7E, 0x7D, 0x7C}), errors.New("EOF")},
		{bytes.NewReader([]byte{0x04, 0x7F, 0x7E, 0x77, 0x7C}), errors.New("Error while reading result type. Unexpected type got: 77. Valid types: 0x7F, 0x7E, 0x7D and 0x7C.")}}

	for _, testCase := range negativeTestCases {
		_, err := NewResultType(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

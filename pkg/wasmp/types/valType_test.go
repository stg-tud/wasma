package types

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseValueType struct {
	reader        io.Reader
	expectedValue string
}

func TestNewValueType(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseValueType{
		{bytes.NewReader([]byte{0x7F}), "i32"},
		{bytes.NewReader([]byte{0x7E}), "i64"},
		{bytes.NewReader([]byte{0x7D}), "f32"},
		{bytes.NewReader([]byte{0x7C}), "f64"},
	}

	for _, testCase := range positiveTestCases {
		valueType, _ := NewValType(testCase.reader)
		test_utilities2.CompareString(testCase.expectedValue, valueType, t)

	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF")},
		{Reader: bytes.NewReader([]byte{0x00}), Err: errors.New("Error while reading result type. Unexpected type got: 0. Valid types: 0x7F, 0x7E, 0x7D and 0x7C.")}}

	for _, testCase := range negativeTestCases {
		_, err := NewValType(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

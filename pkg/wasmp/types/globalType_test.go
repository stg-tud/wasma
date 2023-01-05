package types

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseGlobalType struct {
	reader        io.Reader
	expectedValue GlobalType
}

func TestNewGlobalType(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseGlobalType{
		{bytes.NewReader([]byte{0x7F, 0x00}), GlobalType{"i32", 0x00}},
		{bytes.NewReader([]byte{0x7E, 0x01}), GlobalType{"i64", 0x01}}}

	for _, testCase := range positiveTestCases {
		globalType, _ := NewGlobalType(testCase.reader)
		test_utilities2.CompareString(testCase.expectedValue.ValType, globalType.ValType, t)
		test_utilities2.CompareByte(testCase.expectedValue.Mut, globalType.Mut, t)
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF")},
		{Reader: bytes.NewReader([]byte{0x00, 0x00}), Err: errors.New("Error while reading result type. Unexpected type got: 0. Valid types: 0x7F, 0x7E, 0x7D and 0x7C.")},
		{Reader: bytes.NewReader([]byte{0x7F, 0x02}), Err: errors.New("Error while reading global type Mut value. Expected 0x00 or 0x01, but got: 2.")}}

	for _, testCase := range negativeTestCases {
		_, err := NewGlobalType(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

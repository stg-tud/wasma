package types

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseFunctionType struct {
	reader        io.Reader
	expectedValue FunctionType
}

func TestNewFunctionType(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseFunctionType{
		{bytes.NewReader([]byte{0x60, 0x04, 0x7F, 0x7E, 0x7D, 0x7C, 0x00}), FunctionType{[]string{"i32", "i64", "f32", "f64"}, []string{}}},
		{bytes.NewReader([]byte{0x60, 0x00, 0x04, 0x7F, 0x7E, 0x7D, 0x7C}), FunctionType{[]string{}, []string{"i32", "i64", "f32", "f64"}}},
		{bytes.NewReader([]byte{0x60, 0x02, 0x7F, 0x7F, 0x01, 0x7F}), FunctionType{[]string{"i32", "i32"}, []string{"i32"}}}}

	for _, testCase := range positiveTestCases {
		functionType, _ := NewFunctionType(testCase.reader)
		test_utilities2.CompareStrings(testCase.expectedValue.ParameterTypes, functionType.ParameterTypes, t)
		test_utilities2.CompareStrings(testCase.expectedValue.ResultTypes, functionType.ResultTypes, t)
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF")},
		{Reader: bytes.NewReader([]byte{0x00, 0x04, 0x7F, 0x7E, 0x7D, 0x7C, 0x00}), Err: errors.New("Error while reading type section. Expected 0x60 but got: 0")},
		{Reader: bytes.NewReader([]byte{0x60, 0x04, 0x7F, 0x7E, 0x7D, 0x7C}), Err: errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, err := NewFunctionType(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

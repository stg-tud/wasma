package values

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseFloat32 struct {
	reader        io.Reader
	expectedValue float32
}

type TestCaseFloat64 struct {
	reader        io.Reader
	expectedValue float64
}

func TestReadF32(t *testing.T) {
	// Positive test case
	positiveTestCases := []TestCaseFloat32{
		{bytes.NewReader([]byte{0x14, 0xAE, 0x29, 0xC2}), -42.42},
		{bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00}), 0.0},
		{bytes.NewReader([]byte{0x14, 0xAE, 0x29, 0x42}), 42.42}}

	for _, testCase := range positiveTestCases {
		float32Value, _ := ReadF32(testCase.reader)
		test_utilities2.CompareFloat32(testCase.expectedValue, float32Value, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, err := ReadF32(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

func TestReadF64(t *testing.T) {
	// Positive test case
	positiveTestCases := []TestCaseFloat64{
		{bytes.NewReader([]byte{0xC2, 0x01, 0x00, 0x00, 0x00, 0x00, 0x80, 0xC2}), -2199023255552.2199023255552},
		{bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}), 0.0},
		{bytes.NewReader([]byte{0xC2, 0x01, 0x00, 0x00, 0x00, 0x00, 0x80, 0x42}), 2199023255552.2199023255552}}

	for _, testCase := range positiveTestCases {
		float64Value, _ := ReadF64(testCase.reader)
		test_utilities2.CompareFloat64(testCase.expectedValue, float64Value, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, err := ReadF64(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

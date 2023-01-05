package types

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseTableType struct {
	reader        io.Reader
	expectedValue TableType
}

func TestNewTableType(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseTableType{
		{bytes.NewReader([]byte{0x70, 0x00, 0x01}), TableType{&Limit{1, 0, 0x00}}},
		{bytes.NewReader([]byte{0x70, 0x01, 0x01, 0x64}), TableType{&Limit{1, 100, 0x01}}}}

	for _, testCase := range positiveTestCases {
		tableType, _ := NewTableType(testCase.reader)
		test_utilities2.CompareUInt32(testCase.expectedValue.Limit.Min, tableType.Limit.Min, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Limit.Max, tableType.Limit.Max, t)
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF")},
		{Reader: bytes.NewReader([]byte{0x03, 0x01}), Err: errors.New("Error while reading table type. Expected 0x70 but got: 3")},
		{Reader: bytes.NewReader([]byte{0x70, 0x01, 0x01}), Err: errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, err := NewTableType(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

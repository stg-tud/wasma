package values

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
)

type TestCaseUInt32 struct {
	reader        io.Reader
	expectedValue uint32
}

type TestCaseUInt64 struct {
	reader        io.Reader
	expectedValue uint64
}

type TestCaseInt32 struct {
	reader        io.Reader
	expectedValue int32
}

type TestCaseSInt33 struct {
	firstByte     byte
	reader        io.Reader
	expectedValue int64
}

type TestCaseInt64 struct {
	reader        io.Reader
	expectedValue int64
}

func TestReadU32(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseUInt32{
		// min
		{bytes.NewReader([]byte{0x00}), 0},
		{bytes.NewReader([]byte{0x2A}), 42},
		{bytes.NewReader([]byte{0xE9, 0x8C, 0x18}), 394857},
		// max
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x0F}), 4294967295}}

	for _, testCase := range positiveTestCases {
		actualValue, _ := ReadU32(testCase.reader)
		test_utilities2.CompareUInt32(testCase.expectedValue, actualValue, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		// empty reader
		{bytes.NewReader([]byte{}), errors.New("EOF")},
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, actualErr := ReadU32(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, actualErr, t)
	}
}

func TestReadU64(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseUInt64{
		// min
		{bytes.NewReader([]byte{0x00}), 0},
		{bytes.NewReader([]byte{0x2A}), 42},
		{bytes.NewReader([]byte{0xF4, 0xFA, 0xB6, 0x95, 0x91, 0xDE, 0x96, 0x8D, 0x01}), 79475934879399284},
		// max
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0F}), 18446744073709551615}}

	for _, testCase := range positiveTestCases {
		actualValue, _ := ReadU64(testCase.reader)
		test_utilities2.CompareUInt64(testCase.expectedValue, actualValue, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		// empty reader
		{bytes.NewReader([]byte{}), errors.New("EOF")},
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, actualErr := ReadU64(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, actualErr, t)
	}
}

func TestReadS32(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseInt32{
		// min
		{bytes.NewReader([]byte{0x80, 0x80, 0x80, 0x80, 0x78}), -2147483648},
		{bytes.NewReader([]byte{0x97, 0xF3, 0x67}), -394857},
		{bytes.NewReader([]byte{0x00}), 0},
		{bytes.NewReader([]byte{0x2A}), 42},
		{bytes.NewReader([]byte{0xE9, 0x8C, 0x18}), 394857},
		// max
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x07}), 2147483647}}

	for _, testCase := range positiveTestCases {
		actualValue, _ := ReadS32(testCase.reader)
		test_utilities2.CompareInt32(testCase.expectedValue, actualValue, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		// empty reader
		{bytes.NewReader([]byte{}), errors.New("EOF")},
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, actualErr := ReadS32(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, actualErr, t)
	}
}

func TestReadS33(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseSInt33{
		// min
		{0x80, bytes.NewReader([]byte{0x80, 0x80, 0x80, 0x70}), -4294967296},
		{0x97, bytes.NewReader([]byte{0xF3, 0xE7, 0xFF, 0x1F}), -394857},
		{0x00, bytes.NewReader([]byte{}), 0},
		{0x2A, bytes.NewReader([]byte{}), 42},
		{0xE9, bytes.NewReader([]byte{0x8C, 0x18}), 394857},
		// max
		{0xFF, bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x07}), 4294967295}}

	for _, testCase := range positiveTestCases {
		actualValue, _ := ReadS33(testCase.firstByte, testCase.reader)
		test_utilities2.CompareInt64(testCase.expectedValue, actualValue, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, actualErr := ReadS33(0xFF, testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, actualErr, t)
	}
}

func TestReadS64(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseInt64{
		// min
		{bytes.NewReader([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x7F}), -9223372036854775808},
		{bytes.NewReader([]byte{0x8C, 0x85, 0xC9, 0xEA, 0xEE, 0xA1, 0xE9, 0xF2, 0x7E}), -79475934879399284},
		{bytes.NewReader([]byte{0x00}), 0},
		{bytes.NewReader([]byte{0x2A}), 42},
		{bytes.NewReader([]byte{0xF4, 0xFA, 0xB6, 0x95, 0x91, 0xDE, 0x96, 0x8D, 0x01}), 79475934879399284},
		// max
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00}), 9223372036854775807}}

	for _, testCase := range positiveTestCases {
		actualValue, _ := ReadS64(testCase.reader)
		test_utilities2.CompareInt64(testCase.expectedValue, actualValue, t)
	}

	// Negative test case
	negativeTestCases := []test_utilities2.TestCaseError{
		// empty reader
		{bytes.NewReader([]byte{}), errors.New("EOF")},
		{bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}), errors.New("EOF")}}

	for _, testCase := range negativeTestCases {
		_, actualErr := ReadS64(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, actualErr, t)
	}
}

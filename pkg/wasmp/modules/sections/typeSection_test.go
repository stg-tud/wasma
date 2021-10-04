package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	"wasma/pkg/wasmp/types"
	"wasma/pkg/wasmp/values"
)

type TestCaseTypeSection struct {
	reader        io.Reader
	expectedValue TypeSection
}

func TestNewTypeSection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseTypeSection{{bytes.NewReader([]byte{0x01, 0x13, 0x03, 0x60, 0x04, 0x7F, 0x7E, 0x7D, 0x7C, 0x04, 0x7F, 0x7E, 0x7D, 0x7C, 0x60, 0x00, 0x00, 0x60, 0x00, 0x01, 0x7F}),
		TypeSection{0x01, 19, map[uint32]types.FunctionType{
			0: {[]string{"i32", "i64", "f32", "f64"}, []string{"i32", "i64", "f32", "f64"}},
			1: {[]string{}, []string{}},
			2: {[]string{}, []string{"i32"}}}, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x01, sectionId, t)
		typeSection, _ := NewTypeSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, typeSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, typeSection.Size, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.FunctionTypes), len(typeSection.FunctionTypes), t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, typeSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, typeSection.StartContent, t)
		for functionIdx, functionType := range testCase.expectedValue.FunctionTypes {
			test_utilities2.CompareStrings(functionType.ParameterTypes, typeSection.FunctionTypes[functionIdx].ParameterTypes, t)
			test_utilities2.CompareStrings(functionType.ResultTypes, typeSection.FunctionTypes[functionIdx].ResultTypes, t)
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewTypeSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	"wasma/pkg/wasmp/values"
)

type TestCaseExportSection struct {
	reader        io.Reader
	expectedValue ExportSection
}

func TestNewExportSection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseExportSection{{bytes.NewReader([]byte{0x07, 0x09, 0x01, 0x05, 0x65, 0x54, 0x65, 0x73, 0x74, 0x00, 0x01}), ExportSection{0x07, 9, []*Export{{"eTest", &ExportDesc{ExportType: 0x00, FuncIdx: 0x01}}}, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x07, sectionId, t)
		exportSection, _ := NewExportSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, exportSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, exportSection.Size, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.Exports), len(exportSection.Exports), t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, exportSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, exportSection.StartContent, t)
		for i, export := range testCase.expectedValue.Exports {
			test_utilities2.CompareString(export.Name, exportSection.Exports[i].Name, t)
			expectedValue, _, _ := export.ExportDesc.GetIdx()
			actualValue, _, _ := exportSection.Exports[i].ExportDesc.GetIdx()
			test_utilities2.CompareUInt32(expectedValue, actualValue, t)
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewExportSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

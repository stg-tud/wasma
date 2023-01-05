package sections

import (
	"bytes"
	"errors"
	"io"
	"testing"
	test_utilities2 "wasma/internal/test_utilities"
	types2 "wasma/pkg/wasmp/types"
	"wasma/pkg/wasmp/values"
)

type TestCaseTableSection struct {
	reader        io.Reader
	expectedValue TableSection
}

func TestNewTableSection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseTableSection{{bytes.NewReader([]byte{0x04, 0x04, 0x01, 0x70, 0x00, 0x00}), TableSection{0x04, 4, map[uint32]*types2.TableType{0: {Limit: &types2.Limit{Min: 0, Max: 0, Type: 0x00}}}, 0, 2}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x04, sectionId, t)
		tableSection, _ := NewTableSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, tableSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, tableSection.Size, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.Tables), len(tableSection.Tables), t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, tableSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, tableSection.StartContent, t)
		for i, table := range testCase.expectedValue.Tables {
			test_utilities2.CompareUInt32(table.Limit.Min, tableSection.Tables[i].Limit.Min, t)
			test_utilities2.CompareUInt32(table.Limit.Max, tableSection.Tables[i].Limit.Max, t)
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{Reader: bytes.NewReader([]byte{}), Err: errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewTableSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

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

type TestCaseImportSection struct {
	reader        io.Reader
	expectedValue ImportSection
}

func TestNewImportSection(t *testing.T) {
	// Positive test cases
	positiveTestCases := []TestCaseImportSection{{bytes.NewReader([]byte{0x02, 0x11, 0x01, 0x06, 0x69, 0x54, 0x65, 0x73, 0x74, 0x4D, 0x06, 0x69, 0x54, 0x65, 0x73, 0x74, 0x4E, 0x00, 0x01}),
		ImportSection{
			0x02,
			17,
			[]ImportValue{},
			map[uint32]*Import{0: &Import{"iTestM", "iTestN", &ImportDesc{0x00, 0x01, &types2.TableType{}, &types2.Limit{}, &types2.GlobalType{}}}},
			map[uint32]*Import{},
			map[uint32]*Import{},
			map[uint32]*Import{},
			0,
			2,
			make(map[uint32][]uint32)}}}

	for _, testCase := range positiveTestCases {
		sectionId, _ := values.ReadNextByte(testCase.reader)
		test_utilities2.CompareByte(0x02, sectionId, t)
		importSection, _ := NewImportSection(testCase.reader)
		test_utilities2.CompareByte(testCase.expectedValue.Id, importSection.Id, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.Size, importSection.Size, t)
		test_utilities2.CompareInt(len(testCase.expectedValue.FuncImports), len(importSection.FuncImports), t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartSection, importSection.StartSection, t)
		test_utilities2.CompareUInt32(testCase.expectedValue.StartContent, importSection.StartContent, t)
		for i, imp := range testCase.expectedValue.FuncImports {
			test_utilities2.CompareString(imp.Name, importSection.FuncImports[i].Name, t)
			test_utilities2.CompareString(imp.ModName, importSection.FuncImports[i].ModName, t)
			test_utilities2.CompareByte(imp.ImportDesc.ImportType, importSection.FuncImports[i].ImportDesc.ImportType, t)
			test_utilities2.CompareUInt32(imp.ImportDesc.TypeIdx, importSection.FuncImports[i].ImportDesc.TypeIdx, t)

			if imp.ImportDesc.TableType == nil {
				test_utilities2.CompareUInt32(imp.ImportDesc.TableType.Limit.Min, importSection.FuncImports[i].ImportDesc.TableType.Limit.Min, t)
				test_utilities2.CompareUInt32(imp.ImportDesc.TableType.Limit.Max, importSection.FuncImports[i].ImportDesc.TableType.Limit.Max, t)
			} else {
				if importSection.FuncImports[i].ImportDesc.TableType != nil {
					t.Error("unexpected nil for TableType")
				}
			}

			if imp.ImportDesc.MemType == nil {
				test_utilities2.CompareUInt32(imp.ImportDesc.MemType.Min, importSection.FuncImports[i].ImportDesc.MemType.Min, t)
				test_utilities2.CompareUInt32(imp.ImportDesc.MemType.Max, importSection.FuncImports[i].ImportDesc.MemType.Max, t)
			} else {
				if importSection.FuncImports[i].ImportDesc.MemType != nil {
					t.Error("unexpected nil for MemType")
				}
			}

			if imp.ImportDesc.GlobalType == nil {
				test_utilities2.CompareString(imp.ImportDesc.GlobalType.ValType, importSection.FuncImports[i].ImportDesc.GlobalType.ValType, t)
				test_utilities2.CompareByte(imp.ImportDesc.GlobalType.Mut, importSection.FuncImports[i].ImportDesc.GlobalType.Mut, t)
			} else {
				if importSection.FuncImports[i].ImportDesc.GlobalType != nil {
					t.Error("unexpected nil for GlobalType")
				}
			}
		}
		values.ResetByteCounter()
	}

	// Negative test cases
	negativeTestCases := []test_utilities2.TestCaseError{
		{bytes.NewReader([]byte{}), errors.New("EOF ==> section size could not be determined")}}

	for _, testCase := range negativeTestCases {
		_, err := NewImportSection(testCase.Reader)
		test_utilities2.CompareErrorMessage(testCase.Err, err, t)
	}
}

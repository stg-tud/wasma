package sections

import (
	"errors"
	"fmt"
	"io"
	"log"
	types2 "wasma/pkg/wasmp/types"
	values2 "wasma/pkg/wasmp/values"
)

type ImportDesc struct {
	ImportType byte
	TypeIdx    uint32
	TableType  *types2.TableType
	MemType    *types2.Limit
	GlobalType *types2.GlobalType
}

// Parses an import section starting with the Size.
// The section Id is parsed before and will be automatically set.
func newImportDesc(reader io.Reader) (*ImportDesc, error) {
	importDesc := new(ImportDesc)
	nextByte, err := values2.ReadNextByte(reader)
	if err != nil {
		return nil, err
	}

	importDesc.ImportType = nextByte

	switch importDesc.ImportType {
	case 0x00:
		typeIdx, err := values2.ReadU32(reader)
		if err != nil {
			return nil, err
		}
		importDesc.TypeIdx = typeIdx
	case 0x01:
		tableType, err := types2.NewTableType(reader)
		if err != nil {
			return nil, err
		}
		importDesc.TableType = tableType
	case 0x02:
		memType, err := types2.NewMemoryType(reader)
		if err != nil {
			return nil, err
		}
		importDesc.MemType = memType
	case 0x03:
		globalType, err := types2.NewGlobalType(reader)
		if err != nil {
			return nil, err
		}
		importDesc.GlobalType = globalType
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Error while reading importdesc. No valid identifier, got: %x", importDesc.ImportType))
	}

	return importDesc, nil
}

type Import struct {
	ModName    string
	Name       string
	ImportDesc *ImportDesc
}

func newImport(reader io.Reader) (*Import, error) {
	imp := new(Import)

	modName, err := values2.NewName(reader)
	if err != nil {
		return nil, err
	}
	imp.ModName = modName // module Name

	functionName, err := values2.NewName(reader)
	if err != nil {
		return nil, err
	}
	imp.Name = functionName // Name

	imp.ImportDesc, err = newImportDesc(reader)
	if err != nil {
		return nil, err
	}

	return imp, nil
}

type ImportValue struct {
	Imp   *Import
	Index uint32
}

type ImportSection struct {
	Id      byte
	Size    uint32
	Imports []ImportValue
	// Deviation from specification for easier processing:
	// Imports are divided into func, table, mem and global imports
	// key: funcIdx
	FuncImports   map[uint32]*Import
	TableImports  map[uint32]*Import
	MemImports    map[uint32]*Import
	GlobalImports map[uint32]*Import
	StartSection  uint32
	StartContent  uint32
	// Deviation from specification for easier processing:
	// Additional field
	// key: typeIdx, value: list of funcIdxs
	FuncIdxs map[uint32][]uint32
}

// Returns a map that assigns all imported functions to their
// corresponding type.
func getImportsFuncIdxs(importSection *ImportSection) map[uint32][]uint32 {
	// key: typeIdx, value: list of funcIdxs
	funcIdxs := make(map[uint32][]uint32)
	for funcIdx, funcImport := range importSection.FuncImports {
		funcIdxs[funcImport.ImportDesc.TypeIdx] = append(funcIdxs[funcImport.ImportDesc.TypeIdx], funcIdx)
	}

	return funcIdxs
}

func NewImportSection(reader io.Reader) (*ImportSection, error) {
	var err error
	importSection := new(ImportSection)
	importSection.Id = 0x02
	importSection.FuncImports = make(map[uint32]*Import)
	importSection.MemImports = make(map[uint32]*Import)
	importSection.TableImports = make(map[uint32]*Import)
	importSection.GlobalImports = make(map[uint32]*Import)
	importSection.StartSection = values2.GetByteCounter()

	importSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	importSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, importSection.Size)
	if err != nil {
		values2.UpdateByteCounter(importSection.StartContent - 1 + importSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(importSection.StartContent - 1 + importSection.Size)
		return nil, err
	}

	log.Printf("Identified import section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		importSection.Id,
		importSection.StartSection,
		importSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		imp, err := newImport(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(importSection.StartContent - 1 + importSection.Size)
			return nil, err
		}

		switch imp.ImportDesc.ImportType {
		case 0x00:
			funcIdx := NewFunctionIndex()
			importSection.FuncImports[funcIdx] = imp
			importSection.Imports = append(importSection.Imports, ImportValue{imp, funcIdx})
		case 0x01:
			tableIndex := NewTableIndex()
			importSection.TableImports[tableIndex] = imp
			importSection.Imports = append(importSection.Imports, ImportValue{imp, tableIndex})
		case 0x02:
			memoryIndex := NewMemoryIndex()
			importSection.MemImports[memoryIndex] = imp
			importSection.Imports = append(importSection.Imports, ImportValue{imp, memoryIndex})
		case 0x03:
			globalIdx := NewGlobalIndex()
			importSection.GlobalImports[globalIdx] = imp
			importSection.Imports = append(importSection.Imports, ImportValue{imp, globalIdx})
		default:
			return nil, errors.New("invalid import type")

		}

	}

	err = contentCompletelyRead("import", importSection.Size, importSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(importSection.StartContent - 1 + importSection.Size)
		return nil, err
	}

	SetInitValueForFunctionIndexAfterImport()
	SetInitValueForMemoryIndexAfterImport()
	SetInitValueForTableIndexAfterImport()
	SetInitValueForGlobalAfterImport()
	importSection.FuncIdxs = getImportsFuncIdxs(importSection)

	return importSection, nil
}

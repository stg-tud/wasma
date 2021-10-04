package sections

import (
	"errors"
	"io"
	"log"
	values2 "wasma/pkg/wasmp/values"
)

type FunctionSection struct {
	Id   byte
	Size uint32
	// key: funcIdx, value: typeIdx
	TypeIdxs map[uint32]uint32
	// key: typeIdx, value: []funcIdxs
	FuncIdxs     map[uint32][]uint32 // additional field not in specification
	StartSection uint32
	StartContent uint32
}

// Parses a Function section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewFunctionSection(reader io.Reader) (*FunctionSection, error) {
	ResetFunctionIndexAfterImport()
	var err error
	functionSection := new(FunctionSection)
	functionSection.Id = 0x03
	// key: funcIdx, value: typeIdx
	functionSection.TypeIdxs = make(map[uint32]uint32)
	// key: typeIdx, value: []funcIdxs
	functionSection.FuncIdxs = make(map[uint32][]uint32)
	functionSection.StartSection = values2.GetByteCounter()

	functionSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	functionSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, functionSection.Size)
	if err != nil {
		values2.UpdateByteCounter(functionSection.StartContent - 1 + functionSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(functionSection.StartContent - 1 + functionSection.Size)
		return nil, err
	}

	log.Printf("Identified function section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		functionSection.Id,
		functionSection.StartSection,
		functionSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		// read typeidx
		typeIdx, err := values2.ReadU32(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(functionSection.StartContent - 1 + functionSection.Size)
			return nil, err
		}
		funcIdx := NewFunctionIndex()
		functionSection.TypeIdxs[funcIdx] = typeIdx

		if funcIdxs, found := functionSection.FuncIdxs[typeIdx]; found {
			functionSection.FuncIdxs[typeIdx] = append(funcIdxs, funcIdx)
		} else {
			functionSection.FuncIdxs[typeIdx] = []uint32{funcIdx}
		}
	}

	err = contentCompletelyRead("Function", functionSection.Size, functionSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(functionSection.StartContent - 1 + functionSection.Size)
		return nil, err
	}

	return functionSection, nil
}

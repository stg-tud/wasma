package sections

import (
	"errors"
	"io"
	"log"
	"wasma/pkg/wasmp/types"
	values2 "wasma/pkg/wasmp/values"
)

type TypeSection struct {
	Id   byte
	Size uint32
	// key: typeIdx
	FunctionTypes map[uint32]types.FunctionType
	StartSection  uint32
	StartContent  uint32
}

// Parses a type section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewTypeSection(reader io.Reader) (*TypeSection, error) {
	var err error
	typeSection := new(TypeSection)
	typeSection.Id = 0x01
	// key: typeIdx, value: function type
	typeSection.FunctionTypes = make(map[uint32]types.FunctionType)
	typeSection.StartSection = values2.GetByteCounter()

	typeSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	typeSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, typeSection.Size)
	if err != nil {
		values2.UpdateByteCounter(typeSection.StartContent - 1 + typeSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(typeSection.StartContent - 1 + typeSection.Size)
		return nil, err
	}

	log.Printf("Identified type section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		typeSection.Id,
		typeSection.StartSection,
		typeSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		typeIdx := i - 1
		functionType, err := types.NewFunctionType(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(typeSection.StartContent - 1 + typeSection.Size)
			return nil, err
		}
		typeSection.FunctionTypes[typeIdx] = functionType
	}

	err = contentCompletelyRead("type", typeSection.Size, typeSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(typeSection.StartContent - 1 + typeSection.Size)
		return nil, err
	}

	return typeSection, nil
}

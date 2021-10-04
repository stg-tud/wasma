package sections

import (
	"errors"
	"io"
	"log"
	types2 "wasma/pkg/wasmp/types"
	values2 "wasma/pkg/wasmp/values"
)

type MemorySection struct {
	Id           byte
	Size         uint32
	MemTypes     map[uint32]*types2.Limit
	StartSection uint32
	StartContent uint32
}

// Parses a memory section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewMemorySection(reader io.Reader) (*MemorySection, error) {
	ResetMemoryIndexAfterImport()
	var err error
	memorySection := new(MemorySection)
	memorySection.Id = 0x05
	memorySection.MemTypes = make(map[uint32]*types2.Limit)
	memorySection.StartSection = values2.GetByteCounter()

	memorySection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	memorySection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, memorySection.Size)
	if err != nil {
		values2.UpdateByteCounter(memorySection.StartContent - 1 + memorySection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(memorySection.StartContent - 1 + memorySection.Size)
		return nil, err
	}

	log.Printf("Identified memory section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		memorySection.Id,
		memorySection.StartSection,
		memorySection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		memType, err := types2.NewMemoryType(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(memorySection.StartContent - 1 + memorySection.Size)
			return nil, err
		}
		memorySection.MemTypes[NewMemoryIndex()] = memType
	}

	err = contentCompletelyRead("memory", memorySection.Size, memorySection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(memorySection.StartContent - 1 + memorySection.Size)
		return nil, err
	}

	return memorySection, nil
}

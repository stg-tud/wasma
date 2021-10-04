package sections

import (
	"errors"
	"io"
	"log"
	values2 "wasma/pkg/wasmp/values"
)

type StartSection struct {
	Id           byte
	Size         uint32
	FuncIdx      uint32
	StartSection uint32
	StartContent uint32
}

// Parses a start section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewStartSection(reader io.Reader) (*StartSection, error) {
	var err error
	startSection := new(StartSection)
	startSection.Id = 0x08
	startSection.StartSection = values2.GetByteCounter()

	startSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	startSection.StartContent = values2.GetPositionOfNextByte()

	startSection.FuncIdx, err = values2.ReadU32(reader)
	if err != nil {
		values2.UpdateByteCounter(startSection.StartContent - 1 + startSection.Size)
		return nil, err
	}

	log.Printf("Identified start section: Id = %x, start = %x, Size = %v, FuncIdx = %v. Parsing...\n",
		startSection.Id,
		startSection.StartSection,
		startSection.Size,
		startSection.FuncIdx)

	err = contentCompletelyRead("start", startSection.Size, startSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(startSection.StartContent - 1 + startSection.Size)
		return nil, err
	}

	return startSection, nil
}

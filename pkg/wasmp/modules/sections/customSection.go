package sections

import (
	"errors"
	"io"
	"log"
	values2 "wasma/pkg/wasmp/values"
)

type CustomSection struct {
	Id           byte
	Size         uint32
	Name         string
	CustomBytes  []byte
	StartSection uint32
	StartContent uint32
}

// Parses a custom section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewCustomSection(reader io.Reader) (*CustomSection, error) {
	var err error
	customSection := new(CustomSection)
	customSection.Id = 0x00
	customSection.StartSection = values2.GetByteCounter()

	customSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	customSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, customSection.Size)
	if err != nil {
		values2.UpdateByteCounter(customSection.StartContent - 1 + customSection.Size)
		return nil, err
	}

	customSection.Name, err = values2.NewName(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(customSection.StartContent - 1 + customSection.Size)
		return nil, err
	}

	log.Printf("Identified custom section: Id = %x, start = %x, Size = %v, Name = %s. Parsing...",
		customSection.Id,
		customSection.StartSection,
		customSection.Size,
		customSection.Name)

	customSection.CustomBytes, err = values2.ReadNextBytes(sectionReader, customSection.Size-(values2.GetByteCounter()-customSection.StartContent+1))
	if err != nil {
		values2.UpdateByteCounter(customSection.StartContent - 1 + customSection.Size)
		return nil, err
	}

	err = contentCompletelyRead("custom", customSection.Size, customSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(customSection.StartContent - 1 + customSection.Size)
		return nil, err
	}

	return customSection, nil
}

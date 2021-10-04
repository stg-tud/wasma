package sections

import (
	"errors"
	"io"
	"log"
	"wasma/pkg/wasmp/instructions"
	"wasma/pkg/wasmp/types"
	values2 "wasma/pkg/wasmp/values"
)

type Global struct {
	GlobalType *types.GlobalType
	Expr       *instructions.Expr
}

type GlobalSection struct {
	Id           byte
	Size         uint32
	Globals      map[uint32]Global
	StartSection uint32
	StartContent uint32
}

// Parses a global section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewGlobalSection(reader io.Reader) (*GlobalSection, error) {
	ResetGlobalIndexAfterImport()
	var err error
	globalSection := new(GlobalSection)
	globalSection.Id = 0x06
	globalSection.Globals = make(map[uint32]Global)
	globalSection.StartSection = values2.GetByteCounter()

	globalSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	globalSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, globalSection.Size)
	if err != nil {
		values2.UpdateByteCounter(globalSection.StartContent - 1 + globalSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(globalSection.StartContent - 1 + globalSection.Size)
		return nil, err
	}

	log.Printf("Identified global section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		globalSection.Id,
		globalSection.StartSection,
		globalSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		globalType, err := types.NewGlobalType(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(globalSection.StartContent - 1 + globalSection.Size)
			return nil, err
		}

		expr, err := instructions.NewExpr(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(globalSection.StartContent - 1 + globalSection.Size)
			return nil, err
		}

		globalSection.Globals[NewGlobalIndex()] = Global{globalType, expr}
	}

	err = contentCompletelyRead("global", globalSection.Size, globalSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(globalSection.StartContent - 1 + globalSection.Size)
		return nil, err
	}

	return globalSection, nil
}

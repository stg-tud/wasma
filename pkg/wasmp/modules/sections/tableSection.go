package sections

import (
	"errors"
	"io"
	"log"
	"wasma/pkg/wasmp/types"
	values2 "wasma/pkg/wasmp/values"
)

type TableSection struct {
	Id           byte
	Size         uint32
	Tables       map[uint32]*types.TableType
	StartSection uint32
	StartContent uint32
}

// Parses a table section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewTableSection(reader io.Reader) (*TableSection, error) {
	ResetTableIndexAfterImport()
	var err error
	tableSection := new(TableSection)
	tableSection.Id = 0x04
	tableSection.Tables = make(map[uint32]*types.TableType)
	tableSection.StartSection = values2.GetByteCounter()

	tableSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	tableSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, tableSection.Size)
	if err != nil {
		values2.UpdateByteCounter(tableSection.StartContent - 1 + tableSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(tableSection.StartContent - 1 + tableSection.Size)
		return nil, err
	}

	log.Printf("Identified table section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		tableSection.Id,
		tableSection.StartSection,
		tableSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		table, err := types.NewTableType(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(tableSection.StartContent - 1 + tableSection.Size)
			return nil, err
		}
		tableSection.Tables[NewTableIndex()] = table
	}

	err = contentCompletelyRead("table", tableSection.Size, tableSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(tableSection.StartContent - 1 + tableSection.Size)
		return nil, err
	}

	return tableSection, nil
}

package sections

import (
	"errors"
	"io"
	"log"
	"wasma/pkg/wasmp/instructions"
	values2 "wasma/pkg/wasmp/values"
)

type Data struct {
	MemIdx uint32
	Expr   *instructions.Expr
	Bytes  []byte
}

func newData(reader io.Reader) (*Data, error) {
	var err error
	data := new(Data)

	data.MemIdx, err = values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	data.Expr, err = instructions.NewExpr(reader)
	if err != nil {
		return nil, err
	}

	vecLen, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	data.Bytes, err = values2.ReadNextBytes(reader, vecLen)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type DataSection struct {
	Id           byte
	Size         uint32
	Datas        []*Data
	StartSection uint32
	StartContent uint32
}

// Parses a data section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewDataSection(reader io.Reader) (*DataSection, error) {
	var err error
	dataSection := new(DataSection)
	dataSection.Id = 0x0B
	dataSection.StartSection = values2.GetByteCounter()

	dataSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	dataSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, dataSection.Size)
	if err != nil {
		values2.UpdateByteCounter(dataSection.StartContent - 1 + dataSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(dataSection.StartContent - 1 + dataSection.Size)
		return nil, err
	}

	log.Printf("Identified data section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		dataSection.Id,
		dataSection.StartSection,
		dataSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		data, err := newData(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(dataSection.StartContent - 1 + dataSection.Size)
			return nil, err
		}
		dataSection.Datas = append(dataSection.Datas, data)
	}

	err = contentCompletelyRead("data", dataSection.Size, dataSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(dataSection.StartContent - 1 + dataSection.Size)
		return nil, err
	}

	return dataSection, nil
}

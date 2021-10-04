package types

import (
	"errors"
	"fmt"
	"io"
	"wasma/pkg/wasmp/values"
)

type TableType struct {
	Limit *Limit
}

func (tableType *TableType) elemType() byte {
	return 0x70
}

func NewTableType(reader io.Reader) (*TableType, error) {
	nextByte, err := values.ReadNextByte(reader)
	if err != nil {
		return nil, err
	}

	tableType := new(TableType)

	if nextByte != 0x70 {
		return tableType, errors.New(fmt.Sprintf("Error while reading table type. Expected 0x70 but got: %x", nextByte))
	}

	limit, err := NewLimit(reader)
	if err != nil {
		return tableType, err
	}

	tableType.Limit = limit

	return tableType, nil
}

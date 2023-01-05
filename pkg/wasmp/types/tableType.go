package types

import (
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

	if nextByte != tableType.elemType() {
		return tableType, fmt.Errorf("error while reading table type. Expected %x but got: %x", tableType.elemType(), nextByte)
	}

	limit, err := NewLimit(reader)
	if err != nil {
		return tableType, err
	}

	tableType.Limit = limit

	return tableType, nil
}

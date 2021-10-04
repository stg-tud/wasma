package types

import (
	"errors"
	"fmt"
	"io"
	"wasma/pkg/wasmp/values"
)

type GlobalType struct {
	ValType string
	Mut     byte
}

func NewGlobalType(reader io.Reader) (*GlobalType, error) {
	globalType := new(GlobalType)

	valType, err := NewValType(reader)
	if err != nil {
		return nil, err
	}
	globalType.ValType = valType

	mut, err := values.ReadNextByte(reader)
	if err != nil {
		return nil, err
	}
	globalType.Mut = mut

	if globalType.Mut != 0x00 && globalType.Mut != 0x01 {
		return globalType, errors.New(fmt.Sprintf("Error while reading global type Mut value. Expected 0x00 or 0x01, but got: %x.", globalType.Mut))
	}

	return globalType, nil
}

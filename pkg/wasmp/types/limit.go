package types

import (
	"fmt"
	"io"
	values2 "wasma/pkg/wasmp/values"
)

type Limit struct {
	Min  uint32
	Max  uint32
	Type byte
}

func NewLimit(reader io.Reader) (*Limit, error) {
	nextByte, err := values2.ReadNextByte(reader)
	if err != nil {
		return nil, err
	}

	limit := new(Limit)
	min, err := values2.ReadU32(reader)

	if err != nil {
		return nil, err
	}

	limit.Type = nextByte

	if nextByte == 0x00 {
		limit.Min = min
	} else if nextByte == 0x01 {
		limit.Min = min

		max, err := values2.ReadU32(reader)
		if err != nil {
			return nil, err
		}
		limit.Max = max
	} else {
		return limit, fmt.Errorf("error while reading Limit. Expected 0x00 or 0x01 but got: %x", nextByte)
	}
	return limit, err
}

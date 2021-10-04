package types

import (
	"errors"
	"fmt"
	"io"
	"wasma/pkg/wasmp/values"
)

func NewValType(reader io.Reader) (string, error) {
	nextByte, err := values.ReadNextByte(reader)
	if err != nil {
		return "", err
	}

	valueType := ""
	var error error = nil

	switch nextByte {
	case 0x7F:
		valueType = "i32"
	case 0x7E:
		valueType = "i64"
	case 0x7D:
		valueType = "f32"
	case 0x7C:
		valueType = "f64"
	default:
		error = errors.New(fmt.Sprintf("Error while reading result type. Unexpected type got: %x. Valid types: 0x7F, 0x7E, 0x7D and 0x7C.", nextByte))
	}
	return valueType, error
}
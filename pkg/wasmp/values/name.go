package values

import (
	"io"
)

// Reads a name value.
func NewName(reader io.Reader) (string, error) {
	vecLen, err := ReadU32(reader)

	if err != nil {
		return "", err
	}

	name, err := ReadNextBytes(reader, uint32(vecLen))
	if err != nil {
		return "", err
	}

	return string(name), nil
}

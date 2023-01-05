package sections

import (
	"bytes"
	"fmt"
	"io"
)

// Reads all bytes of a section an returns a new
// byte reader
func readNextSection(reader io.Reader, size uint32) (*bytes.Reader, error) {
	if size < 1 {
		return nil, fmt.Errorf("sectionn size must be greater or equal 1, but got: %v", size)
	}

	sectionBytes := make([]byte, size)

	_, err := reader.Read(sectionBytes)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(sectionBytes), nil
}

func contentCompletelyRead(section string, size uint32, startByte uint32, currentByte uint32) error {
	//if currentByte-startByte  == size {
	if currentByte-size+1 == startByte {
		return nil
	} else {
		if currentByte-startByte < size {
			return fmt.Errorf("not all Bytes of the %s section were read", section)
		} else {
			return fmt.Errorf("more Bytes were read than the current %s section contains", section)
		}
	}
}

package values

import (
	"bytes"
	"errors"
	"testing"
	"wasma/internal/test_utilities"
)

func TestNewName(t *testing.T) {
	// Positive test case
	// Encoded string consists of: (vector length = 12) + (string value = functionName)
	readerP := bytes.NewReader([]byte{0x0C, 0x66, 0x75, 0x6E, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x4E, 0x61, 0x6D, 0x65})
	name, _ := NewName(readerP)
	test_utilities.CompareString("functionName", name, t)

	// Negative test case
	// Encoded string consists of: (vector length = 12) + (string value = "")
	readerN := bytes.NewReader([]byte{0x0C})
	_, err := NewName(readerN)
	test_utilities.CompareErrorMessage(errors.New("EOF"), err, t)
}

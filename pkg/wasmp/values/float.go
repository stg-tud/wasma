package values

import (
	"encoding/binary"
	"io"
	"math"
)

// Reads a float value that is encoded in the IEEE 754 format
// and returns a float32.
func ReadF32(reader io.Reader) (float32, error) {
	nextBytes, err := ReadNextBytes(reader, 4)

	if err != nil {
		return 0, err
	}

	bytesUint32 := binary.LittleEndian.Uint32(nextBytes)
	floatNumber := math.Float32frombits(bytesUint32)
	return floatNumber, nil

}

// Reads a float value that is encoded in the IEEE 754 format
// and returns a float64.
func ReadF64(reader io.Reader) (float64, error) {
	nextBytes, err := ReadNextBytes(reader, 8)

	if err != nil {
		return 0, err
	}
	bytesUint64 := binary.LittleEndian.Uint64(nextBytes)
	floatNumber := math.Float64frombits(bytesUint64)
	return floatNumber, nil
}

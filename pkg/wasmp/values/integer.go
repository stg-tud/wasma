package values

import (
	"io"
)

// Reads an unsigned integer value that is encoded in the LEB128 format
// and returns an uint64.
func ReadU64(reader io.Reader) (uint64, error) {
	var result uint64
	var shift uint64

	for {
		nextByte, err := ReadNextByte(reader)
		if err != nil {
			return 0, err
		}
		//byteUInt64 := uint64(nextByte)
		//bytel := byteUInt64 & 0x7F //
		//result |= bytel << shift
		result |= uint64(nextByte&0x7F) << shift //
		if nextByte&0x80 == 0 {
			break
		}
		shift += 7
	}
	return result, nil
}

// Reads an unsigned integer value that is encoded in the LEB128 format
// and returns an uint32.
func ReadU32(reader io.Reader) (uint32, error) {
	value, err := ReadU64(reader)
	return uint32(value), err
}

// Reads an signed integer value that is encoded in the LEB128 format
// and returns an int64.
func ReadS33(firstByte byte, reader io.Reader) (int64, error) {
	var result int64 = 0
	var shift int64 = 0
	const int33ValueBits int64 = 4294967295 // bits 1 - 32
	const int33Sign int64 = 4294967296      // bit 33

	first := true
	var nextByte byte
	var err error

	for {
		if first {
			first = false
			nextByte = firstByte
		} else {
			nextByte, err = ReadNextByte(reader)
		}
		if err != nil {
			return 0, err
		}
		result |= int64(nextByte&0x7F) << shift
		shift += 7
		if 0x80&nextByte == 0 {
			value := result & int33ValueBits
			sign := result & int33Sign
			if shift < 36 && sign == int33Sign {
				return -int33Sign + value, nil
			}
			return result & int33ValueBits, err
		}
	}
}

// Reads an signed integer value that is encoded in the LEB128 format
// and returns an int64.
func ReadS64(reader io.Reader) (int64, error) {
	var result int64 = 0
	var shift int64 = 0

	for {
		nextByte, err := ReadNextByte(reader)
		if err != nil {
			return 0, err
		}
		result |= int64(nextByte&0x7F) << shift
		shift += 7
		if 0x80&nextByte == 0 {
			if shift < 64 && (nextByte&0x40) != 0 {
				return result | (^0 << shift), nil
			}
			return result, err
		}
	}
}

// Reads an signed integer value that is encoded in the LEB128 format
// and returns an int32.
func ReadS32(reader io.Reader) (int32, error) {
	value, err := ReadS64(reader)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

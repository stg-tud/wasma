package values

import (
	"io"
)

// Number of read bytes.
var byteCounter uint32
var byteCounterInit = true

// opcode counter
var opcounterOn = false
var elseCounter = 0 // 0x05
var endCounter = 0  // 0x0B

func ResetOpcodeCounter() {
	elseCounter = 0
	endCounter = 0
}

func IncrementElseCounter() {
	if opcounterOn {
		elseCounter++
	}
}

func IncrementEndCounter() {
	if opcounterOn {
		endCounter++
	}
}

func OpCounterOn() {
	opcounterOn = true
}

func OpCounterOff() {
	opcounterOn = false
}

func GetElseCounter() int {
	return elseCounter
}

func GetEndCounter() int {
	return endCounter
}

// Resets the variables byteCounter and byteCounterInit
func ResetByteCounter() {
	byteCounter = 0
	byteCounterInit = true
}

// Updates the value of the byte counter.
func UpdateByteCounter(newValue uint32) {
	byteCounter = newValue
}

// Returns the number of read bytes.
func GetByteCounter() uint32 {
	return byteCounter
}

func GetPositionOfNextByte() uint32 {
	if byteCounter <= 0 {
		return 0
	} else {
		return byteCounter + 1
	}
}

// Returns the position for the current byte
func GetPosition() uint32 {
	if byteCounter > 0 {
		return byteCounter - 1
	} else {
		return 0
	}
}

// Reads the next n bytes of a reader.
func ReadNextBytes(file io.Reader, n uint32) ([]byte, error) {
	if n < 1 {
		return nil, nil
	}

	bytes := make([]byte, n)

	_, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}
	if !byteCounterInit {
		byteCounter += n
	} else {
		byteCounterInit = false
		byteCounter += n - 1
	}

	return bytes, nil
}

// Reads the next byte of a reader.
func ReadNextByte(file io.Reader) (byte, error) {
	nextBytes, err := ReadNextBytes(file, 1)

	if err != nil {
		return 0, err
	} else {
		return nextBytes[0], nil
	}
}

package sections

var memoryIndex uint32 = 0
var initValueMemory uint32 = 0

func ResetMemoryIndex() {
	memoryIndex = 0
	initValueMemory = 0
}

func SetInitValueForMemoryIndexAfterImport() {
	initValueMemory = memoryIndex
}

func ResetMemoryIndexAfterImport() {
	memoryIndex = initValueMemory
}

func NewMemoryIndex() uint32 {
	currentMemoryIndex := memoryIndex
	memoryIndex++
	return currentMemoryIndex
}

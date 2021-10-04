package sections

var functionIndex uint32 = 0
var initValue uint32 = 0

func ResetFunctionIndex() {
	functionIndex = 0
	initValue = 0
}

func SetInitValueForFunctionIndexAfterImport() {
	initValue = functionIndex
}

func ResetFunctionIndexAfterImport() {
	functionIndex = initValue
}

// Returns the current function index and increases
// it by one.
func NewFunctionIndex() uint32 {
	currentFunctionIndex := functionIndex
	functionIndex++
	return currentFunctionIndex
}

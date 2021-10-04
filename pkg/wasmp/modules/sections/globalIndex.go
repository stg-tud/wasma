package sections

var globalIndex uint32 = 0
var initValueGlobal uint32 = 0

func ResetGlobalIndex() {
	globalIndex = 0
	initValueGlobal = 0
}

func SetInitValueForGlobalAfterImport() {
	initValueGlobal = globalIndex
}

func ResetGlobalIndexAfterImport() {
	globalIndex = initValueGlobal
}

func NewGlobalIndex() uint32 {
	currentGlobalIndex := globalIndex
	globalIndex++
	return currentGlobalIndex
}

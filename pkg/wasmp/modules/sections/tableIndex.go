package sections

var tableIndex uint32 = 0
var initValueTable uint32 = 0

func ResetTableIndex() {
	tableIndex = 0
	initValueTable = 0
}

func SetInitValueForTableIndexAfterImport() {
	initValueTable = tableIndex
}

func ResetTableIndexAfterImport() {
	tableIndex = initValueTable
}

func NewTableIndex() uint32 {
	currentTableIndex := tableIndex
	tableIndex++
	return currentTableIndex
}

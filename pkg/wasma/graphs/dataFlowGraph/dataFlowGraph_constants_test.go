package dataFlowGraph

import (
	"testing"
)

// Constants

func TestConstUnaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constUnaryInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;1":    0,
		"1;2":      0,
		"3;1;4":    0,
		"4;5":      0,
		"6;1.5;7":  0,
		"7;8":      0,
		"9;1.5;10": 0,
		"10;11":    0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 0, expectedEdges)
}

func TestConstBinaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constBinaryInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;2":     0,
		"1;1;2":     0,
		"2;3":       0,
		"4;1;6":     0,
		"5;1;6":     0,
		"6;7":       0,
		"8;1.5;10":  0,
		"9;1.5;10":  0,
		"10;11":     0,
		"12;1.5;14": 0,
		"13;1.5;14": 0,
		"14;15":     0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 0, expectedEdges)
}

func TestConstTestInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constTestInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;1": 0,
		"1;2":   0,
		"3;1;4": 0,
		"4;5":   0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 0, expectedEdges)
}

func TestConstComparisonInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constComparisonInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;2":     0,
		"1;1;2":     0,
		"2;3":       0,
		"4;1;6":     0,
		"5;1;6":     0,
		"6;7":       0,
		"8;1.5;10":  0,
		"9;1.5;10":  0,
		"10;11":     0,
		"12;1.5;14": 0,
		"13;1.5;14": 0,
		"14;15":     0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 0, expectedEdges)
}

func TestConstConvertInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constConvertInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;1":    0,
		"1;2":      0,
		"3;1;4":    0,
		"4;5":      0,
		"6;1.5;7":  0,
		"7;8":      0,
		"9;1.5;10": 0,
		"10;11":    0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 0, expectedEdges)
}

func TestConstCallInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constCallInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;1":   0,
		"2;1;3":   0,
		"4;1.5;5": 0,
		"6;1.5;7": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 0, expectedEdges)
}

func TestConstCall_indirectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constCall_indirectInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;2":    0,
		"1;0;2":    0,
		"3;1;5":    0,
		"4;1;5":    0,
		"6;1.5;8":  0,
		"7;2;8":    0,
		"9;1.5;11": 0,
		"10;3;11":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 0, expectedEdges)
}

func TestConstIfInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constIfInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;1": 0,
		"2;0;4": 0,
		"3;1;4": 0,
	}

	testDataFlow(t, dataFlowGraph, 3, 0, 0, expectedEdges)
}

func TestConstBr_ifInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constBr_ifInstructions.wasm")

	expectedEdges := map[string]int{
		"1;1;2": 0,
	}

	testDataFlow(t, dataFlowGraph, 1, 0, 0, expectedEdges)
}

func TestConstBr_tableInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constBr_tableInstructions.wasm")

	expectedEdges := map[string]int{
		"3;1;4": 0,
	}

	testDataFlow(t, dataFlowGraph, 1, 0, 0, expectedEdges)
}

func TestConstDropInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constDropInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;1":   0,
		"2;1;3":   0,
		"4;1.5;5": 0,
		"6;1.5;7": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 0, expectedEdges)
}

func TestConstSelectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constSelectInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;3":          0,
		"1;1;3":          0,
		"2;1;3":          0,
		"3;0 or 1;4":     0,
		"5;0;8":          0,
		"6;1;8":          0,
		"7;1;8":          0,
		"8;0 or 1;9":     0,
		"10;0;13":        0,
		"11;1.5;13":      0,
		"12;1;13":        0,
		"13;0 or 1.5;14": 0,
		"15;0;18":        0,
		"16;1.5;18":      0,
		"17;1;18":        0,
		"18;0 or 1.5;19": 0,
	}

	testDataFlow(t, dataFlowGraph, 16, 0, 0, expectedEdges)
}

func TestConstLocalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constLocalInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;1":     0,
		"1;1;L4":    0,
		"2;1;3":     0,
		"3;1;L5":    0,
		"4;1.5;5":   0,
		"5;1.5;L6":  0,
		"6;1.5;7":   0,
		"7;1.5;L7":  0,
		"8;1;9":     0,
		"9;1;L4":    0,
		"9;1;10":    0,
		"11;1;12":   0,
		"12;1;L5":   0,
		"12;1;13":   0,
		"14;1.5;15": 0,
		"15;1.5;L6": 0,
		"15;1.5;16": 0,
		"17;1.5;18": 0,
		"18;1.5;L7": 0,
		"18;1.5;19": 0,
		"20;1;21":   0,
		"21;1;P0":   0,
		"22;1;23":   0,
		"23;1;P1":   0,
		"24;1.5;25": 0,
		"25;1.5;P2": 0,
		"26;1.5;27": 0,
		"27;1.5;P3": 0,
	}

	testDataFlow(t, dataFlowGraph, 28, 8, 0, expectedEdges)
}

func TestConstGlobalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constGlobalInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1;1":    0,
		"1;1;G0":   0,
		"2;1;3":    0,
		"3;1;G1":   0,
		"4;1.5;5":  0,
		"5;1.5;G2": 0,
		"6;1.5;7":  0,
		"7;1.5;G3": 0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 4, expectedEdges)
}

func TestConstReturnInstructions0(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constReturnInstructions0.wasm")

	expectedEdges := map[string]int{
		"3;1;return": 0,
		"4;1;5":      0,
		"7;1;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 3, 0, 0, expectedEdges)
}

func TestConstReturnInstructions1(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/constants/constReturnInstructions1.wasm")

	expectedEdges := map[string]int{
		"0;1;1":      0,
		"2;1;return": 0,
		"4;1;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 3, 0, 0, expectedEdges)
}

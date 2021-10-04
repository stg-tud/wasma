package dataFlowGraph

import (
	"testing"
)

// Calls

func TestCallUnaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callUnaryInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1":   0,
		"1;2":   0,
		"3;4":   0,
		"4;5":   0,
		"6;7":   0,
		"7;8":   0,
		"9;10":  0,
		"10;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 0, expectedEdges)
}

func TestCallBinaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callBinaryInstructions.wasm")

	expectedEdges := map[string]int{
		"0;2":   0,
		"1;2":   0,
		"2;3":   0,
		"4;6":   0,
		"5;6":   0,
		"6;7":   0,
		"8;10":  0,
		"9;10":  0,
		"10;11": 0,
		"12;14": 0,
		"13;14": 0,
		"14;15": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 0, expectedEdges)
}

func TestCallTestInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callTestInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1": 0,
		"1;2": 0,
		"3;4": 0,
		"4;5": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 0, expectedEdges)
}

func TestCallComparisonInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callComparisonInstructions.wasm")

	expectedEdges := map[string]int{
		"0;2":     0,
		"1;0;2":   0,
		"2;3":     0,
		"4;6":     0,
		"5;0;6":   0,
		"6;7":     0,
		"8;10":    0,
		"9;0;10":  0,
		"10;11":   0,
		"12;14":   0,
		"13;0;14": 0,
		"14;15":   0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 0, expectedEdges)
}

func TestCallConvertInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callConvertInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1":   0,
		"1;2":   0,
		"3;4":   0,
		"4;5":   0,
		"6;7":   0,
		"7;8":   0,
		"9;10":  0,
		"10;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 0, expectedEdges)
}

func TestCallCallInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callCallInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1": 0,
		"2;3": 0,
		"4;5": 0,
		"6;7": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 0, expectedEdges)
}

func TestCallCall_indirectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callCall_indirectInstructions.wasm")

	expectedEdges := map[string]int{
		"0;2":     0,
		"1;0;2":   0,
		"3;5":     0,
		"4;1;5":   0,
		"6;8":     0,
		"7;2;8":   0,
		"9;11":    0,
		"10;3;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 0, expectedEdges)
}

func TestCallIfInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callIfInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1":   0,
		"2;0;4": 0,
		"3;1;4": 0,
	}

	testDataFlow(t, dataFlowGraph, 3, 0, 0, expectedEdges)
}

func TestCallBr_ifInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callBr_ifInstructions.wasm")

	expectedEdges := map[string]int{
		"1;2": 0,
	}

	testDataFlow(t, dataFlowGraph, 1, 0, 0, expectedEdges)
}

func TestCallBr_tableInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callBr_tableInstructions.wasm")

	expectedEdges := map[string]int{
		"3;4": 0,
	}

	testDataFlow(t, dataFlowGraph, 1, 0, 0, expectedEdges)
}

func TestCallDropInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callDropInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1": 0,
		"2;3": 0,
		"4;5": 0,
		"6;7": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 0, expectedEdges)
}

func TestCallSelectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callSelectInstructions.wasm")

	expectedEdges := map[string]int{
		"0;3":   0,
		"1;3":   0,
		"2;3":   0,
		"3;4":   0,
		"5;8":   0,
		"6;8":   0,
		"7;8":   0,
		"8;9":   0,
		"10;13": 0,
		"11;13": 0,
		"12;13": 0,
		"13;14": 0,
		"15;18": 0,
		"16;18": 0,
		"17;18": 0,
		"18;19": 0,
	}

	testDataFlow(t, dataFlowGraph, 16, 0, 0, expectedEdges)
}

func TestCallLocalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callLocalInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1":   0,
		"1;L4":  0,
		"2;3":   0,
		"3;L5":  0,
		"4;5":   0,
		"5;L6":  0,
		"6;7":   0,
		"7;L7":  0,
		"8;9":   0,
		"9;L4":  0,
		"9;10":  0,
		"11;12": 0,
		"12;L5": 0,
		"12;13": 0,
		"14;15": 0,
		"15;L6": 0,
		"15;16": 0,
		"17;18": 0,
		"18;L7": 0,
		"18;19": 0,
		"20;21": 0,
		"21;P0": 0,
		"22;23": 0,
		"23;P1": 0,
		"24;25": 0,
		"25;P2": 0,
		"26;27": 0,
		"27;P3": 0,
	}

	testDataFlow(t, dataFlowGraph, 28, 8, 0, expectedEdges)
}

func TestCallGlobalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callGlobalInstructions.wasm")

	expectedEdges := map[string]int{
		"0;1":  0,
		"1;G0": 0,
		"2;3":  0,
		"3;G1": 0,
		"4;5":  0,
		"5;G2": 0,
		"6;7":  0,
		"7;G3": 0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 4, expectedEdges)
}

func TestCallReturnInstructions0(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callReturnInstructions0.wasm")

	expectedEdges := map[string]int{
		"3;return": 0,
		"4;5":      0,
		"7;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 3, 0, 0, expectedEdges)
}

func TestCallReturnInstructions1(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call/callReturnInstructions1.wasm")

	expectedEdges := map[string]int{
		"0;1":      0,
		"2;return": 0,
		"4;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 3, 0, 0, expectedEdges)
}

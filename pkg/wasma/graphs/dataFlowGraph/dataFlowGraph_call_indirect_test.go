package dataFlowGraph

import (
	"testing"
)

// Indirect calls

func TestCallIndirectUnaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectUnaryInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":   0,
		"1;2":     0,
		"2;3":     0,
		"4;1;5":   0,
		"5;6":     0,
		"6;7":     0,
		"8;2;9":   0,
		"9;10":    0,
		"10;11":   0,
		"12;3;13": 0,
		"13;14":   0,
		"14;15":   0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 0, expectedEdges)
}

func TestCallIndirectBinaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectBinaryInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":   0,
		"2;0;3":   0,
		"1;4":     0,
		"3;4":     0,
		"4;5":     0,
		"6;1;7":   0,
		"8;1;9":   0,
		"7;10":    0,
		"9;10":    0,
		"10;11":   0,
		"12;2;13": 0,
		"14;2;15": 0,
		"13;16":   0,
		"15;16":   0,
		"16;17":   0,
		"18;3;19": 0,
		"20;3;21": 0,
		"19;22":   0,
		"21;22":   0,
		"22;23":   0,
	}

	testDataFlow(t, dataFlowGraph, 20, 0, 0, expectedEdges)
}

func TestCallIndirectTestInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectTestInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1": 0,
		"1;2":   0,
		"2;3":   0,
		"4;1;5": 0,
		"5;6":   0,
		"6;7":   0,
	}

	testDataFlow(t, dataFlowGraph, 6, 0, 0, expectedEdges)
}

func TestCallIndirectComparisonInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectComparisonInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":   0,
		"1;3":     0,
		"2;0;3":   0,
		"3;4":     0,
		"5;1;6":   0,
		"6;8":     0,
		"7;0;8":   0,
		"8;9":     0,
		"10;2;11": 0,
		"11;13":   0,
		"12;0;13": 0,
		"13;14":   0,
		"15;3;16": 0,
		"16;18":   0,
		"17;0;18": 0,
		"18;19":   0,
	}

	testDataFlow(t, dataFlowGraph, 16, 0, 0, expectedEdges)
}

func TestCallIndirectConvertInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectConvertInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":   0,
		"1;2":     0,
		"2;3":     0,
		"4;1;5":   0,
		"5;6":     0,
		"6;7":     0,
		"8;2;9":   0,
		"9;10":    0,
		"10;11":   0,
		"12;3;13": 0,
		"13;14":   0,
		"14;15":   0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 0, expectedEdges)
}

func TestCallIndirectCallInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectCallInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":  0,
		"1;2":    0,
		"3;1;4":  0,
		"4;5":    0,
		"6;2;7":  0,
		"7;8":    0,
		"9;3;10": 0,
		"10;11":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 0, expectedEdges)
}

func TestCallIndirectCall_indirectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectCall_indirectInstructions.wasm")

	expectedEdges := map[string]int{
		"0;4;1":   0,
		"1;3":     0,
		"2;0;3":   0,
		"4;5;5":   0,
		"5;7":     0,
		"6;1;7":   0,
		"8;6;9":   0,
		"9;11":    0,
		"10;2;11": 0,
		"12;7;13": 0,
		"13;15":   0,
		"14;3;15": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 0, expectedEdges)
}

func TestCallIndirectIfInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectIfInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1": 0,
		"1;2":   0,
		"3;0;5": 0,
		"4;1;5": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 0, expectedEdges)
}

func TestCallIndirectBr_ifInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectBr_ifInstructions.wasm")

	expectedEdges := map[string]int{
		"1;0;2": 0,
		"2;3":   0,
	}

	testDataFlow(t, dataFlowGraph, 2, 0, 0, expectedEdges)
}

func TestCallIndirectBr_tableInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectBr_tableInstructions.wasm")

	expectedEdges := map[string]int{
		"3;0;4": 0,
		"4;5":   0,
	}

	testDataFlow(t, dataFlowGraph, 2, 0, 0, expectedEdges)
}

func TestCallIndirectDropInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectDropInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":  0,
		"1;2":    0,
		"3;1;4":  0,
		"4;5":    0,
		"6;2;7":  0,
		"7;8":    0,
		"9;3;10": 0,
		"10;11":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 0, expectedEdges)
}

func TestCallIndirectSelectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectSelectInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":    0,
		"2;1;3":    0,
		"4;2;5":    0,
		"1;6":      0,
		"3;6":      0,
		"5;6":      0,
		"6;7":      0,
		"8;3;9":    0,
		"10;4;11":  0,
		"12;5;13":  0,
		"9;14":     0,
		"11;14":    0,
		"13;14":    0,
		"14;15":    0,
		"16;6;17":  0,
		"18;7;19":  0,
		"20;8;21":  0,
		"17;22":    0,
		"19;22":    0,
		"21;22":    0,
		"22;23":    0,
		"24;9;25":  0,
		"26;10;27": 0,
		"28;11;29": 0,
		"25;30":    0,
		"27;30":    0,
		"29;30":    0,
		"30;31":    0,
	}

	testDataFlow(t, dataFlowGraph, 28, 0, 0, expectedEdges)
}

func TestCallIndirectLocalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectLocalInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":   0,
		"1;2":     0,
		"2;L4":    0,
		"3;1;4":   0,
		"4;5":     0,
		"5;L5":    0,
		"6;2;7":   0,
		"7;8":     0,
		"8;L6":    0,
		"9;3;10":  0,
		"10;11":   0,
		"11;L7":   0,
		"12;0;13": 0,
		"13;14":   0,
		"14;L4":   0,
		"14;15":   0,
		"16;1;17": 0,
		"17;18":   0,
		"18;L5":   0,
		"18;19":   0,
		"20;2;21": 0,
		"21;22":   0,
		"22;L6":   0,
		"22;23":   0,
		"24;3;25": 0,
		"25;26":   0,
		"26;L7":   0,
		"26;27":   0,
		"28;0;29": 0,
		"29;30":   0,
		"30;P0":   0,
		"31;1;32": 0,
		"32;33":   0,
		"33;P1":   0,
		"34;2;35": 0,
		"35;36":   0,
		"36;P2":   0,
		"37;3;38": 0,
		"38;39":   0,
		"39;P3":   0,
	}

	testDataFlow(t, dataFlowGraph, 40, 8, 0, expectedEdges)
}

func TestCallIndirectGlobalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectGlobalInstructions.wasm")

	expectedEdges := map[string]int{
		"0;0;1":  0,
		"1;2":    0,
		"2;G0":   0,
		"3;1;4":  0,
		"4;5":    0,
		"5;G1":   0,
		"6;2;7":  0,
		"7;8":    0,
		"8;G2":   0,
		"9;3;10": 0,
		"10;11":  0,
		"11;G3":  0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 4, expectedEdges)
}

func TestCallIndirectReturnInstructions0(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectReturnInstructions0.wasm")

	expectedEdges := map[string]int{
		"3;0;4":     0,
		"5;0;6":     0,
		"4;return":  0,
		"6;7":       0,
		"9;0;10":    0,
		"10;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 6, 0, 0, expectedEdges)
}

func TestCallIndirectReturnInstructions1(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/call_indirect/call_indirectReturnInstructions1.wasm")

	expectedEdges := map[string]int{
		"0;0;1":    0,
		"1;2":      0,
		"3;1;4":    0,
		"4;return": 0,
		"6;2;7":    0,
		"7;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 6, 0, 0, expectedEdges)
}

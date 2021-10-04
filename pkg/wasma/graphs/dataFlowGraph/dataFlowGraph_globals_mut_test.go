package dataFlowGraph

import "testing"

// Globals mut

func TestGlobalMutUnaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutUnaryInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0":  0,
		"0;1":   0,
		"1;2":   0,
		"G1;3":  0,
		"3;4":   0,
		"4;5":   0,
		"G2;6":  0,
		"6;7":   0,
		"7;8":   0,
		"G3;9":  0,
		"9;10":  0,
		"10;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 4, expectedEdges)
}

func TestGlobalMutBinaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutBinaryInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0":  0,
		"G0;1":  0,
		"0;2":   0,
		"1;2":   0,
		"2;3":   0,
		"G1;4":  0,
		"G1;5":  0,
		"4;6":   0,
		"5;6":   0,
		"6;7":   0,
		"G2;8":  0,
		"G2;9":  0,
		"8;10":  0,
		"9;10":  0,
		"10;11": 0,
		"G3;12": 0,
		"G3;13": 0,
		"12;14": 0,
		"13;14": 0,
		"14;15": 0,
	}

	testDataFlow(t, dataFlowGraph, 16, 0, 4, expectedEdges)
}

func TestGlobalMutTestInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutTestInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0": 0,
		"0;1":  0,
		"1;2":  0,
		"G1;3": 0,
		"3;4":  0,
		"4;5":  0,
	}

	testDataFlow(t, dataFlowGraph, 6, 0, 2, expectedEdges)
}

func TestGlobalMutComparisonInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutComparisonInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0":    0,
		"0;2":     0,
		"1;0;2":   0,
		"2;3":     0,
		"G1;4":    0,
		"4;6":     0,
		"5;0;6":   0,
		"6;7":     0,
		"G2;8":    0,
		"8;10":    0,
		"9;0;10":  0,
		"10;11":   0,
		"G3;12":   0,
		"12;14":   0,
		"13;0;14": 0,
		"14;15":   0,
	}

	testDataFlow(t, dataFlowGraph, 16, 0, 4, expectedEdges)
}

func TestGlobalMutConvertInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutConvertInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0":  0,
		"0;1":   0,
		"1;2":   0,
		"G1;3":  0,
		"3;4":   0,
		"4;5":   0,
		"G2;6":  0,
		"6;7":   0,
		"7;8":   0,
		"G3;9":  0,
		"9;10":  0,
		"10;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 4, expectedEdges)
}

func TestGlobalMutCallInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutCallInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0": 0,
		"0;1":  0,
		"G1;2": 0,
		"2;3":  0,
		"G2;4": 0,
		"4;5":  0,
		"G3;6": 0,
		"6;7":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 4, expectedEdges)
}

func TestGlobalMutCall_indirectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutCall_indirectInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0":    0,
		"0;2":     0,
		"1;0;2":   0,
		"G1;3":    0,
		"3;5":     0,
		"4;1;5":   0,
		"G2;6":    0,
		"6;8":     0,
		"7;2;8":   0,
		"G3;9":    0,
		"9;11":    0,
		"10;3;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 4, expectedEdges)
}

func TestGlobalMutIfInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutIfInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0":  0,
		"0;1":   0,
		"3;1;4": 0,
		"2;0;4": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 1, expectedEdges)
}

func TestGlobalMutBr_ifInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutBr_ifInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;1": 0,
		"1;2":  0,
	}

	testDataFlow(t, dataFlowGraph, 2, 0, 1, expectedEdges)
}

func TestGlobalMutBr_tableInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutBr_tableInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;3": 0,
		"3;4":  0,
	}

	testDataFlow(t, dataFlowGraph, 2, 0, 1, expectedEdges)
}

func TestGlobalMutDropInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutDropInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0": 0,
		"0;1":  0,
		"G1;2": 0,
		"2;3":  0,
		"G2;4": 0,
		"4;5":  0,
		"G3;6": 0,
		"6;7":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 0, 4, expectedEdges)
}

func TestGlobalMutSelectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutSelectInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0":                          0,
		"G1;1":                          0,
		"G2;2":                          0,
		"0;3":                           0,
		"1;3":                           0,
		"2;3":                           0,
		"3;value(G0) or value(G1);4":    0,
		"G3;5":                          0,
		"G4;6":                          0,
		"G5;7":                          0,
		"5;8":                           0,
		"6;8":                           0,
		"7;8":                           0,
		"8;value(G3) or value(G4);9":    0,
		"G6;10":                         0,
		"G7;11":                         0,
		"G8;12":                         0,
		"10;13":                         0,
		"11;13":                         0,
		"12;13":                         0,
		"13;value(G6) or value(G7);14":  0,
		"G9;15":                         0,
		"G10;16":                        0,
		"G11;17":                        0,
		"15;18":                         0,
		"16;18":                         0,
		"17;18":                         0,
		"18;value(G9) or value(G10);19": 0,
	}

	testDataFlow(t, dataFlowGraph, 28, 0, 12, expectedEdges)
}

func TestGlobalMutLocalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutLocalInstructions.wasm")

	expectedEdges := map[string]int{
		"G0;0":            0,
		"G0;8":            0,
		"G0;20":           0,
		"0;1":             0,
		"1;value(G0);L4":  0,
		"8;9":             0,
		"9;value(G0);L4":  0,
		"9;value(G0);10":  0,
		"20;21":           0,
		"21;value(G0);P0": 0,
		"G1;2":            0,
		"G1;11":           0,
		"G1;22":           0,
		"2;3":             0,
		"3;value(G1);L5":  0,
		"11;12":           0,
		"12;value(G1);13": 0,
		"12;value(G1);L5": 0,
		"22;23":           0,
		"23;value(G1);P1": 0,
		"G2;4":            0,
		"G2;14":           0,
		"G2;24":           0,
		"4;5":             0,
		"5;value(G2);L6":  0,
		"14;15":           0,
		"15;value(G2);16": 0,
		"15;value(G2);L6": 0,
		"24;25":           0,
		"25;value(G2);P2": 0,
		"G3;6":            0,
		"G3;17":           0,
		"G3;26":           0,
		"6;7":             0,
		"7;value(G3);L7":  0,
		"17;18":           0,
		"18;value(G3);19": 0,
		"18;value(G3);L7": 0,
		"26;27":           0,
		"27;value(G3);P3": 0,
	}

	testDataFlow(t, dataFlowGraph, 32, 8, 4, expectedEdges)
}

func TestGlobalMutGlobalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutGlobalInstructions.wasm")

	expectedEdges := map[string]int{
		"G4;0":           0,
		"0;1":            0,
		"1;value(G4);G0": 0,
		"G5;2":           0,
		"2;3":            0,
		"3;value(G5);G1": 0,
		"G6;4":           0,
		"4;5":            0,
		"5;value(G6);G2": 0,
		"G7;6":           0,
		"6;7":            0,
		"7;value(G7);G3": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 0, 8, expectedEdges)
}

func TestGlobalMutReturnInstructions0(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutReturnInstructions0.wasm")

	expectedEdges := map[string]int{
		"G0;3":     0,
		"G0;4":     0,
		"G0;7":     0,
		"3;return": 0,
		"7;return": 0,
		"4;5":      0,
	}

	testDataFlow(t, dataFlowGraph, 4, 0, 1, expectedEdges)
}

func TestGlobalMutReturnInstructions1(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/globals_mut/global_mutReturnInstructions1.wasm")

	expectedEdges := map[string]int{
		"G0;0":     0,
		"0;1":      0,
		"G1;2":     0,
		"2;return": 0,
		"G2;4":     0,
		"4;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 6, 0, 3, expectedEdges)
}

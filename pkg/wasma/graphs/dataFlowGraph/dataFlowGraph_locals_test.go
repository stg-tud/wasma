package dataFlowGraph

import "testing"

// Locals

func TestLocalUnaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localUnaryInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0":  0,
		"0;1":   0,
		"1;2":   0,
		"L1;3":  0,
		"3;4":   0,
		"4;5":   0,
		"L2;6":  0,
		"6;7":   0,
		"7;8":   0,
		"L3;9":  0,
		"9;10":  0,
		"10;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 4, 0, expectedEdges)
}

func TestLocalBinaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localBinaryInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0":  0,
		"L0;1":  0,
		"0;2":   0,
		"1;2":   0,
		"2;3":   0,
		"L1;4":  0,
		"L1;5":  0,
		"4;6":   0,
		"5;6":   0,
		"6;7":   0,
		"L2;8":  0,
		"L2;9":  0,
		"8;10":  0,
		"9;10":  0,
		"10;11": 0,
		"L3;12": 0,
		"L3;13": 0,
		"12;14": 0,
		"13;14": 0,
		"14;15": 0,
	}

	testDataFlow(t, dataFlowGraph, 16, 4, 0, expectedEdges)
}

func TestLocalTestInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localTestInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0": 0,
		"0;1":  0,
		"1;2":  0,
		"L1;3": 0,
		"3;4":  0,
		"4;5":  0,
	}

	testDataFlow(t, dataFlowGraph, 6, 2, 0, expectedEdges)
}

func TestLocalComparisonInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localComparisonInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0":    0,
		"0;2":     0,
		"1;0;2":   0,
		"2;3":     0,
		"L1;4":    0,
		"4;6":     0,
		"5;0;6":   0,
		"6;7":     0,
		"L2;8":    0,
		"8;10":    0,
		"9;0;10":  0,
		"10;11":   0,
		"L3;12":   0,
		"12;14":   0,
		"13;0;14": 0,
		"14;15":   0,
	}

	testDataFlow(t, dataFlowGraph, 16, 4, 0, expectedEdges)
}

func TestLocalConvertInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localConvertInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0":  0,
		"0;1":   0,
		"1;2":   0,
		"L1;3":  0,
		"3;4":   0,
		"4;5":   0,
		"L2;6":  0,
		"6;7":   0,
		"7;8":   0,
		"L3;9":  0,
		"9;10":  0,
		"10;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 4, 0, expectedEdges)
}

func TestLocalCallInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localCallInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0": 0,
		"0;1":  0,
		"L1;2": 0,
		"2;3":  0,
		"L2;4": 0,
		"4;5":  0,
		"L3;6": 0,
		"6;7":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 4, 0, expectedEdges)
}

func TestLocalCall_indirectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localCall_indirectInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0":    0,
		"0;2":     0,
		"1;0;2":   0,
		"L1;3":    0,
		"3;5":     0,
		"4;1;5":   0,
		"L2;6":    0,
		"6;8":     0,
		"7;2;8":   0,
		"L3;9":    0,
		"9;11":    0,
		"10;3;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 4, 0, expectedEdges)
}

func TestLocalIfInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localIfInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0":  0,
		"0;1":   0,
		"3;1;4": 0,
		"2;0;4": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 1, 0, expectedEdges)
}

func TestLocalBr_ifInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localBr_ifInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;1": 0,
		"1;2":  0,
	}

	testDataFlow(t, dataFlowGraph, 2, 1, 0, expectedEdges)
}

func TestLocalBr_tableInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localBr_tableInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;3": 0,
		"3;4":  0,
	}

	testDataFlow(t, dataFlowGraph, 2, 1, 0, expectedEdges)
}

func TestLocalDropInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localDropInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0": 0,
		"0;1":  0,
		"L1;2": 0,
		"2;3":  0,
		"L2;4": 0,
		"4;5":  0,
		"L3;6": 0,
		"6;7":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 4, 0, expectedEdges)
}

func TestLocalSelectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localSelectInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;2":                           0,
		"L4;0":                           0,
		"L5;1":                           0,
		"0;3":                            0,
		"1;3":                            0,
		"2;3":                            0,
		"3;value(L4) or value(L5);4":     0,
		"L1;7":                           0,
		"L6;5":                           0,
		"L7;6":                           0,
		"5;8":                            0,
		"6;8":                            0,
		"7;8":                            0,
		"8;value(L6) or value(L7);9":     0,
		"L2;12":                          0,
		"L8;10":                          0,
		"L9;11":                          0,
		"10;13":                          0,
		"11;13":                          0,
		"12;13":                          0,
		"13;value(L8) or value(L9);14":   0,
		"L3;17":                          0,
		"L10;15":                         0,
		"L11;16":                         0,
		"15;18":                          0,
		"16;18":                          0,
		"17;18":                          0,
		"18;value(L10) or value(L11);19": 0,
	}

	testDataFlow(t, dataFlowGraph, 28, 12, 0, expectedEdges)
}

func TestLocalLocalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localLocalInstructions.wasm")

	expectedEdges := map[string]int{
		"L8;0":             0,
		"L8;8":             0,
		"L8;20":            0,
		"0;1":              0,
		"1;value(L8);L4":   0,
		"8;9":              0,
		"9;value(L8);L4":   0,
		"9;value(L8);10":   0,
		"20;21":            0,
		"21;value(L8);P0":  0,
		"L9;2":             0,
		"L9;11":            0,
		"L9;22":            0,
		"2;3":              0,
		"3;value(L9);L5":   0,
		"11;12":            0,
		"12;value(L9);13":  0,
		"12;value(L9);L5":  0,
		"22;23":            0,
		"23;value(L9);P1":  0,
		"L10;4":            0,
		"L10;14":           0,
		"L10;24":           0,
		"4;5":              0,
		"5;value(L10);L6":  0,
		"14;15":            0,
		"15;value(L10);16": 0,
		"15;value(L10);L6": 0,
		"24;25":            0,
		"25;value(L10);P2": 0,
		"L11;6":            0,
		"L11;17":           0,
		"L11;26":           0,
		"6;7":              0,
		"7;value(L11);L7":  0,
		"17;18":            0,
		"18;value(L11);19": 0,
		"18;value(L11);L7": 0,
		"26;27":            0,
		"27;value(L11);P3": 0,
	}

	testDataFlow(t, dataFlowGraph, 32, 12, 0, expectedEdges)
}

func TestLocalGlobalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localGlobalInstructions.wasm")

	expectedEdges := map[string]int{
		"L0;0":           0,
		"0;1":            0,
		"1;value(L0);G0": 0,
		"L1;2":           0,
		"2;3":            0,
		"3;value(L1);G1": 0,
		"L2;4":           0,
		"4;5":            0,
		"5;value(L2);G2": 0,
		"L3;6":           0,
		"6;7":            0,
		"7;value(L3);G3": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 4, 4, expectedEdges)
}

func TestLocalReturnInstructions0(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localReturnInstructions0.wasm")

	expectedEdges := map[string]int{
		"L0;3":     0,
		"L0;4":     0,
		"L0;7":     0,
		"3;return": 0,
		"7;return": 0,
		"4;5":      0,
	}

	testDataFlow(t, dataFlowGraph, 4, 1, 0, expectedEdges)
}

func TestLocalReturnInstructions1(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/locals/localReturnInstructions1.wasm")

	expectedEdges := map[string]int{
		"L0;0":     0,
		"0;1":      0,
		"L1;2":     0,
		"2;return": 0,
		"L2;4":     0,
		"4;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 6, 3, 0, expectedEdges)
}

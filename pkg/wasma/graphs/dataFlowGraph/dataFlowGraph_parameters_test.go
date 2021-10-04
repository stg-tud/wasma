package dataFlowGraph

import (
	"testing"
)

// Parameters

func TestParamUnaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramUnaryInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0":  0,
		"0;1":   0,
		"1;2":   0,
		"P1;3":  0,
		"3;4":   0,
		"4;5":   0,
		"P2;6":  0,
		"6;7":   0,
		"7;8":   0,
		"P3;9":  0,
		"9;10":  0,
		"10;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 4, 0, expectedEdges)
}

func TestParamBinaryInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramBinaryInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0":  0,
		"P0;1":  0,
		"0;2":   0,
		"1;2":   0,
		"2;3":   0,
		"P1;4":  0,
		"P1;5":  0,
		"4;6":   0,
		"5;6":   0,
		"6;7":   0,
		"P2;8":  0,
		"P2;9":  0,
		"8;10":  0,
		"9;10":  0,
		"10;11": 0,
		"P3;12": 0,
		"P3;13": 0,
		"12;14": 0,
		"13;14": 0,
		"14;15": 0,
	}

	testDataFlow(t, dataFlowGraph, 16, 4, 0, expectedEdges)
}

func TestParamTestInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramTestInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0": 0,
		"0;1":  0,
		"1;2":  0,
		"P1;3": 0,
		"3;4":  0,
		"4;5":  0,
	}

	testDataFlow(t, dataFlowGraph, 6, 2, 0, expectedEdges)
}

func TestParamComparisonInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramComparisonInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0":    0,
		"0;2":     0,
		"1;0;2":   0,
		"2;3":     0,
		"P1;4":    0,
		"4;6":     0,
		"5;0;6":   0,
		"6;7":     0,
		"P2;8":    0,
		"8;10":    0,
		"9;0;10":  0,
		"10;11":   0,
		"P3;12":   0,
		"12;14":   0,
		"13;0;14": 0,
		"14;15":   0,
	}

	testDataFlow(t, dataFlowGraph, 16, 4, 0, expectedEdges)
}

func TestParamConvertInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramConvertInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0":  0,
		"0;1":   0,
		"1;2":   0,
		"P1;3":  0,
		"3;4":   0,
		"4;5":   0,
		"P2;6":  0,
		"6;7":   0,
		"7;8":   0,
		"P3;9":  0,
		"9;10":  0,
		"10;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 4, 0, expectedEdges)
}

func TestParamCallInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramCallInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0": 0,
		"0;1":  0,
		"P1;2": 0,
		"2;3":  0,
		"P2;4": 0,
		"4;5":  0,
		"P3;6": 0,
		"6;7":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 4, 0, expectedEdges)
}

func TestParamCall_indirectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramCall_indirectInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0":    0,
		"0;2":     0,
		"1;0;2":   0,
		"P1;3":    0,
		"3;5":     0,
		"4;1;5":   0,
		"P2;6":    0,
		"6;8":     0,
		"7;2;8":   0,
		"P3;9":    0,
		"9;11":    0,
		"10;3;11": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 4, 0, expectedEdges)
}

func TestParamIfInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramIfInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0":  0,
		"0;1":   0,
		"3;1;4": 0,
		"2;0;4": 0,
	}

	testDataFlow(t, dataFlowGraph, 4, 1, 0, expectedEdges)
}

func TestParamBr_ifInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramBr_ifInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;1": 0,
		"1;2":  0,
	}

	testDataFlow(t, dataFlowGraph, 2, 1, 0, expectedEdges)
}

func TestParamBr_tableInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramBr_tableInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;3": 0,
		"3;4":  0,
	}

	testDataFlow(t, dataFlowGraph, 2, 1, 0, expectedEdges)
}

func TestParamDropInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramDropInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0": 0,
		"0;1":  0,
		"P1;2": 0,
		"2;3":  0,
		"P2;4": 0,
		"4;5":  0,
		"P3;6": 0,
		"6;7":  0,
	}

	testDataFlow(t, dataFlowGraph, 8, 4, 0, expectedEdges)
}

func TestParamSelectInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramSelectInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;2":                           0,
		"P4;0":                           0,
		"P5;1":                           0,
		"0;3":                            0,
		"1;3":                            0,
		"2;3":                            0,
		"3;value(P4) or value(P5);4":     0,
		"P1;7":                           0,
		"P6;5":                           0,
		"P7;6":                           0,
		"5;8":                            0,
		"6;8":                            0,
		"7;8":                            0,
		"8;value(P6) or value(P7);9":     0,
		"P2;12":                          0,
		"P8;10":                          0,
		"P9;11":                          0,
		"10;13":                          0,
		"11;13":                          0,
		"12;13":                          0,
		"13;value(P8) or value(P9);14":   0,
		"P3;17":                          0,
		"P10;15":                         0,
		"P11;16":                         0,
		"15;18":                          0,
		"16;18":                          0,
		"17;18":                          0,
		"18;value(P10) or value(P11);19": 0,
	}

	testDataFlow(t, dataFlowGraph, 28, 12, 0, expectedEdges)
}

func TestParamLocalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramLocalInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0":             0,
		"P0;8":             0,
		"P0;20":            0,
		"0;1":              0,
		"1;value(P0);L8":   0,
		"8;9":              0,
		"9;value(P0);L8":   0,
		"9;value(P0);10":   0,
		"20;21":            0,
		"21;value(P0);P4":  0,
		"P1;2":             0,
		"P1;11":            0,
		"P1;22":            0,
		"2;3":              0,
		"3;value(P1);L9":   0,
		"11;12":            0,
		"12;value(P1);L9":  0,
		"12;value(P1);13":  0,
		"22;23":            0,
		"23;value(P1);P5":  0,
		"P2;4":             0,
		"P2;14":            0,
		"P2;24":            0,
		"4;5":              0,
		"5;value(P2);L10":  0,
		"14;15":            0,
		"15;value(P2);L10": 0,
		"15;value(P2);16":  0,
		"24;25":            0,
		"25;value(P2);P6":  0,
		"P3;6":             0,
		"P3;17":            0,
		"P3;26":            0,
		"6;7":              0,
		"7;value(P3);L11":  0,
		"17;18":            0,
		"18;value(P3);L11": 0,
		"18;value(P3);19":  0,
		"26;27":            0,
		"27;value(P3);P7":  0,
	}

	testDataFlow(t, dataFlowGraph, 32, 12, 0, expectedEdges)
}

func TestParamGlobalInstructions(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramGlobalInstructions.wasm")

	expectedEdges := map[string]int{
		"P0;0":           0,
		"0;1":            0,
		"1;value(P0);G0": 0,
		"P1;2":           0,
		"2;3":            0,
		"3;value(P1);G1": 0,
		"P2;4":           0,
		"4;5":            0,
		"5;value(P2);G2": 0,
		"P3;6":           0,
		"6;7":            0,
		"7;value(P3);G3": 0,
	}

	testDataFlow(t, dataFlowGraph, 12, 4, 4, expectedEdges)
}

func TestParamReturnInstructions0(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramReturnInstructions0.wasm")

	expectedEdges := map[string]int{
		"P0;3":     0,
		"P0;4":     0,
		"P0;7":     0,
		"3;return": 0,
		"7;return": 0,
		"4;5":      0,
	}

	testDataFlow(t, dataFlowGraph, 4, 1, 0, expectedEdges)
}

func TestParamReturnInstructions1(t *testing.T) {
	dataFlowGraph := loadDataFlowGraph(t, "../../../../test/wabesa/dataFlowGraph/dataInputs/parameters/paramReturnInstructions1.wasm")

	expectedEdges := map[string]int{
		"P0;0":     0,
		"0;1":      0,
		"P1;2":     0,
		"2;return": 0,
		"P2;4":     0,
		"4;return": 0,
	}

	testDataFlow(t, dataFlowGraph, 6, 3, 0, expectedEdges)
}

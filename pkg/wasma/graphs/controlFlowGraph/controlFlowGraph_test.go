package controlFlowGraph

import (
	"fmt"
	"log"
	"sort"
	"testing"
	"wasma/internal/test_utilities"
	"wasma/pkg/wasmp/instructions"
	"wasma/pkg/wasmp/parser"
)

func TestNoInstructions(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/noInstructions.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 0, len(cfgFunc1), "Check number of nodes")
}

func TestNoControlInstructions(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/noControlInstructions.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 1, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{0: "local.get", 1: "local.get", 2: "i32.add"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestBlock(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/block.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 2, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{1: "local.get", 2: "local.get", 3: "i32.add"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestNestedBlock(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/nestedBlock.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 3, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "block", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{2: "local.get", 3: "local.get", 4: "i32.add"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestLoop(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/loop.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 2, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "loop", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{1: "local.get", 2: "local.get", 3: "i32.add"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestNestedLoop(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/nestedLoop.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 3, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "loop", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "loop", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{2: "local.get", 3: "local.get", 4: "i32.add"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestLoopBrIf(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/loopBrIf.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 4, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "loop", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, false, []Edge{{5, ""}}, "", map[uint32]string{1: "local.get", 2: "i32.const", 3: "i32.sub", 4: "local.tee"})
		} else if index == 5 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}, {6, ""}}, "br_if", make(map[uint32]string))
		} else if index == 6 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{6: "nop"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestLoopBrTable(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/loopBrTable.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "loop", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{{6, ""}}, "", map[uint32]string{2: "local.get", 3: "i32.const", 4: "i32.sub", 5: "local.tee"})
		} else if index == 6 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}, {7, ""}}, "br_table", make(map[uint32]string))
		} else if index == 7 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{7: "nop"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestNestedLoopBrIf(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/nestedLoopBrIf.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "loop", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "loop", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{{6, ""}}, "", map[uint32]string{2: "local.get", 3: "i32.const", 4: "i32.sub", 5: "local.tee"})
		} else if index == 6 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}, {7, ""}}, "br_if", make(map[uint32]string))
		} else if index == 7 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{7: "nop"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestIfOnlyThen(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/ifOnlyThen.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, false, []Edge{{1, ""}}, "", map[uint32]string{0: "local.get"})
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, "then"}, {4, ""}}, "if", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{{3, ""}}, "", map[uint32]string{2: "local.get"})
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 4 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{4: "i32.const"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestIfOnlyElse(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/ifOnlyElse.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, false, []Edge{{1, ""}}, "", map[uint32]string{0: "local.get"})
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, "else"}, {4, ""}}, "if", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{{3, ""}}, "", map[uint32]string{2: "local.get"})
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 4 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{4: "i32.const"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestIfThenElseReturn(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/ifThenElseReturn.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 6, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, false, []Edge{{1, ""}}, "", map[uint32]string{0: "local.get"})
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, "then"}, {4, "else"}}, "if", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{{3, ""}}, "", map[uint32]string{2: "local.get"})
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 4 {
			checkCFGNode(t, index, node, false, []Edge{{5, ""}}, "", map[uint32]string{4: "local.get"})
		} else if index == 5 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestIfThenElse(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/ifThenElse.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, false, []Edge{{1, ""}}, "", map[uint32]string{0: "local.get"})
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, "then"}, {3, "else"}}, "if", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{{4, ""}}, "", map[uint32]string{2: "local.get"})
		} else if index == 3 {
			checkCFGNode(t, index, node, false, []Edge{{4, ""}}, "", map[uint32]string{3: "local.get"})
		} else if index == 4 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{4: "nop"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestNestedIf(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/nestedIf.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 28, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, false, []Edge{{1, ""}}, "", map[uint32]string{0: "local.get"})
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, "then"}, {17, "else"}}, "if", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{{3, ""}}, "", map[uint32]string{2: "local.get"})
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{{4, "then"}, {10, "else"}}, "if", make(map[uint32]string))
		} else if index == 4 {
			checkCFGNode(t, index, node, false, []Edge{{5, ""}}, "", map[uint32]string{4: "local.get"})
		} else if index == 5 {
			checkCFGNode(t, index, node, true, []Edge{{6, "then"}, {9, ""}}, "if", make(map[uint32]string))
		} else if index == 6 {
			checkCFGNode(t, index, node, false, []Edge{{9, ""}}, "", map[uint32]string{6: "nop", 7: "nop", 8: "nop"})
		} else if index == 9 {
			checkCFGNode(t, index, node, false, []Edge{{16, ""}}, "", map[uint32]string{9: "nop"})
		} else if index == 10 {
			checkCFGNode(t, index, node, false, []Edge{{11, ""}}, "", map[uint32]string{10: "local.get"})
		} else if index == 11 {
			checkCFGNode(t, index, node, true, []Edge{{12, "else"}, {15, ""}}, "if", make(map[uint32]string))
		} else if index == 12 {
			checkCFGNode(t, index, node, false, []Edge{{15, ""}}, "", map[uint32]string{12: "nop", 13: "nop", 14: "nop"})
		} else if index == 15 {
			checkCFGNode(t, index, node, false, []Edge{{16, ""}}, "", map[uint32]string{15: "nop"})
		} else if index == 16 {
			checkCFGNode(t, index, node, false, []Edge{{39, ""}}, "", map[uint32]string{16: "nop"})
		} else if index == 17 {
			checkCFGNode(t, index, node, false, []Edge{{18, ""}}, "", map[uint32]string{17: "local.get"})
		} else if index == 18 {
			checkCFGNode(t, index, node, true, []Edge{{19, "then"}, {29, "else"}}, "if", make(map[uint32]string))
		} else if index == 19 {
			checkCFGNode(t, index, node, false, []Edge{{20, ""}}, "", map[uint32]string{19: "local.get"})
		} else if index == 20 {
			checkCFGNode(t, index, node, true, []Edge{{21, "then"}, {25, "else"}}, "if", make(map[uint32]string))
		} else if index == 21 {
			checkCFGNode(t, index, node, false, []Edge{{24, ""}}, "", map[uint32]string{21: "nop", 22: "nop", 23: "nop"})
		} else if index == 24 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 25 {
			checkCFGNode(t, index, node, false, []Edge{{28, ""}}, "", map[uint32]string{25: "nop", 26: "nop", 27: "nop"})
		} else if index == 28 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 29 {
			checkCFGNode(t, index, node, false, []Edge{{30, ""}}, "", map[uint32]string{29: "local.get"})
		} else if index == 30 {
			checkCFGNode(t, index, node, true, []Edge{{31, "then"}, {34, "else"}}, "if", make(map[uint32]string))
		} else if index == 31 {
			checkCFGNode(t, index, node, false, []Edge{{37, ""}}, "", map[uint32]string{31: "nop", 32: "nop", 33: "nop"})
		} else if index == 34 {
			checkCFGNode(t, index, node, false, []Edge{{37, ""}}, "", map[uint32]string{34: "nop", 35: "nop", 36: "nop"})
		} else if index == 37 {
			checkCFGNode(t, index, node, false, []Edge{{38, ""}}, "", map[uint32]string{37: "nop"})
		} else if index == 38 {
			checkCFGNode(t, index, node, false, []Edge{{39, ""}}, "", map[uint32]string{38: "nop"})
		} else if index == 39 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{39: "nop"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestNestedIfContinuation(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/nestedIfContinuation.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 13, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, false, []Edge{{2, ""}}, "", map[uint32]string{1: "local.get"})
		} else if index == 2 {
			checkCFGNode(t, index, node, true, []Edge{{3, "then"}, {10, "else"}}, "if", make(map[uint32]string))
		} else if index == 3 {
			checkCFGNode(t, index, node, false, []Edge{{4, ""}}, "", map[uint32]string{3: "local.get"})
		} else if index == 4 {
			checkCFGNode(t, index, node, true, []Edge{{5, "then"}, {9, "else"}}, "if", make(map[uint32]string))
		} else if index == 5 {
			checkCFGNode(t, index, node, false, []Edge{{6, ""}}, "", map[uint32]string{5: "local.get"})
		} else if index == 6 {
			checkCFGNode(t, index, node, true, []Edge{{7, "then"}, {8, "else"}}, "if", make(map[uint32]string))
		} else if index == 7 {
			checkCFGNode(t, index, node, false, []Edge{{11, ""}}, "", map[uint32]string{7: "nop"})
		} else if index == 8 {
			checkCFGNode(t, index, node, false, []Edge{{11, ""}}, "", map[uint32]string{8: "nop"})
		} else if index == 9 {
			checkCFGNode(t, index, node, false, []Edge{{11, ""}}, "", map[uint32]string{9: "nop"})
		} else if index == 10 {
			checkCFGNode(t, index, node, false, []Edge{{11, ""}}, "", map[uint32]string{10: "nop"})
		} else if index == 11 {
			checkCFGNode(t, index, node, false, []Edge{{12, ""}}, "", map[uint32]string{11: "i32.const"})
		} else if index == 12 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestBr(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/br.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 23, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "block", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, true, []Edge{{3, ""}}, "block", make(map[uint32]string))
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{{4, ""}}, "block", make(map[uint32]string))
		} else if index == 4 {
			checkCFGNode(t, index, node, true, []Edge{{5, ""}}, "loop", make(map[uint32]string))
		} else if index == 5 {
			checkCFGNode(t, index, node, false, []Edge{{6, ""}}, "", map[uint32]string{5: "local.get"})
		} else if index == 6 {
			checkCFGNode(t, index, node, true, []Edge{{7, "then"}, {14, "else"}}, "if", make(map[uint32]string))
		} else if index == 7 {
			checkCFGNode(t, index, node, false, []Edge{{8, ""}}, "", map[uint32]string{7: "local.get"})
		} else if index == 8 {
			checkCFGNode(t, index, node, true, []Edge{{9, "then"}, {13, "else"}}, "if", make(map[uint32]string))
		} else if index == 9 {
			checkCFGNode(t, index, node, false, []Edge{{10, ""}}, "", map[uint32]string{9: "local.get"})
		} else if index == 10 {
			checkCFGNode(t, index, node, true, []Edge{{11, "then"}, {12, "else"}}, "if", make(map[uint32]string))
		} else if index == 11 {
			checkCFGNode(t, index, node, true, []Edge{{15, ""}}, "br", make(map[uint32]string))
		} else if index == 12 {
			checkCFGNode(t, index, node, true, []Edge{{17, ""}}, "br", make(map[uint32]string))
		} else if index == 13 {
			checkCFGNode(t, index, node, true, []Edge{{19, ""}}, "br", make(map[uint32]string))
		} else if index == 14 {
			checkCFGNode(t, index, node, true, []Edge{{21, ""}}, "br", make(map[uint32]string))
		} else if index == 15 {
			checkCFGNode(t, index, node, false, []Edge{{16, ""}}, "", map[uint32]string{15: "i32.const"})
		} else if index == 16 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 17 {
			checkCFGNode(t, index, node, false, []Edge{{18, ""}}, "", map[uint32]string{17: "i32.const"})
		} else if index == 18 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 19 {
			checkCFGNode(t, index, node, false, []Edge{{20, ""}}, "", map[uint32]string{19: "i32.const"})
		} else if index == 20 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 21 {
			checkCFGNode(t, index, node, false, []Edge{{22, ""}}, "", map[uint32]string{21: "i32.const"})
		} else if index == 22 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestBrIf(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/brIf.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 29, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "block", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, true, []Edge{{3, ""}}, "block", make(map[uint32]string))
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{{4, ""}}, "block", make(map[uint32]string))
		} else if index == 4 {
			checkCFGNode(t, index, node, true, []Edge{{5, ""}}, "loop", make(map[uint32]string))
		} else if index == 5 {
			checkCFGNode(t, index, node, false, []Edge{{6, ""}}, "", map[uint32]string{5: "local.get"})
		} else if index == 6 {
			checkCFGNode(t, index, node, true, []Edge{{7, "then"}, {17, "else"}}, "if", make(map[uint32]string))
		} else if index == 7 {
			checkCFGNode(t, index, node, false, []Edge{{8, ""}}, "", map[uint32]string{7: "local.get"})
		} else if index == 8 {
			checkCFGNode(t, index, node, true, []Edge{{9, "then"}, {15, "else"}}, "if", make(map[uint32]string))
		} else if index == 9 {
			checkCFGNode(t, index, node, false, []Edge{{10, ""}}, "", map[uint32]string{9: "local.get"})
		} else if index == 10 {
			checkCFGNode(t, index, node, true, []Edge{{11, "then"}, {13, "else"}}, "if", make(map[uint32]string))
		} else if index == 11 {
			checkCFGNode(t, index, node, false, []Edge{{12, ""}}, "", map[uint32]string{11: "local.get"})
		} else if index == 12 {
			checkCFGNode(t, index, node, true, []Edge{{19, ""}, {21, ""}}, "br_if", make(map[uint32]string))
		} else if index == 13 {
			checkCFGNode(t, index, node, false, []Edge{{14, ""}}, "", map[uint32]string{13: "local.get"})
		} else if index == 14 {
			checkCFGNode(t, index, node, true, []Edge{{19, ""}, {23, ""}}, "br_if", make(map[uint32]string))
		} else if index == 15 {
			checkCFGNode(t, index, node, false, []Edge{{16, ""}}, "", map[uint32]string{15: "local.get"})
		} else if index == 16 {
			checkCFGNode(t, index, node, true, []Edge{{19, ""}, {25, ""}}, "br_if", make(map[uint32]string))
		} else if index == 17 {
			checkCFGNode(t, index, node, false, []Edge{{18, ""}}, "", map[uint32]string{17: "local.get"})
		} else if index == 18 {
			checkCFGNode(t, index, node, true, []Edge{{19, ""}, {27, ""}}, "br_if", make(map[uint32]string))
		} else if index == 19 {
			checkCFGNode(t, index, node, false, []Edge{{20, ""}}, "", map[uint32]string{19: "i32.const"})
		} else if index == 20 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 21 {
			checkCFGNode(t, index, node, false, []Edge{{22, ""}}, "", map[uint32]string{21: "i32.const"})
		} else if index == 22 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 23 {
			checkCFGNode(t, index, node, false, []Edge{{24, ""}}, "", map[uint32]string{23: "i32.const"})
		} else if index == 24 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 25 {
			checkCFGNode(t, index, node, false, []Edge{{26, ""}}, "", map[uint32]string{25: "i32.const"})
		} else if index == 26 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 27 {
			checkCFGNode(t, index, node, false, []Edge{{28, ""}}, "", map[uint32]string{27: "i32.const"})
		} else if index == 28 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestBrTable(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/brTable.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 14, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "block", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, true, []Edge{{3, ""}}, "block", make(map[uint32]string))
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{{4, ""}}, "block", make(map[uint32]string))
		} else if index == 4 {
			checkCFGNode(t, index, node, false, []Edge{{5, ""}}, "", map[uint32]string{4: "local.get"})
		} else if index == 5 {
			checkCFGNode(t, index, node, true, []Edge{{6, ""}, {8, ""}, {10, ""}, {12, ""}}, "br_table", make(map[uint32]string))
		} else if index == 6 {
			checkCFGNode(t, index, node, false, []Edge{{7, ""}}, "", map[uint32]string{6: "i32.const"})
		} else if index == 7 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 8 {
			checkCFGNode(t, index, node, false, []Edge{{9, ""}}, "", map[uint32]string{8: "i32.const"})
		} else if index == 9 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 10 {
			checkCFGNode(t, index, node, false, []Edge{{11, ""}}, "", map[uint32]string{10: "i32.const"})
		} else if index == 11 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 12 {
			checkCFGNode(t, index, node, false, []Edge{{13, ""}}, "", map[uint32]string{12: "i32.const"})
		} else if index == 13 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestBrTableGroup(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/brTableGroup.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 14, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "block", make(map[uint32]string))
		} else if index == 2 {
			checkCFGNode(t, index, node, true, []Edge{{3, ""}}, "block", make(map[uint32]string))
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{{4, ""}}, "block", make(map[uint32]string))
		} else if index == 4 {
			checkCFGNode(t, index, node, false, []Edge{{5, ""}}, "", map[uint32]string{4: "local.get"})
		} else if index == 5 {
			checkCFGNode(t, index, node, true, []Edge{{6, ""}, {8, ""}, {10, ""}, {12, ""}}, "br_table", make(map[uint32]string))
		} else if index == 6 {
			checkCFGNode(t, index, node, false, []Edge{{7, ""}}, "", map[uint32]string{6: "i32.const"})
		} else if index == 7 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 8 {
			checkCFGNode(t, index, node, false, []Edge{{9, ""}}, "", map[uint32]string{8: "i32.const"})
		} else if index == 9 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 10 {
			checkCFGNode(t, index, node, false, []Edge{{11, ""}}, "", map[uint32]string{10: "i32.const"})
		} else if index == 11 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 12 {
			checkCFGNode(t, index, node, false, []Edge{{13, ""}}, "", map[uint32]string{12: "i32.const"})
		} else if index == 13 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestReturn(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/return.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 8, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
		} else if index == 1 {
			checkCFGNode(t, index, node, false, []Edge{{2, ""}}, "", map[uint32]string{1: "local.get"})
		} else if index == 2 {
			checkCFGNode(t, index, node, true, []Edge{{3, "then"}, {6, "else"}}, "if", make(map[uint32]string))
		} else if index == 3 {
			checkCFGNode(t, index, node, false, []Edge{{5, ""}}, "", map[uint32]string{3: "nop", 4: "i32.const"})
		} else if index == 5 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
		} else if index == 6 {
			checkCFGNode(t, index, node, false, []Edge{{9, ""}}, "", map[uint32]string{6: "nop", 7: "nop", 8: "nop"})
		} else if index == 9 {
			checkCFGNode(t, index, node, false, []Edge{{12, ""}}, "", map[uint32]string{9: "nop", 10: "nop", 11: "nop"})
		} else if index == 12 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{12: "i32.const"})
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestReturnUnreachable(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/returnUnreachable.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree
	unreachableCode, err := controlFlowGraph[0].GetUnreachableCode()
	if err != nil {
		t.Fatal(err.Error())
	}

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 11, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, false, []Edge{{1, ""}}, "", map[uint32]string{0: "local.get"})
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, "then"}, {5, "else"}}, "if", make(map[uint32]string))
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 2 {
			checkCFGNode(t, index, node, false, []Edge{{3, ""}}, "", map[uint32]string{2: "i32.const"})
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 3 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 4 {
			checkCFGNode(t, index, node, false, []Edge{{8, ""}}, "", map[uint32]string{4: "nop"})
			test_utilities.AssertTrueM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 5 {
			checkCFGNode(t, index, node, false, []Edge{{6, ""}}, "", map[uint32]string{5: "i32.const"})
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 6 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 7 {
			checkCFGNode(t, index, node, false, []Edge{{8, ""}}, "", map[uint32]string{7: "nop"})
			test_utilities.AssertTrueM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 8 {
			checkCFGNode(t, index, node, false, []Edge{{9, ""}}, "", map[uint32]string{8: "i32.const"})
			test_utilities.AssertTrueM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 9 {
			checkCFGNode(t, index, node, true, []Edge{}, "return", make(map[uint32]string))
			test_utilities.AssertTrueM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 10 {
			checkCFGNode(t, index, node, false, []Edge{}, "return", map[uint32]string{10: "nop"})
			test_utilities.AssertTrueM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestBrUnreachable(t *testing.T) {
	controlFlowGraph := loadControlFlowGraph(t, "../../../../test/wabesa/controlFlowGraph/brUnreachable.wasm")

	// every benchmark case contains only one function
	cfgFunc1 := controlFlowGraph[0].Tree
	unreachableCode, err := controlFlowGraph[0].GetUnreachableCode()
	if err != nil {
		t.Fatal(err.Error())
	}

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 9, len(cfgFunc1), "Check number of nodes")

	for index, node := range cfgFunc1 {
		if index == 0 {
			checkCFGNode(t, index, node, true, []Edge{{1, ""}}, "block", make(map[uint32]string))
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 1 {
			checkCFGNode(t, index, node, true, []Edge{{2, ""}}, "block", make(map[uint32]string))
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 2 {
			checkCFGNode(t, index, node, true, []Edge{{3, ""}}, "block", make(map[uint32]string))
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 3 {
			checkCFGNode(t, index, node, false, []Edge{{6, ""}}, "", map[uint32]string{3: "nop", 4: "nop", 5: "nop"})
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 6 {
			checkCFGNode(t, index, node, true, []Edge{{12, ""}}, "br", make(map[uint32]string))
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 7 {
			checkCFGNode(t, index, node, false, []Edge{{8, ""}}, "", map[uint32]string{7: "nop"})
			test_utilities.AssertTrueM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 8 {
			checkCFGNode(t, index, node, false, []Edge{{10, ""}}, "", map[uint32]string{8: "nop", 9: "nop"})
			test_utilities.AssertTrueM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 10 {
			checkCFGNode(t, index, node, false, []Edge{{12, ""}}, "", map[uint32]string{10: "nop", 11: "nop"})
			test_utilities.AssertTrueM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else if index == 12 {
			checkCFGNode(t, index, node, false, []Edge{}, "", map[uint32]string{12: "nop"})
			test_utilities.AssertFalseM(t, In(index, unreachableCode), fmt.Sprintf("node %v: check reachability", index))
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func loadControlFlowGraph(t *testing.T, file string) map[uint32]*CFG {
	module, err := parser.Parse(file)
	if err != nil {
		t.Fatal(err.Error())
	}

	return NewControlFlowGraph(module, true, 0)
}

func checkCFGNode(t *testing.T, nodeIndex uint32, node *CFGNode, isControl bool, successors []Edge, controlType string, block map[uint32]string) {
	test_utilities.AssertEqualM(t, isControl, node.Control, fmt.Sprintf("node %v: check isControl", nodeIndex))
	test_utilities.AssertTrueM(t, compareCFGEdgeSlice(successors, node.Successors), fmt.Sprintf("node %v: check successors", nodeIndex))
	if !node.Control {
		test_utilities.AssertTrueM(t, compareBlock(t, nodeIndex, block, node.Block), fmt.Sprintf("node %v: check block", nodeIndex))
	} else {
		test_utilities.AssertEqualM(t, controlType, node.Name, fmt.Sprintf("node %v: check controlType", nodeIndex))
	}
}

func compareBlock(t *testing.T, nodeIndex uint32, expected map[uint32]string, actual map[uint32]instructions.Instruction) bool {
	test_utilities.AssertEqualM(t, len(expected), len(actual), fmt.Sprintf("node %v: check block", nodeIndex))

	for i, expectedValue := range expected {
		if actualValue, found := actual[i]; found {
			if expectedValue != actualValue.Name() {
				return false
			}
		} else {
			for key, value := range actual {
				log.Printf("%v -> %v", key, value)
			}
			t.Fatalf("node %v: index %v is invalid", nodeIndex, i)
		}

	}
	return true
}

func compareCFGEdgeSlice(expected []Edge, actual []Edge) bool {
	if len(expected) != len(actual) {
		return false
	}

	sort.Slice(actual, func(i, j int) bool {
		return actual[i].TargetNode < actual[j].TargetNode
	})
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].TargetNode < expected[j].TargetNode
	})

	for i := 0; i < len(expected); i++ {
		if expected[i].TargetNode != actual[i].TargetNode ||
			expected[i].Tag != actual[i].Tag {
			return false
		}
	}

	return true
}

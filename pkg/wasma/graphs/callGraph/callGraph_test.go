package callGraph

import (
	"fmt"
	"sort"
	"testing"
	"wasma/internal/test_utilities"
	"wasma/pkg/wasmp/parser"
)

func TestEntrypointExportCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/entrypointExportCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 1, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 2, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{1, false}}, 0, true)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{0, false}}, []Function{}, 0, false)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestEntrypointExportIndirectCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/entrypointExportIndirectCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 1, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 2, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{1, true}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{0, true}}, 1, true)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestEntrypointStartCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/entrypointStartCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 1, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{3, false}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{3, false}}, []Function{}, 0, false)
		} else if index == 2 {
			checkNode(t, index, node, callGraph, []Function{{3, false}}, []Function{}, 0, false)
		} else if index == 3 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{{0, false}, {1, false}, {2, false}}, 0, false)
		} else if index == 4 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{3, false}}, 0, true)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestEntrypointStartIndirectCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/entrypointStartIndirectCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 1, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{3, true}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{3, true}}, []Function{}, 0, false)
		} else if index == 2 {
			checkNode(t, index, node, callGraph, []Function{{3, true}}, []Function{}, 0, false)
		} else if index == 3 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{{0, true}, {1, true}, {2, true}}, 4, false)
		} else if index == 4 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{3, false}}, 0, true)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestImportCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/importCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 0, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 2, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{1, false}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{0, false}}, 0, false)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestImportIndirectCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/importIndirectCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 0, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 2, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{1, true}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{0, true}}, 1, false)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestIndirectCallMultipleTargets(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/indirectCallMultipleTargets.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 0, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 4, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{6, true}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{6, true}}, []Function{}, 0, false)
		} else if index == 2 {
			checkNode(t, index, node, callGraph, []Function{{6, true}}, []Function{}, 0, false)
		} else if index == 6 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{0, true}, {1, true}, {2, true}}, 1, false)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestIndirectCallNoFittingTableEntry(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/indirectCallNoFittingTableEntry.wasm")

	// Check number of edges.
	test_utilities.AssertEqualM(t, 1, len(callGraph.GetEdges()), "Check number of edges.")
	for index, node := range callGraph.GetEdges() {
		if index == 6 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{}, 1, false)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestMultipleCalls(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/multipleCalls.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 1, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{}, 0, false)
		} else if index == 2 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{}, 0, false)
		} else if index == 3 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{}, 0, false)
		} else if index == 4 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{0, false}, {1, false}, {2, false}, {3, false}}, 0, true)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestMultipleIndirectCalls(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/multipleIndirectCalls.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 1, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 5, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{4, true}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{4, true}}, []Function{}, 0, false)
		} else if index == 2 {
			checkNode(t, index, node, callGraph, []Function{{4, true}}, []Function{}, 0, false)
		} else if index == 3 {
			checkNode(t, index, node, callGraph, []Function{{4, true}}, []Function{}, 0, false)
		} else if index == 4 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{0, true}, {1, true}, {2, true}, {3, true}}, 5, true)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestNestedCalls(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/nestedCalls.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 3, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 10, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{7, false}}, []Function{}, 0, false)
		} else if index == 2 {
			checkNode(t, index, node, callGraph, []Function{{9, false}}, []Function{{3, false}, {4, false}}, 0, false)
		} else if index == 3 {
			checkNode(t, index, node, callGraph, []Function{{2, false}}, []Function{}, 0, false)
		} else if index == 4 {
			checkNode(t, index, node, callGraph, []Function{{2, false}}, []Function{{0, false}, {5, false}, {6, false}}, 0, false)
		} else if index == 5 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{{7, false}}, 0, true)
		} else if index == 6 {
			checkNode(t, index, node, callGraph, []Function{{4, false}}, []Function{}, 0, true)
		} else if index == 7 {
			checkNode(t, index, node, callGraph, []Function{{5, false}}, []Function{{1, false}, {8, false}}, 0, false)
		} else if index == 8 {
			checkNode(t, index, node, callGraph, []Function{{7, false}}, []Function{}, 0, false)
		} else if index == 9 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{2, false}}, 0, true)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestNestedIndirectCalls(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/nestedIndirectCalls.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 3, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 10, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{4, true}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{7, true}}, []Function{}, 0, false)
		} else if index == 2 {
			checkNode(t, index, node, callGraph, []Function{{9, true}}, []Function{{3, true}, {4, true}}, 2, false)
		} else if index == 3 {
			checkNode(t, index, node, callGraph, []Function{{2, true}}, []Function{}, 0, false)
		} else if index == 4 {
			checkNode(t, index, node, callGraph, []Function{{2, true}}, []Function{{0, true}, {5, true}, {6, true}}, 3, false)
		} else if index == 5 {
			checkNode(t, index, node, callGraph, []Function{{4, true}}, []Function{{7, true}}, 1, true)
		} else if index == 6 {
			checkNode(t, index, node, callGraph, []Function{{4, true}}, []Function{}, 0, true)
		} else if index == 7 {
			checkNode(t, index, node, callGraph, []Function{{5, true}}, []Function{{1, true}, {8, true}}, 2, false)
		} else if index == 8 {
			checkNode(t, index, node, callGraph, []Function{{7, true}}, []Function{}, 0, false)
		} else if index == 9 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{2, true}}, 1, true)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestNoCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/noCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 0, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 0, len(callGraph.GetEdges()), "Check number of nodes")

}

func TestOneCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/oneCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 0, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 2, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{1, false}}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{{0, false}}, []Function{}, 0, false)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func TestOneIndirectCall(t *testing.T) {
	callGraph := loadCallGraph(t, "../../../../test/wabesa/callGraph/oneIndirectCall.wasm")

	// Check number of entry points.
	test_utilities.AssertEqualM(t, 0, len(callGraph.EntryPoints), "Check number of entry points.")

	// Check number of nodes.
	test_utilities.AssertEqualM(t, 2, len(callGraph.GetEdges()), "Check number of nodes")

	for index, node := range callGraph.GetEdges() {
		if index == 0 {
			checkNode(t, index, node, callGraph, []Function{{1, true}}, []Function{}, 0, false)
		} else if index == 1 {
			checkNode(t, index, node, callGraph, []Function{}, []Function{{0, true}}, 1, false)
		} else {
			t.Fatalf("Got node index %v but in the corresponding graph is no node with this index.", index)
		}
	}
}

func loadCallGraph(t *testing.T, file string) *CallGraph {
	module, err := parser.Parse(file)
	if err != nil {
		t.Fatal(err.Error())
	}

	callGraph, err := NewCallGraph(module, true)
	if err != nil {
		t.Fatal(err.Error())
	}

	return callGraph
}

func checkNode(t *testing.T, nodeIndex uint32, node *CallGraphEdges, callGraph *CallGraph, expectedIsCalled []Function, expectedCalls []Function, indirectCalls uint32, isEntryPoint bool) {
	test_utilities.AssertTrueM(t, compareCGEdgeSlice(expectedIsCalled, node.IsCalled), fmt.Sprintf("node %v: check IsCalled", nodeIndex))
	test_utilities.AssertTrueM(t, compareCGEdgeSlice(expectedCalls, node.Calls), fmt.Sprintf("node %v: check Calls", nodeIndex))
	test_utilities.AssertEqualM(t, isEntryPoint, callGraph.IsEntryPoint(nodeIndex), fmt.Sprintf("node %v: check IsEntryPoint", nodeIndex))
	test_utilities.AssertEqualM(t, indirectCalls, node.NumIndirect, fmt.Sprintf("node %v: check the number of indirect calls", nodeIndex))
}

func compareCGEdgeSlice(expected []Function, actual []Function) bool {
	if len(expected) != len(actual) {
		return false
	}

	sort.Slice(actual, func(i, j int) bool {
		return actual[i].FuncIdx < actual[j].FuncIdx
	})
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].FuncIdx < expected[j].FuncIdx
	})

	for i := 0; i < len(expected); i++ {
		if expected[i].FuncIdx != actual[i].FuncIdx ||
			expected[i].IsIndirect != actual[i].IsIndirect {
			return false
		}
	}

	return true
}

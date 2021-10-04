package dataFlowGraph

import (
	"fmt"
	"testing"
	"wasma/internal/test_utilities"
	"wasma/pkg/wasmp/parser"
)

func testDataFlow(t *testing.T, dataFlowGraph map[uint32]*DFG, informationFlow int, locals int, globals int, expectedEdges map[string]int) {
	// The function with index 0 contains the evaluation scenario.
	dfgFunc1 := dataFlowGraph[0]

	// Check information flow.
	test_utilities.AssertEqualM(t, informationFlow, len(dfgFunc1.Tree), "Check information flow")

	// Check number of locals
	test_utilities.AssertEqualM(t, locals, len(dfgFunc1.Environment.Locals), "Check number of locals")

	// Check number of globals
	test_utilities.AssertEqualM(t, globals, len(dfgFunc1.Environment.Globals), "Check number of globals")

	for _, flowEdges := range dfgFunc1.Tree {
		for _, flowEdge := range flowEdges {
			checkFlowEdge(t, flowEdge, expectedEdges)
		}
	}

	for key, value := range expectedEdges {
		if value < 1 {
			t.Fatalf("Edge is missing: %v", key)
		} else if value > 1 {
			t.Fatalf("Edge exists more than once: %v", key)
		}
	}
}

func loadDataFlowGraph(t *testing.T, file string) map[uint32]*DFG {
	module, err := parser.Parse(file)
	if err != nil {
		t.Fatal(err.Error())
	}

	return NewDataFlowGraph(module, true, 0)
}

func checkFlowEdge(t *testing.T, flowEdge FlowEdge, expectedEdges map[string]int) {
	key := ""

	if flowEdge.Variable.Value == "unknown" {
		key = fmt.Sprintf("%v;%v", flowEdge.Output, flowEdge.Input)
	} else {
		key = fmt.Sprintf("%v;%v;%v", flowEdge.Output, flowEdge.Variable.Value, flowEdge.Input)
	}

	if value, found := expectedEdges[key]; found {
		expectedEdges[key] = value + 1
	} else {
		t.Fatalf("Invalid flow edge: %v -%v-> %v", flowEdge.Output, flowEdge.Variable.Value, flowEdge.Input)
	}
}

package structures

import (
	"wasma/pkg/wasmp/types"
)

type MemoryEntry struct {
	Value   string
	Tainted bool
}

type FunctionCall struct {
	FuncIdx     uint32
	Instruction string
	Name        string
}

type Taint struct {
	Tainted bool
	Source  FunctionCall
	Sink    FunctionCall
}

type Variable struct {
	// Types:
	// - P = parameter
	// - L = local
	// - GM = global (mutable)
	// - GC = global (constant)
	// - V = variable
	// - C = constant
	VariableType       string
	PrimaryVariableIdx uint32
	VariableName       string
	VariableDataType   string
	Value              string
	LocalGlobalIn      bool
	Taint              Taint
}

type Environment interface {
	GetNumOfFuncParameterFuncIdx(funcIdx uint32) (*types.FunctionType, error)
	GetNumOfFuncParameterTypeIdx(typeIdx uint32) (*types.FunctionType, error)
	Pop() Variable
	Push(element Variable)
	SetLocal(localIdx uint32, newValue Variable) (Variable, Variable, error)
	GetLocal(localIdx uint32) (Variable, Variable, error)
	SetGlobal(globalIdx uint32, newValue Variable) (Variable, Variable, error)
	GetGlobal(globalIdx uint32) (Variable, Variable, error)
	NewParameter(variableDataType string) Variable
	NewParameterWithTaint(variableDataType string, taint Taint) Variable
	NewLocal(variableDataType string) Variable
	NewGlobal(mut byte, variableDataType string) Variable
	NewVariable(variableDataType string) Variable
	NewConstant(value string, variableDataType string) Variable
	AddInput(instrIdx uint32, input Variable)
	AddOutput(instrIdx uint32, output Variable)
}

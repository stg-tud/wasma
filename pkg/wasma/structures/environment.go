package structures

import (
	"errors"
	"fmt"
	"wasma/pkg/wasmp/modules"
	"wasma/pkg/wasmp/structures"
	"wasma/pkg/wasmp/types"
)

type DataFlowEdge struct {
	Input  []structures.Variable
	Output []structures.Variable
}

type Environment struct {
	module             *modules.Module
	primaryVariableIdx uint32
	localIdx           uint32
	globalIdx          uint32
	variableIdx        uint32
	constantIdx        uint32
	returnIdx          uint32
	Stack              Stack
	// key: localIdx
	Locals map[uint32]structures.Variable
	// key: globalIdx
	Globals map[uint32]structures.Variable
	// key: primaryVariableIdx
	Variables map[uint32]structures.Variable
	// key: primaryVariableIdx, value: returnPoint
	ReturnPoints map[uint32]string
	// key: instrIdx
	Flow map[uint32]DataFlowEdge
	// key: Adress in memory
	Memory                map[uint32]structures.MemoryEntry
	IsEntireMemoryTainted bool
}

func (environment *Environment) SetGlobalIdx(newValue uint32) {
	environment.globalIdx = newValue
}

func (environment *Environment) getNumOfFuncParameterFuncIdxCheckImportFunction(funcIdx uint32) (*types.FunctionType, error) {
	if importSection, err := environment.module.GetImportSection(); err == nil {
		if imp, found := importSection.FuncImports[funcIdx]; found {
			if imp.ImportDesc.ImportType == 0x00 {
				if typeSection, err := environment.module.GetTypeSection(); err == nil {
					if functionType, found := typeSection.FunctionTypes[imp.ImportDesc.TypeIdx]; found {
						return &functionType, nil
					} else {
						return nil, errors.New(fmt.Sprintf("function type for typeIdx %v not found", imp.ImportDesc.TypeIdx))
					}
				} else {
					return nil, errors.New("type section does not exist")
				}
			} else {
				return nil, errors.New("function import has no valid import type")
			}
		} else {
			return nil, errors.New("import section does not exist")
		}
	} else {
		return nil, errors.New("import section does not exist")
	}
}

func (environment *Environment) GetNumOfFuncParameterFuncIdx(funcIdx uint32) (*types.FunctionType, error) {
	if functionSection, err := environment.module.GetFunctionSection(); err == nil {
		if typeSection, err := environment.module.GetTypeSection(); err == nil {
			if typeIdx, found := functionSection.TypeIdxs[funcIdx]; found {
				if functionType, found := typeSection.FunctionTypes[typeIdx]; found {
					return &functionType, nil
				} else {
					return nil, errors.New(fmt.Sprintf("function type for typeIdx %v not found", typeIdx))
				}
			} else {
				return environment.getNumOfFuncParameterFuncIdxCheckImportFunction(funcIdx)
			}
		} else {
			return nil, errors.New("type section does not exist")
		}
	} else {
		return nil, errors.New("function section does not exist")
	}
}

func (environment *Environment) GetNumOfFuncParameterTypeIdx(typeIdx uint32) (*types.FunctionType, error) {
	if typeSection, err := environment.module.GetTypeSection(); err == nil {
		if functionType, found := typeSection.FunctionTypes[typeIdx]; found {
			return &functionType, nil
		} else {
			return nil, errors.New(fmt.Sprintf("function type for typeIdx %v not found", typeIdx))
		}
	} else {
		return nil, errors.New("type section does not exist")
	}
}

func (environment *Environment) Pop() structures.Variable {
	return environment.Stack.Pop()
}

func (environment *Environment) Push(element structures.Variable) {
	environment.Stack.Push(element)
}

func (environment *Environment) SetLocal(localIdx uint32, newValue structures.Variable) (structures.Variable, structures.Variable, error) {
	if local, found := environment.Locals[localIdx]; found {
		if newValue.Value == "unknown" && newValue.VariableType != "C" && newValue.VariableType != "V" {
			local.Value = fmt.Sprintf("value(%v)", newValue.VariableName)
		} else {
			local.Value = newValue.Value
		}
		environment.Locals[localIdx] = local

		originalLocal, currentLocal, _ := environment.GetLocal(localIdx)
		return originalLocal, currentLocal, nil
	} else {
		return structures.Variable{}, structures.Variable{}, errors.New(fmt.Sprintf("localIdx: %v is no valid index", localIdx))
	}
}

func (environment *Environment) GetLocal(localIdx uint32) (structures.Variable, structures.Variable, error) {
	if originalValue, found := environment.Locals[localIdx]; found {
		currentValue := originalValue
		currentValue.PrimaryVariableIdx = environment.primaryVariableIdx
		environment.primaryVariableIdx++
		return originalValue, currentValue, nil
	}
	return structures.Variable{}, structures.Variable{}, errors.New(fmt.Sprintf("localIdx: %v is no valid index", localIdx))
}

func (environment *Environment) SetGlobal(globalIdx uint32, newValue structures.Variable) (structures.Variable, structures.Variable, error) {
	if global, found := environment.Globals[globalIdx]; found {
		if newValue.Value == "unknown" && newValue.VariableType != "C" && newValue.VariableType != "V" {
			global.Value = fmt.Sprintf("value(%v)", newValue.VariableName)
		} else {
			global.Value = newValue.Value
		}
		environment.Globals[globalIdx] = global

		originalGlobal, currentGlobal, _ := environment.GetGlobal(globalIdx)
		return originalGlobal, currentGlobal, nil
	} else {
		return structures.Variable{}, structures.Variable{}, errors.New(fmt.Sprintf("globalIdx: %v is no valid index", globalIdx))
	}
}

func (environment *Environment) GetGlobal(globalIdx uint32) (structures.Variable, structures.Variable, error) {
	if originalValue, found := environment.Globals[globalIdx]; found {
		currentValue := originalValue
		currentValue.PrimaryVariableIdx = environment.primaryVariableIdx
		environment.primaryVariableIdx++
		return originalValue, currentValue, nil
	}
	return structures.Variable{}, structures.Variable{}, errors.New(fmt.Sprintf("globalIdx: %v is no valid index", globalIdx))
}

func (environment *Environment) NewParameter(variableDataType string) structures.Variable {
	taint := structures.Taint{Tainted: false}
	newParameter := environment.NewParameterWithTaint(variableDataType, taint)
	return newParameter
}

func (environment *Environment) NewParameterWithTaint(variableDataType string, taint structures.Taint) structures.Variable {
	newParameter := structures.Variable{
		VariableType:       "P",
		PrimaryVariableIdx: environment.primaryVariableIdx,
		VariableName:       fmt.Sprintf("P%v", environment.localIdx),
		VariableDataType:   variableDataType,
		Value:              "unknown",
		LocalGlobalIn:      false,
		Taint:              taint}
	environment.Variables[environment.primaryVariableIdx] = newParameter
	environment.Locals[environment.localIdx] = newParameter
	environment.primaryVariableIdx++
	environment.localIdx++
	return newParameter
}

func (environment *Environment) NewLocal(variableDataType string) structures.Variable {
	taint := structures.Taint{Tainted: false}
	newLocal := structures.Variable{
		VariableType:       "L",
		PrimaryVariableIdx: environment.primaryVariableIdx,
		VariableName:       fmt.Sprintf("L%v", environment.localIdx),
		VariableDataType:   variableDataType,
		Value:              "unknown",
		LocalGlobalIn:      false,
		Taint:              taint}
	environment.Variables[environment.primaryVariableIdx] = newLocal
	environment.Locals[environment.localIdx] = newLocal
	environment.primaryVariableIdx++
	environment.localIdx++
	return newLocal
}

func (environment *Environment) NewGlobal(mut byte, variableDataType string) structures.Variable {
	var variableType string
	if mut == 0x00 {
		variableType = "GC"
	} else {
		variableType = "GM"
	}
	taint := structures.Taint{Tainted: false}
	newGlobal := structures.Variable{
		VariableType:       variableType,
		PrimaryVariableIdx: environment.primaryVariableIdx,
		VariableName:       fmt.Sprintf("G%v", environment.globalIdx),
		VariableDataType:   variableDataType,
		Value:              "unknown",
		LocalGlobalIn:      false,
		Taint:              taint}
	environment.Variables[environment.primaryVariableIdx] = newGlobal
	environment.Globals[environment.globalIdx] = newGlobal
	environment.primaryVariableIdx++
	environment.globalIdx++
	return newGlobal
}
func (environment *Environment) NewVariable(variableDataType string) structures.Variable {
	taint := structures.Taint{Tainted: false}
	newVariable := structures.Variable{
		VariableType:       "V",
		PrimaryVariableIdx: environment.primaryVariableIdx,
		VariableName:       fmt.Sprintf("V%v", environment.variableIdx),
		VariableDataType:   variableDataType,
		Value:              "unknown",
		LocalGlobalIn:      false,
		Taint:              taint}
	environment.Variables[environment.primaryVariableIdx] = newVariable
	environment.primaryVariableIdx++
	environment.variableIdx++
	return newVariable
}
func (environment *Environment) NewConstant(value string, variableDataType string) structures.Variable {
	taint := structures.Taint{Tainted: false}
	newConstant := structures.Variable{
		VariableType:       "C",
		PrimaryVariableIdx: environment.primaryVariableIdx,
		VariableName:       fmt.Sprintf("C%v", environment.constantIdx),
		VariableDataType:   variableDataType,
		Value:              value,
		LocalGlobalIn:      false,
		Taint:              taint}
	environment.Variables[environment.primaryVariableIdx] = newConstant
	environment.primaryVariableIdx++
	environment.constantIdx++
	return newConstant
}

func (environment *Environment) AddInput(instrIdx uint32, input structures.Variable) {
	if dataFlowEdge, found := environment.Flow[instrIdx]; found {
		dataFlowEdge.Input = append(dataFlowEdge.Input, input)
		environment.Flow[instrIdx] = dataFlowEdge
	} else {
		environment.Flow[instrIdx] = DataFlowEdge{[]structures.Variable{input}, []structures.Variable{}}
	}
}

func (environment *Environment) AddOutput(instrIdx uint32, output structures.Variable) {
	if dataFlowEdge, found := environment.Flow[instrIdx]; found {
		dataFlowEdge.Output = append(dataFlowEdge.Output, output)
		environment.Flow[instrIdx] = dataFlowEdge
	} else {
		environment.Flow[instrIdx] = DataFlowEdge{[]structures.Variable{}, []structures.Variable{output}}
	}
}

func (environment *Environment) NewReturnPoint() {
	for _, element := range environment.Stack.GetStack() {
		environment.ReturnPoints[element.PrimaryVariableIdx] = fmt.Sprintf("R%v", environment.returnIdx)
		environment.returnIdx++
	}
}

func NewEnvironment(module *modules.Module) *Environment {
	environment := new(Environment)
	environment.module = module
	environment.primaryVariableIdx = 0
	environment.localIdx = 0
	environment.globalIdx = 0
	environment.variableIdx = 0
	environment.returnIdx = 0

	// key: localIdx
	environment.Locals = make(map[uint32]structures.Variable)
	// key: globalIdx
	environment.Globals = make(map[uint32]structures.Variable)
	// key: primaryVariableIdx
	environment.Variables = make(map[uint32]structures.Variable)
	// key: primaryVariableIdx, value: returnPoint
	environment.ReturnPoints = make(map[uint32]string)
	// key: instrIdx
	environment.Flow = make(map[uint32]DataFlowEdge)

	return environment
}

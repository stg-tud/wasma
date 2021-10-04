package instructions

import (
	"fmt"
	"io"
	"log"
	"wasma/pkg/wasmp/structures"
	"wasma/pkg/wasmp/values"
)

func newDrop(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{
		0x1A,
		"drop",
		values.GetPosition(),
		func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
			variable := environment.Pop()
			environment.AddInput(instrIdx, variable)
		}}, nil
}

func newSelect(reader io.Reader) (Instruction, error) {
	return &SimpleInstruction{
		0x1B,
		"select",
		values.GetPosition(),
		func(instrIdx uint32, instruction Instruction, environment structures.Environment) {
			constant := environment.Pop()
			value2 := environment.Pop()
			value1 := environment.Pop()

			if value2.VariableDataType != value1.VariableDataType {
				log.Fatalf("error: select instruction %v, %v != %v", instrIdx, value2.VariableDataType, value1.VariableDataType)
			}

			newVariable := environment.NewVariable(value1.VariableDataType)
			if value1.Value != "unknown" || value2.Value != "unknown" {
				newVariable.Value = fmt.Sprintf("%v or %v", value1.Value, value2.Value)
			} else {
				if value1.VariableType == "P" || value1.VariableType == "L" || value1.VariableType == "GM" || value1.VariableType == "GC" {
					newVariable.Value = fmt.Sprintf("value(%v)", value1.VariableName)
				} else {
					newVariable.Value = "unknown"
				}
				if value2.VariableType == "P" || value2.VariableType == "L" || value2.VariableType == "GM" || value2.VariableType == "GC" {
					newVariable.Value = newVariable.Value + " or " + fmt.Sprintf("value(%v)", value2.VariableName)
				} else {
					if newVariable.Value != "unknown" {
						newVariable.Value = newVariable.Value + " or unknown"
					}
				}
			}
			environment.Push(newVariable)

			environment.AddInput(instrIdx, constant)
			environment.AddInput(instrIdx, value2)
			environment.AddInput(instrIdx, value1)
			environment.AddOutput(instrIdx, newVariable)
		}}, nil
}

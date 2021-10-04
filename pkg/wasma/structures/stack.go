package structures

import (
	"wasma/pkg/wasmp/structures"
)

type Stack struct {
	elements []structures.Variable
}

func (stack *Stack) Pop() structures.Variable {
	element := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return element
}

func (stack *Stack) Push(element structures.Variable) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack) GetStack() []structures.Variable {
	return stack.elements
}

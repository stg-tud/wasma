package instructions

import (
	"io"
	"wasma/pkg/wasmp/values"
)

type Expr struct {
	Instructions []Instruction
}

func NewExpr(reader io.Reader) (*Expr, error) {
	expr := new(Expr)

	for {
		nextByte, err := values.ReadNextByte(reader)
		if err != nil {
			return nil, err
		}

		if nextByte == 0x0B {
			values.IncrementEndCounter()
			break
		}

		instruction, err := mapOpcode(nextByte, reader)
		if err != nil {
			return nil, err
		}

		expr.Instructions = append(expr.Instructions, instruction)
	}

	return expr, nil
}

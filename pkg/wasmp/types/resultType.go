package types

import (
	"io"
	"wasma/pkg/wasmp/values"
)

type ResultType struct {
	ValType []string
}

func NewResultType(reader io.Reader) (*ResultType, error) {
	resultType := new(ResultType)
	vecLen, err := values.ReadU32(reader)
	if err != nil {return nil, err}

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		valType, err := NewValType(reader)
		if err != nil {
			return resultType, err
		}
		resultType.ValType = append(resultType.ValType, valType)
	}
	return resultType, nil
}

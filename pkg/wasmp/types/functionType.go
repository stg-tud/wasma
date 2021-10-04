package types

import (
	"errors"
	"fmt"
	"io"
	values2 "wasma/pkg/wasmp/values"
)

type FunctionType struct {
	ParameterTypes []string
	ResultTypes    []string
}

func (functionType *FunctionType) ParameterTypesToString() string {
	parameterTypesString := ""
	for index, parameterType := range functionType.ParameterTypes {
		if index == 0 {
			parameterTypesString = parameterTypesString + parameterType
		} else {
			parameterTypesString = parameterTypesString + ", " + parameterType
		}
	}
	return parameterTypesString
}

// WebAssembly version 1.0 supports only one return value
func (functionType *FunctionType) ResultTypeToString() string {
	if functionType.ResultTypes == nil {
		return "nil"
	} else if len(functionType.ResultTypes) == 1 {
		return functionType.ResultTypes[0]
	} else {
		resultTypesString := "("
		for index, resultType := range functionType.ResultTypes {
			if index == 0 {
				resultTypesString = resultTypesString + resultType
			} else {
				resultTypesString = resultTypesString + ", " + resultType
			}
		}
		return resultTypesString + ")"
	}
}

func NewFunctionType(reader io.Reader) (FunctionType, error) {
	nextByte, err := values2.ReadNextByte(reader)
	if err != nil {
		return FunctionType{}, err
	}

	if nextByte != 0x60 {
		return FunctionType{}, errors.New(fmt.Sprintf("Error while reading type section. Expected 0x60 but got: %x", nextByte))
	}

	// br parameter types
	parameterTypes, err := readType(reader)
	if err != nil {
		return FunctionType{}, err
	}

	// br return types
	resultTypes, err := readType(reader)
	if err != nil {
		return FunctionType{}, err
	}

	return FunctionType{ParameterTypes: parameterTypes, ResultTypes: resultTypes}, nil
}

func readType(reader io.Reader) ([]string, error) {
	n, err := values2.ReadU32(reader)
	if err != nil {
		return []string{}, err
	}

	var valtypes []string

	var i uint32 = 1
	for ; i <= n; i++ {
		valtype, err := NewValType(reader)
		if err != nil {
			return []string{}, err
		}
		valtypes = append(valtypes, valtype)
	}
	return valtypes, nil
}

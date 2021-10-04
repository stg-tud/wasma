package sections

import (
	"errors"
	"io"
	"log"
	"wasma/pkg/wasmp/instructions"
	"wasma/pkg/wasmp/types"
	values2 "wasma/pkg/wasmp/values"
)

type Code struct {
	Size     uint32
	Function *Function
}

func newCode(reader io.Reader) (*Code, error) {
	var err error
	code := new(Code)
	code.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	code.Function, err = newFunction(reader)
	if err != nil {
		return nil, err
	}

	return code, nil
}

type Function struct {
	Locals []*Local
	Expr   *instructions.Expr
}

func newFunction(reader io.Reader) (*Function, error) {
	function := new(Function)
	vecLen, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		local, err := newLocal(reader)
		if err != nil {
			return nil, err
		}
		function.Locals = append(function.Locals, local)
	}

	function.Expr, err = instructions.NewExpr(reader)
	if err != nil {
		return nil, err
	}

	return function, nil
}

type Local struct {
	N       uint32 // Number of locals with the corresponding type ValType.
	ValType string
}

func newLocal(reader io.Reader) (*Local, error) {
	var err error
	local := new(Local)
	local.N, err = values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	local.ValType, err = types.NewValType(reader)
	if err != nil {
		return nil, err
	}

	return local, nil
}

type CodeSection struct {
	Id   byte
	Size uint32
	// key: funcIdx
	Codes        map[uint32]*Code
	StartSection uint32
	StartContent uint32
}

// Parses a code section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewCodeSection(reader io.Reader) (*CodeSection, error) {
	ResetFunctionIndexAfterImport()
	var err error
	codeSection := new(CodeSection)
	codeSection.Id = 0x0A
	// key: funcIdx
	codeSection.Codes = make(map[uint32]*Code)
	codeSection.StartSection = values2.GetByteCounter()

	codeSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	codeSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, codeSection.Size)
	if err != nil {
		values2.UpdateByteCounter(codeSection.StartContent - 1 + codeSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(codeSection.StartContent - 1 + codeSection.Size)
		return nil, err
	}

	log.Printf("Identified code section: Id = %x, start = %x, Size = %v, vecL: %v. Parsing...\n",
		codeSection.Id,
		codeSection.StartSection,
		codeSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		code, err := newCode(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(codeSection.StartContent - 1 + codeSection.Size)
			return nil, err
		}
		codeSection.Codes[NewFunctionIndex()] = code
	}

	err = contentCompletelyRead("code", codeSection.Size, codeSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(codeSection.StartContent - 1 + codeSection.Size)
		return nil, err
	}

	return codeSection, nil
}

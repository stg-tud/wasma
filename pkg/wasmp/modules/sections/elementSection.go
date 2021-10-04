package sections

import (
	"errors"
	"io"
	"log"
	"wasma/pkg/wasmp/instructions"
	values2 "wasma/pkg/wasmp/values"
)

type Element struct {
	TableIdx uint32
	Expr     *instructions.Expr
	FuncIdxs []uint32
}

func newElement(reader io.Reader) (*Element, error) {
	var err error
	element := new(Element)
	element.TableIdx, err = values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	element.Expr, err = instructions.NewExpr(reader)
	if err != nil {
		return nil, err
	}

	vecLen, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		funcIdx, err := values2.ReadU32(reader)
		if err != nil {
			return nil, err
		}
		element.FuncIdxs = append(element.FuncIdxs, funcIdx)
	}

	return element, nil
}

type ElementSection struct {
	Id           byte
	Size         uint32
	Elements     []*Element
	StartSection uint32
	StartContent uint32
	// Deviation from specification for easier processing:
	// Additional field
	// key: funcIdx
	FuncIdxs map[uint32]bool
}

// Returns a list of all funcIdxs that can be used
// for an indirect call.
func getTableFuncIdxs(elementSection *ElementSection) map[uint32]bool {
	// key: funcIdx
	var funcIdxs = make(map[uint32]bool)
	for _, element := range elementSection.Elements {
		for _, funcIdx := range element.FuncIdxs {
			funcIdxs[funcIdx] = true
		}
	}

	return funcIdxs
}

// Parses an element section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewElementSection(reader io.Reader) (*ElementSection, error) {
	var err error
	elementSection := new(ElementSection)
	elementSection.Id = 0x09
	elementSection.StartSection = values2.GetByteCounter()

	elementSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	elementSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, elementSection.Size)
	if err != nil {
		values2.UpdateByteCounter(elementSection.StartContent - 1 + elementSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(elementSection.StartContent - 1 + elementSection.Size)
		return nil, err
	}

	log.Printf("Identified element section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		elementSection.Id,
		elementSection.StartSection,
		elementSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		element, err := newElement(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(elementSection.StartContent - 1 + elementSection.Size)
			return nil, err
		}
		elementSection.Elements = append(elementSection.Elements, element)
	}

	err = contentCompletelyRead("element", elementSection.Size, elementSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(elementSection.StartContent - 1 + elementSection.Size)
		return nil, err
	}

	elementSection.FuncIdxs = getTableFuncIdxs(elementSection)

	return elementSection, nil
}

package sections

import (
	"errors"
	"fmt"
	"io"
	"log"
	values2 "wasma/pkg/wasmp/values"
)

type ExportDesc struct {
	ExportType byte
	FuncIdx    uint32
	TableIdx   uint32
	MemIdx     uint32
	GlobalIdx  uint32
}

func (exportDesc *ExportDesc) GetIdx() (uint32, string, error) {
	switch exportDesc.ExportType {
	case 0x00:
		return exportDesc.FuncIdx, "funcIdx", nil
	case 0x01:
		return exportDesc.TableIdx, "tableIdx", nil
	case 0x02:
		return exportDesc.MemIdx, "memIdx", nil
	case 0x03:
		return exportDesc.GlobalIdx, "globalIdx", nil
	default:
		return 0, "", errors.New("invalid exportdesc")
	}
}

func (exportDesc *ExportDesc) GetIdxValue() string {
	switch exportDesc.ExportType {
	case 0x00:
		return fmt.Sprintf("%v", exportDesc.FuncIdx)
	case 0x01:
		return fmt.Sprintf("%v", exportDesc.TableIdx)
	case 0x02:
		return fmt.Sprintf("%v", exportDesc.MemIdx)
	case 0x03:
		return fmt.Sprintf("%v", exportDesc.GlobalIdx)
	default:
		return ""
	}
}

func newExportDesc(reader io.Reader) (*ExportDesc, error) {
	var err error
	exportDesc := new(ExportDesc)
	exportDesc.ExportType, err = values2.ReadNextByte(reader)
	if err != nil {
		return nil, err
	}

	idx, err := values2.ReadU32(reader)
	if err != nil {
		return nil, err
	}
	switch exportDesc.ExportType {
	case 0x00:
		exportDesc.FuncIdx = idx
	case 0x01:
		exportDesc.TableIdx = idx
	case 0x02:
		exportDesc.MemIdx = idx
	case 0x03:
		exportDesc.GlobalIdx = idx
	default:
		log.Fatalf("Error while reading ExportDesc. No valid identifier, got: %x", exportDesc.ExportType)
	}

	return exportDesc, nil
}

type Export struct {
	Name       string
	ExportDesc *ExportDesc
}

func (export *Export) ToString() string {
	switch export.ExportDesc.ExportType {
	case 0x00:
		return fmt.Sprintf("func[%v] -> \"%v\"", export.ExportDesc.GetIdxValue(), export.Name)
	case 0x01:
		return fmt.Sprintf("table[%v] -> \"%v\"", export.ExportDesc.GetIdxValue(), export.Name)
	case 0x02:
		return fmt.Sprintf("memory[%v] -> \"%v\"", export.ExportDesc.GetIdxValue(), export.Name)
	case 0x03:
		return fmt.Sprintf("global[%v] -> \"%v\"", export.ExportDesc.GetIdxValue(), export.Name)
	default:
		return ""
	}
}

func newExport(reader io.Reader) (*Export, error) {
	var err error
	export := new(Export)
	export.Name, err = values2.NewName(reader) // Name
	if err != nil {
		return nil, err
	}

	export.ExportDesc, err = newExportDesc(reader)
	if err != nil {
		return nil, err
	}
	return export, nil
}

type ExportSection struct {
	Id           byte
	Size         uint32
	Exports      []*Export
	StartSection uint32
	StartContent uint32
}

// Parses an export section starting with the Size.
// The section Id is parsed before and will be automatically set.
func NewExportSection(reader io.Reader) (*ExportSection, error) {
	var err error
	exportSection := new(ExportSection)
	exportSection.Id = 0x07
	exportSection.StartSection = values2.GetByteCounter()

	exportSection.Size, err = values2.ReadU32(reader)
	if err != nil {
		return nil, errors.New(err.Error() + " ==> section size could not be determined")
	}

	exportSection.StartContent = values2.GetPositionOfNextByte()

	sectionReader, err := readNextSection(reader, exportSection.Size)
	if err != nil {
		values2.UpdateByteCounter(exportSection.StartContent - 1 + exportSection.Size)
		return nil, err
	}

	vecLen, err := values2.ReadU32(sectionReader)
	if err != nil {
		values2.UpdateByteCounter(exportSection.StartContent - 1 + exportSection.Size)
		return nil, err
	}

	log.Printf("Identified export section: Id = %x, start = %x, Size = %v, vecL = %v. Parsing...\n",
		exportSection.Id,
		exportSection.StartSection,
		exportSection.Size,
		vecLen)

	var i uint32 = 1
	for ; i <= vecLen; i++ {
		export, err := newExport(sectionReader)
		if err != nil {
			values2.UpdateByteCounter(exportSection.StartContent - 1 + exportSection.Size)
			return nil, err
		}
		exportSection.Exports = append(exportSection.Exports, export)
	}

	err = contentCompletelyRead("export", exportSection.Size, exportSection.StartContent, values2.GetByteCounter())
	if err != nil {
		values2.UpdateByteCounter(exportSection.StartContent - 1 + exportSection.Size)
		return nil, err
	}

	return exportSection, nil
}

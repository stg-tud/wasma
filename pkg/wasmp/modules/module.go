package modules

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	sections2 "wasma/pkg/wasmp/modules/sections"
	"wasma/pkg/wasmp/values"
)

type Module struct {
	customSectionsAll      []*sections2.CustomSection
	customSections         []*sections2.CustomSection
	typeSection            *sections2.TypeSection
	customSectionsType     []*sections2.CustomSection
	importSection          *sections2.ImportSection
	customSectionsImport   []*sections2.CustomSection
	functionSection        *sections2.FunctionSection
	customSectionsFunction []*sections2.CustomSection
	tableSection           *sections2.TableSection
	customSectionsTable    []*sections2.CustomSection
	memorySection          *sections2.MemorySection
	customSectionsMemory   []*sections2.CustomSection
	globalSection          *sections2.GlobalSection
	customSectionsGlobal   []*sections2.CustomSection
	exportSection          *sections2.ExportSection
	customSectionsExport   []*sections2.CustomSection
	startSection           *sections2.StartSection
	customSectionsStart    []*sections2.CustomSection
	elementSection         *sections2.ElementSection
	customSectionsElement  []*sections2.CustomSection
	codeSection            *sections2.CodeSection
	customSectionsCode     []*sections2.CustomSection
	dataSection            *sections2.DataSection
	customSectionsData     []*sections2.CustomSection
	nameSection            *sections2.CustomSection
}

func (module *Module) GetCustomSectionsAll() ([]*sections2.CustomSection, error) {
	if module.customSectionsAll != nil {
		return module.customSectionsAll, nil
	} else {
		return nil, errors.New("custom sections do not exist")
	}
}

func (module *Module) GetCustomSections() ([]*sections2.CustomSection, error) {
	if module.customSections != nil {
		return module.customSections, nil
	} else {
		return nil, errors.New("custom sections do not exist")
	}
}

func (module *Module) GetCustomSectionsType() ([]*sections2.CustomSection, error) {
	if module.customSectionsType != nil {
		return module.customSectionsType, nil
	} else {
		return nil, errors.New("custom sections after type section do not exist")
	}
}

func (module *Module) GetCustomSectionsImport() ([]*sections2.CustomSection, error) {
	if module.customSectionsImport != nil {
		return module.customSectionsImport, nil
	} else {
		return nil, errors.New("custom sections after import section do not exist")
	}
}

func (module *Module) GetCustomSectionsFunction() ([]*sections2.CustomSection, error) {
	if module.customSectionsFunction != nil {
		return module.customSectionsFunction, nil
	} else {
		return nil, errors.New("custom sections after function section do not exist")
	}
}

func (module *Module) GetCustomSectionsTable() ([]*sections2.CustomSection, error) {
	if module.customSectionsTable != nil {
		return module.customSectionsTable, nil
	} else {
		return nil, errors.New("custom sections after table section do not exist")
	}
}

func (module *Module) GetCustomSectionsMemory() ([]*sections2.CustomSection, error) {
	if module.customSectionsMemory != nil {
		return module.customSectionsMemory, nil
	} else {
		return nil, errors.New("custom sections after memory section do not exist")
	}
}

func (module *Module) GetCustomSectionsGlobal() ([]*sections2.CustomSection, error) {
	if module.customSectionsGlobal != nil {
		return module.customSectionsGlobal, nil
	} else {
		return nil, errors.New("custom sections after global section do not exist")
	}
}

func (module *Module) GetCustomSectionsExport() ([]*sections2.CustomSection, error) {
	if module.customSectionsExport != nil {
		return module.customSectionsExport, nil
	} else {
		return nil, errors.New("custom sections after export section do not exist")
	}
}

func (module *Module) GetCustomSectionsStart() ([]*sections2.CustomSection, error) {
	if module.customSectionsStart != nil {
		return module.customSectionsStart, nil
	} else {
		return nil, errors.New("custom sections after start section do not exist")
	}
}

func (module *Module) GetCustomSectionsElement() ([]*sections2.CustomSection, error) {
	if module.customSectionsElement != nil {
		return module.customSectionsElement, nil
	} else {
		return nil, errors.New("custom sections after element section do not exist")
	}
}

func (module *Module) GetCustomSectionsCode() ([]*sections2.CustomSection, error) {
	if module.customSectionsCode != nil {
		return module.customSectionsCode, nil
	} else {
		return nil, errors.New("custom sections after code section do not exist")
	}
}

func (module *Module) GetCustomSectionsData() ([]*sections2.CustomSection, error) {
	if module.customSectionsData != nil {
		return module.customSectionsData, nil
	} else {
		return nil, errors.New("custom sections after data section do not exist")
	}
}

func (module *Module) GetTypeSection() (*sections2.TypeSection, error) {
	if module.typeSection != nil {
		return module.typeSection, nil
	} else {
		return nil, errors.New("type section does not exist")
	}
}

func (module *Module) GetImportSection() (*sections2.ImportSection, error) {
	if module.importSection != nil {
		return module.importSection, nil
	} else {
		return nil, errors.New("import section does not exist")
	}
}

func (module *Module) GetFunctionSection() (*sections2.FunctionSection, error) {
	if module.functionSection != nil {
		return module.functionSection, nil
	} else {
		return nil, errors.New("function section does not exist")
	}
}

func (module *Module) GetTableSection() (*sections2.TableSection, error) {
	if module.tableSection != nil {
		return module.tableSection, nil
	} else {
		return nil, errors.New("table section does not exist")
	}
}

func (module *Module) GetMemorySection() (*sections2.MemorySection, error) {
	if module.memorySection != nil {
		return module.memorySection, nil
	} else {
		return nil, errors.New("memory section does not exist")
	}
}

func (module *Module) GetGlobalSection() (*sections2.GlobalSection, error) {
	if module.globalSection != nil {
		return module.globalSection, nil
	} else {
		return nil, errors.New("global section does not exist")
	}
}

func (module *Module) GetExportSection() (*sections2.ExportSection, error) {
	if module.exportSection != nil {
		return module.exportSection, nil
	} else {
		return nil, errors.New("export section does not exist")
	}
}

func (module *Module) GetStartSection() (*sections2.StartSection, error) {
	if module.startSection != nil {
		return module.startSection, nil
	} else {
		return nil, errors.New("start section does not exist")
	}
}

func (module *Module) GetElementSection() (*sections2.ElementSection, error) {
	if module.elementSection != nil {
		return module.elementSection, nil
	} else {
		return nil, errors.New("element section does not exist")
	}
}

func (module *Module) GetCodeSection() (*sections2.CodeSection, error) {
	if module.codeSection != nil {
		return module.codeSection, nil
	} else {
		return nil, errors.New("code section does not exist")
	}
}

func (module *Module) GetDataSection() (*sections2.DataSection, error) {
	if module.dataSection != nil {
		return module.dataSection, nil
	} else {
		return nil, errors.New("data section does not exist")
	}
}

func (module *Module) GetNameSection() (*sections2.CustomSection, error) {
	if module.nameSection != nil {
		return module.nameSection, nil
	} else {
		return nil, errors.New("name section does not exist")
	}
}

func NewModule(reader io.Reader) (*Module, error) {
	// The counters must be reset for each new module.
	values.ResetByteCounter()
	sections2.ResetFunctionIndex()
	sections2.ResetMemoryIndex()
	sections2.ResetTableIndex()
	sections2.ResetGlobalIndex()
	values.ResetOpcodeCounter()

	magic, err := values.ReadNextBytes(reader, 4)
	if err != nil {
		return nil, err
	}
	log.Printf("magic: %x\n", magic)
	if hex.EncodeToString(magic) != "0061736d" {
		return nil, errors.New(fmt.Sprintf("expected magic 0061736d but got: %x", magic))
	}

	version, err := values.ReadNextBytes(reader, 4)
	if err != nil {
		return nil, err
	}
	log.Printf("version: %x\n", version)
	if hex.EncodeToString(version) != "01000000" {
		return nil, errors.New(fmt.Sprintf("expected version 01000000 but got: %x", magic))
	}

	log.Println("reading wasm module")
	module := new(Module)
	var lastSection = 0x00

	for {
		sectionId, err := values.ReadNextByte(reader)
		if err != nil && err.Error() == "EOF" {
			return module, nil
		} else if err != nil {
			return nil, err
		} else {
			switch sectionId {
			case 0x00:
				newCustomSection, err := sections2.NewCustomSection(reader)
				if err != nil {
					return nil, err
				}
				module.customSectionsAll = append(module.customSectionsAll, newCustomSection)
				switch lastSection {
				case 0x00:
					module.customSections = append(module.customSections, newCustomSection)
				case 0x01:
					module.customSectionsType = append(module.customSectionsType, newCustomSection)
				case 0x02:
					module.customSectionsImport = append(module.customSectionsImport, newCustomSection)
				case 0x03:
					module.customSectionsFunction = append(module.customSectionsFunction, newCustomSection)
				case 0x04:
					module.customSectionsTable = append(module.customSectionsTable, newCustomSection)
				case 0x05:
					module.customSectionsMemory = append(module.customSectionsMemory, newCustomSection)
				case 0x06:
					module.customSectionsGlobal = append(module.customSectionsGlobal, newCustomSection)
				case 0x07:
					module.customSectionsExport = append(module.customSectionsExport, newCustomSection)
				case 0x08:
					module.customSectionsStart = append(module.customSectionsStart, newCustomSection)
				case 0x09:
					module.customSectionsElement = append(module.customSectionsElement, newCustomSection)
				case 0x0A:
					module.customSectionsCode = append(module.customSectionsCode, newCustomSection)
				case 0x0B:
					module.customSectionsData = append(module.customSectionsData, newCustomSection)
				}
				if newCustomSection.Name == "name" && lastSection == 0x0B {
					module.nameSection = newCustomSection
				}
			case 0x01:
				module.typeSection, err = sections2.NewTypeSection(reader)
				lastSection = 0x01
			case 0x02:
				module.importSection, err = sections2.NewImportSection(reader)
				lastSection = 0x02
			case 0x03:
				module.functionSection, err = sections2.NewFunctionSection(reader)
				lastSection = 0x03
			case 0x04:
				module.tableSection, err = sections2.NewTableSection(reader)
				lastSection = 0x04
			case 0x05:
				module.memorySection, err = sections2.NewMemorySection(reader)
				lastSection = 0x05
			case 0x06:
				module.globalSection, err = sections2.NewGlobalSection(reader)
				lastSection = 0x06
			case 0x07:
				module.exportSection, err = sections2.NewExportSection(reader)
				lastSection = 0x07
			case 0x08:
				module.startSection, err = sections2.NewStartSection(reader)
				lastSection = 0x08
			case 0x09:
				module.elementSection, err = sections2.NewElementSection(reader)
				lastSection = 0x09
			case 0x0A:
				values.OpCounterOn()
				module.codeSection, err = sections2.NewCodeSection(reader)
				values.OpCounterOff()
				lastSection = 0x0A
			case 0x0B:
				module.dataSection, err = sections2.NewDataSection(reader)
				lastSection = 0x0B
			default:
				return nil, errors.New(fmt.Sprintf("no valid section id found, got: %x", sectionId))
			}
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

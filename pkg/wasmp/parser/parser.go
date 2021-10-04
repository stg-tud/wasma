package parser

import (
	"bytes"
	"io/ioutil"
	"log"
	"wasma/pkg/wasmp/modules"
)

func Parse(file string) (*modules.Module, error) {
	inputBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	log.Printf("Start parsing: %s", file)

	module, err := modules.NewModule(bytes.NewReader(inputBytes))
	if err != nil {
		return nil, err
	}

	log.Printf("Parsing complete for: %s", file)

	return module, nil
}

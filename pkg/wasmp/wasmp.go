package main

import (
	"flag"
	"log"
	"wasma/pkg/wasmp/parser"
)

func main() {
	file := flag.String("file", "", "wasm file that should be analyzed")

	flag.Parse()

	if *file == "" {
		log.Fatal("The parameter file is mandatory.")
	}

	_, err := parser.Parse(*file)
	if err != nil {
		log.Fatal(err.Error())
	}
}

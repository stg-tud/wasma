package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"wasma/pkg/wasma"
	"wasma/pkg/wasma/graphs/controlFlowGraph"
	"wasma/pkg/wasma/output"
	"wasma/pkg/wasmp/modules"
)

type MinerAnalysis struct{}

type NGram struct {
	n      int
	items  []string
	nGrams *map[string]int
}

func (nGram *NGram) getNGram() string {
	nGramString := ""
	first := true
	for _, item := range nGram.items {
		if first {
			nGramString = nGramString + item
			first = false
		}
		nGramString = nGramString + " " + item
	}
	return nGramString
}

func (nGram *NGram) addNGram(nGramString string) {
	if value, found := (*nGram.nGrams)[nGramString]; found {
		value = value + 1
		(*nGram.nGrams)[nGramString] = value
	} else {
		(*nGram.nGrams)[nGramString] = 1
	}
}

func (nGram *NGram) addIntr(item string) {
	if len(nGram.items) >= nGram.n {
		nGram.addNGram(nGram.getNGram())
		for i := 0; i < nGram.n-1; i++ {
			nGram.items[i] = nGram.items[i+1]
		}
		nGram.items[nGram.n-1] = item
	} else {
		nGram.items = append(nGram.items, item)
	}
}

func NewNGram(n int, nGrams *map[string]int) NGram {
	return NGram{n, []string{}, nGrams}
}

func walkCFG(node uint32, subCFG *controlFlowGraph.CFG, nGram NGram, visited map[uint32]bool) {
	currentNode := subCFG.Tree[node]
	if currentNode.Control {
		nGram.addIntr(currentNode.Instruction.Name())
	} else {
		var instructions []uint32
		for key, _ := range currentNode.Block {
			instructions = append(instructions, key)
		}
		sort.Slice(instructions, func(i, j int) bool { return instructions[i] < instructions[j] })

		for _, instruction := range instructions {
			nGram.addIntr(currentNode.Block[instruction].Name())
		}
	}

	if _, found := visited[node]; found {
		return
	} else {
		visited[node] = true
	}

	for _, successor := range currentNode.Successors {
		walkCFG(successor.TargetNode, subCFG, nGram, visited)
	}
}

type ResultRecord struct {
	nGram    string
	numNGram int
	percent  float64
}

func (minerAnalysis *MinerAnalysis) Analyze(module *modules.Module, args map[string]string) {
	configFile, err := os.Open(args["con"])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	scanner := bufio.NewScanner(configFile)
	i := 1
	signature := make(map[string]bool)
	println("Miner Signature")
	for scanner.Scan() {
		fmt.Printf("%v: %v\n", i, scanner.Text())
		signature[scanner.Text()] = true
		i++
	}

	outputFile, err := output.OpenOrCreateTXT(args["out"])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	cfg := controlFlowGraph.NewControlFlowGraph(module, true, 0)
	fiveGrams := make(map[string]int)
	nGram := NewNGram(5, &fiveGrams)

	i = 1
	for _, subCfg := range cfg {
		//fmt.Printf("SubCFG: %v\n", i)
		walkCFG(0, subCfg, nGram, make(map[uint32]bool))
		i++
	}

	numNGrams := 0
	for _, count := range fiveGrams {
		numNGrams = numNGrams + count
	}

	var fiveGramResults []ResultRecord
	for fiveGram, count := range fiveGrams {
		percent := 0.0
		if numNGrams > 0 {
			percent = float64(count) / float64(numNGrams) * 100
		}
		fiveGramResults = append(fiveGramResults, ResultRecord{fiveGram, count, percent})
	}

	sort.Slice(fiveGramResults, func(i, j int) bool { return fiveGramResults[i].numNGram > fiveGramResults[j].numNGram })

	counter := 0
	for _, fiveGramResult := range fiveGramResults {
		if found := signature[fiveGramResult.nGram]; found {
			counter++
		}
	}

	if counter >= 8 {
		outputFile.WriteString(fmt.Sprintf("FOUND MINER: %v\n", args["file"]))
	} else {
		outputFile.WriteString(fmt.Sprintf("NO MINER: %v\n", args["file"]))
	}
}

func (minerAnalysis *MinerAnalysis) Name() string {
	return "Miner Detection Analysis"
}

func main() {
	analysis := wasma.NewWasmA()
	analysis.Start(&MinerAnalysis{})
}

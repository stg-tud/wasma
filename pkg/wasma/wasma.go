package wasma

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"wasma/pkg/wasmp/modules"
	"wasma/pkg/wasmp/parser"
)

type Analysis interface {
	Analyze(module *modules.Module, args map[string]string)
	Name() string
}

type WasmA struct {
	analyses   []Analysis
	inputFiles []string
	output     string
	con        string
	//analysis              int
	funcIdx               int
	funcName              string
	funcParams            string
	logFilePath           string
	considerIndirectCalls string
}

func (wasma *WasmA) Start(analysis Analysis) {
	log.Println("Start analysis")

	for _, wasmFile := range wasma.inputFiles {
		// parse wasm file
		if !strings.HasPrefix(wasmFile, "#") {
			module, err := parser.Parse(wasmFile)
			if err != nil {
				log.Fatal(err.Error())
			}

			args := map[string]string{
				"file": wasmFile,
				"out":  wasma.output,
				"con":  wasma.con,
				"fi":   strconv.Itoa(wasma.funcIdx),
				"fn":   wasma.funcName,
				"fp":   wasma.funcParams,
				"ic":   wasma.considerIndirectCalls,
			}

			// execute analysis
			analysis.Analyze(module, args)
		}
	}
}

var outputFile *os.File

// Analyses List of all available analyses.
var Analyses []Analysis

// AddAnalysis Add new analysis to the list of available analyses.
func AddAnalysis(analysis Analysis) {
	Analyses = append(Analyses, analysis)
}

func NewWasmA() WasmA {
	file := flag.String("file", "", "wasm file that should be analyzed")
	out := flag.String("out", "", "output path")
	con := flag.String("con", "", "path to config file")
	files := flag.String("files", "", "file containing a list of wasm files")
	funcIdx := flag.Int("fi", -1, "select a function by its function index")
	funcName := flag.String("fn", "", "select a function by its function name")
	funcParams := flag.String("fp", "", "select a list of parameters by its argument position seperatet by comma (e.g. 0,3,4)")
	logFilePath := flag.String("log", "", "log file")
	findIndirectCalls := flag.String("ic", "true", "if the flag is true indirect calls a considered during the analysis otherwise not (default: true)")
	//analysesList := flag.Bool("list", false, "if true a list of all available analyses is shown.")
	//analysis := flag.Int("analysis", 0, "select an analysis by its index.")
	flag.Parse()

	//if len(Analyses) <= *analysis {
	//	fmt.Printf("Analysis with index %v does not exist. Display a list with available analyses using the parameter '-list'\n", *analysis)
	//}
	//
	//if *analysesList {
	//	for index, analysis := range Analyses {
	//		log.Fatalf("(%v) %v\n", index, analysis.Name())
	//	}
	//}

	if *file == "" && *files == "" {
		log.Fatal("One of the two parameters '-file' or '-files' is mandatory.")
	}

	if *file != "" && *files != "" {
		log.Fatal("Select only one of the two parameters '-file' or '-files'.")
	}

	if *funcIdx != -1 && *funcName != "" {
		log.Fatal("Select only one of the two parameters '-fi' or '-fn'.")
	}

	if *funcIdx < -1 {
		log.Fatal("Function index must be greater or equal zero or -1 if no function should be selected.")
	}

	var wasmFiles []string

	if *files != "" {
		wasmFiles = readFileList(*files)
	} else {
		wasmFiles = append(wasmFiles, *file)
	}

	// logging
	if *logFilePath != "" {
		logFile, err := os.OpenFile(*logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error when opening log file: %v", err)
		}
		defer logFile.Close()

		log.SetOutput(logFile)
		log.Printf("log file set to: %v\n", *logFile)
	}

	return WasmA{Analyses, wasmFiles, *out, *con, *funcIdx, *funcName, *funcParams, *logFilePath, *findIndirectCalls}

}

func readFileList(fileList string) []string {
	var wasmFiles []string
	file, err := os.Open(fileList)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wasmFiles = append(wasmFiles, scanner.Text())
	}
	return wasmFiles
}

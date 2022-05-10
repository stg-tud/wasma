# WasmA: A Static Analysis Framework for WebAssembly

WasmA is a framework to create static analyses for WebAssembly binaries. Additionally, WasmA
provides several standard analyses.

### Provided Standrad Analyses

- CallGraph: Generates a call graph for a given binary and saves it as a dot file.
- ControlFlowGraph: Generates a control-flow graph for a given binary and saves it as a dot file.
- DataFlowGraph: Generates a data-flow graph for a given binary and saves it as a dot file.
- Data: Reads the data section of a binary and saves it as a csv file.
- Disassembly: Outputs a disassembly of a binary's function that can be specified by its function index.
- Exports: Outputs the exports of a given binary.
- Imports: Outputs the imports of a given binary.
- InstructionCount: Counts the occurrence of instructions and saves the results as txt file.
- OverapproximationBrTable: Determines the overapproximation for br_table instructions and saves the results as csv file.
- OverapproximationCallIndirect: Determines the overapproximation for call_indirect instructions and saves the results as csv file.
- SectionDetails: Saves the details of a binary's sections to a txt file.
- SectionList: Saves a list of a binary's sections to a txt file.

## Installing WasmA

There are two options to use the WasmA framework. The first one is to use the framework as a local installation
and the second one is to use WasmA inside a docker container.

### Local Installation

The local installation has the following **dependencies**:

- [golang](https://golang.org/doc/install)
- make

#### Installation

1. Install dependencies
2. Clone this repository
3. Change into the *wasma* directory of the cloned repository
4. Execute the following command to build the standard analyses:


    make build


### Running WasmA as Docker Container

Dependencies:
- [docker](https://www.docker.com/)
- [golang](https://golang.org/doc/install)
- make

#### Installation

1. Install dependencies
2. Clone this repository
3. Change into the *wasma* directory of the cloned repository
4. Execute the following command to build the docker image and container:


    ./docker-build.sh

The build script creates four directories:

- **$HOME/wasma**: Working directory.
- **$HOME/wasma/data**: This directory can be used to provide data that should be analyzed.
- **$HOME/wasma/bin**: This directory is intended as import and export directory for the docker container.
New created analyses can be exported to the local host using the script *docker-save-app.sh*. The analysis applications
are exported to this directory *$HOME/wasma/bin*. The script *docker-load-app.sh* allows to load an existing analysis application
into the docker container.
- **$HOME/wasma/analyses**: This directory can be used to export and import source code for analyses by using
the scripts *docker-save-analysis.sh* and *docker-load-analysis.sh*.

## Using WasmA

### Local

#### Provided Standard Analyses

After the local installation the standard analyses in the directory */bin* can be used immediately.
A list of all provided standard analyses can be found [here](#provided-standrad-analyses).
A description of the standard analyses' parameters can be found [here](#parameter-of-a-wasma-analysis).

#### Creating a New Analysis

A new analysis can be created by using the two go packages *wasma* and *wasmp*. For an easy start use
the following template for your analysis:

    package main
    
    import (
        "fmt"
        "log"
        "path/filepath"
        "strings"
        "wasma/pkg/wasma"
        "wasma/pkg/wasma/graphs"
        "wasma/pkg/wasmp/modules"
    )

    type NEWANALYSIS struct{}

    func (nEWANALYSIS *NEWANALYSIS) Analyze(module *modules.Module, args map[string]string) {
        // add code
    }

    func (nEWANALYSIS *NEWANALYSIS) Name() string {
        return "NEWANALYSIS"
    }

    func main() {
        analysis := wasma.NewWasmA()
        analysis.Start(&NEWANALYSIS{})
    }

The first step for a new analysis is to define a new empty *struct*. This *struct* has to implement
the two functions *Analyze* and *Name*. The function *Analyze* includes the analysis code and the function *Name*
returns the name of the analysis.

Every analysis needs a *main* function. In this *main* function the *wasma* framework has to be instantiated
by:

    analysis := wasma.NewWasmA()

The analysis is started by passing an instance of the previously defined *struct* to the *Start* function:

    analysis.Start(&NEWANALYSIS{})

The new analysis then can be compiled to an executable application with:

    go build -o <output file> <source file>

### Docker

To start and use the docker container enter the following commands:

    docker start wasma
    docker exec -it wasma sh



#### Using standard analyses

To list all provided standard analyses enter:

    list.sh

#### Create your own analysis application

To create a new analysis enter:

    create.sh <ANALYSIS NAME>

## Parameter of a WasmA Analysis

Every analysis that has been created using WasmA has the following parameters:

- **-file**: Specifies a binary file that should be analyzed.
- **-files**: Specifies a txt file with a list of binary files that should be analyzed.
- **-out**: Specifies an output directory for the analysis results.
- **-fi**: Specifies the function index to select a specific function for an analysis.
- **-log**: Specifies the location for the log file.
- **-ic**: Flag to specify if indirect calls should be considered.
    

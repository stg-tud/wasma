package output

import (
	"encoding/csv"
	"log"
	"os"
)

type CSVFile struct {
	filePath  string
	file      *os.File
	csvWriter *csv.Writer
}

func (csvFile *CSVFile) Write(record []string) {
	err := csvFile.csvWriter.Write(record)
	if err != nil {
		log.Println(err.Error())
	}
}

func (csvFile *CSVFile) WriteAll(record [][]string) {
	csvFile.csvWriter.WriteAll(record)
}

func (csvFile *CSVFile) Close() {
	csvFile.csvWriter.Flush()
	csvFile.file.Close()
}

func NewCSVFile(filePath string) (*CSVFile, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	} else {
		log.Printf("created file: %v\n", file)
	}

	return &CSVFile{filePath, file, csv.NewWriter(file)}, nil
}

func OpenOrCreateCSV(filePath string) (*CSVFile, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err == nil {
		return &CSVFile{filePath, file, csv.NewWriter(file)}, nil
	}

	file, err = os.Create(filePath)
	if err != nil {
		return nil, err
	} else {
		log.Printf("created file: %v\n", file)
	}

	return &CSVFile{filePath, file, csv.NewWriter(file)}, nil
}

func OpenOrCreateCSVHead(filePath string, head []string) (*CSVFile, error) {
	csvFile, err := OpenOrCreateCSV(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(csvFile.filePath)
	if err == nil && len(bytes) == 0 {
		csvFile.Write(head)
	}

	return csvFile, nil
}

func OpenOrCreateTXT(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err == nil {
		return file, nil
	}

	file, err = os.Create(filePath)
	if err != nil {
		return nil, err
	} else {
		log.Printf("created file: %v\n", file)
	}

	return file, nil
}

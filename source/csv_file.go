package source

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvFile struct {
	path string
}

func NewCsvFile(path string) *CsvFile {
	return &CsvFile{path: path}
}

func (csvFile CsvFile) Read() [][]string {
	f := openFileForReading(csvFile)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(f)

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

func openFileForReading(csvFile CsvFile) *os.File {
	f, err := os.Open(csvFile.path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return f
}

func (csvFile CsvFile) Write(data [][]string) {
	destination := createFileForWriting(csvFile)

	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(destination)

	writer := csv.NewWriter(destination)
	defer writer.Flush()

	for _, line := range data {
		err := writer.Write(line)
		if err != nil {
			fmt.Println("Error writing line to file")
		}
	}
}

func createFileForWriting(csvFile CsvFile) *os.File {
	destination, err := os.Create(csvFile.path)
	if err != nil {
		fmt.Println("Failed to open file")
	}
	return destination
}

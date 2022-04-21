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
	source, err := os.Open(csvFile.path)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(source)

	data, err := csv.NewReader(source).ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

func (csvFile CsvFile) Write(data [][]string) {
	destination, err := os.Create(csvFile.path)
	if err != nil {
		fmt.Println("Failed to open file")
	}

	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(destination)

	writer := csv.NewWriter(destination)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		fmt.Println("Error writing line to file")
	}

}

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

func (receiver CsvFile) Read() [][]string {
	f, err := os.Open(receiver.path)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

func (receiver CsvFile) Write(data [][]string) {
	destination, err := os.Create(receiver.path)
	defer destination.Close()

	if err != nil {
		fmt.Println("Failed to open file")
	}

	writer := csv.NewWriter(destination)
	defer writer.Flush()

	for _, line := range data {
		err := writer.Write(line)
		if err != nil {
			fmt.Println("Error writing line to file")
		}
	}
}

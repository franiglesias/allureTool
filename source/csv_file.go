package source

import (
	"encoding/csv"
	"log"
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
		log.Fatalln(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (receiver CsvFile) Write(data [][]string) {
	destination, err := os.Create(receiver.path)
	defer destination.Close()

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	writer := csv.NewWriter(destination)
	defer writer.Flush()

	for _, line := range data {
		err := writer.Write(line)
		if err != nil {
			log.Fatalln("Error writing line to file", err)
		}
	}
}

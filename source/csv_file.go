package source

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/afero"
)

type CsvFile struct {
	path string
	fs   afero.Fs
}

func NewCsvFile(path string) *CsvFile {
	return &CsvFile{path: path, fs: afero.NewOsFs()}
}

func NewCsvFileWithFS(path string, fs afero.Fs) *CsvFile {
	return &CsvFile{path: path, fs: fs}

}

func (csvFile CsvFile) Read() [][]string {
	source, err := csvFile.fs.Open(csvFile.path)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func(f afero.File) {
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
	destination, err := csvFile.fs.Create(csvFile.path)
	if err != nil {
		fmt.Println("Failed to open file")
	}

	defer func(destination afero.File) {
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

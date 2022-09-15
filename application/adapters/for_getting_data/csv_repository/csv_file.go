package csv_repository

import (
	"encoding/csv"
	"github.com/spf13/afero"
	"log"
)

type CSVFile struct {
	Fs   afero.Fs
	Path string
}

func (f CSVFile) Write(data [][]string) error {
	dest, err := f.Fs.Create(f.Path)
	if err != nil {
		return err
	}

	defer func(dest afero.File) {
		err := dest.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(dest)

	writer := csv.NewWriter(dest)
	defer writer.Flush()

	return writer.WriteAll(data)
}

func (f CSVFile) Read() [][]string {
	source, err := f.Fs.Open(f.Path)
	if err != nil {
		log.Fatalf("Cannot read file %s, %#v", f.Path, err)
	}
	defer func(source afero.File) {
		err := source.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(source)

	data, _ := csv.NewReader(source).ReadAll()
	return data
}

package csv_repository

import "github.com/spf13/afero"

type ProjectFile struct {
	project string
	path    string
}

func (f ProjectFile) ReadFromFs(fs afero.Fs) [][]string {
	csv := CSVFile{
		Fs:   fs,
		Path: f.path,
	}
	return csv.Read()
}

package csv_repository

import "github.com/spf13/afero"

type ProjectFile struct {
	Project string
	Path    string
}

func (f ProjectFile) ReadFromFs(fs afero.Fs) [][]string {
	csv := CSVFile{
		Fs:   fs,
		Path: f.Path,
	}
	return csv.Read()
}

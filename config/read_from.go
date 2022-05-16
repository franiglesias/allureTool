package config

import (
	"github.com/spf13/afero"
	"strings"
)

type DataFile struct {
	Path string
	Fs   afero.Fs
}

func NewDataFile(file string, fs afero.Fs) DataFile {
	return DataFile{
		Path: file,
		Fs:   fs,
	}
}

func (df DataFile) ReadLines() []string {
	bytes, err := afero.ReadFile(df.Fs, df.Path)
	if err != nil {
		return nil
	}
	return strings.Split(string(bytes), "\n")
}

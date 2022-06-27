package config

import (
	"github.com/spf13/afero"
	"os"
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
	bytes, _ := df.ReadBytes()
	return strings.Split(string(bytes), "\n")
}

func (df DataFile) ReadBytes() ([]byte, error) {
	bytes, err := afero.ReadFile(df.Fs, df.Path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (df DataFile) WriteBytes(bytes []byte) error {
	return afero.WriteFile(df.Fs, df.Path, bytes, os.ModePerm)
}

func (df DataFile) DuplicateTo(path string) (DataFile, error) {
	dest := NewDataFile(path, df.Fs)
	contents, err := df.ReadBytes()
	if err != nil {
		return dest, err
	}

	return dest, dest.WriteBytes(contents)
}

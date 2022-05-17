package config

import (
	"github.com/spf13/afero"
	"log"
	"strings"
)

type Directory struct {
	path string
	fs   afero.Fs
}

func (d Directory) Files() []string {
	var fPaths []string
	files, err := afero.ReadDir(d.fs, d.path)
	if err != nil {
		log.Fatal("Failed: "+d.path, err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		fPaths = append(fPaths, file.Name())
	}
	return fPaths
}

func FilesInDir(dir string) []string {
	directory := Directory{
		path: dir,
		fs:   afero.NewOsFs(),
	}

	return directory.Files()
}

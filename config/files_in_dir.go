package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func FilesInDir(dir string) []string {
	var fPaths []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("Failed: "+dir, err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		fmt.Println(file.Name())
		fPaths = append(fPaths, file.Name())
	}

	return fPaths
}

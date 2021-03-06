package main

import (
	. "allureTool/config"
	. "allureTool/report"
	. "allureTool/source"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	c := GetConfig()

	NewCsvFile(c.OutputFile()).
		Write(
			MakeEmptyReport().
				BuildWith(aggregatedDataIn(c.PathToReports())).
				FilterWith(filtersIn(c.FiltersFile())).
				ToRaw(),
		)
}

func aggregatedDataIn(folder string) [][]string {
	files := filesInDir(folder)
	var aggregated [][]string
	for _, file := range files {
		aggregated = append(aggregated, NewCsvFile(folder+file).Read()...)
	}
	return aggregated
}

func filesInDir(dir string) []string {
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

func filtersIn(filtersFile string) []string {
	var flat []string
	for _, line := range NewCsvFile(filtersFile).Read() {
		flat = append(flat, line[0])
	}
	return flat
}

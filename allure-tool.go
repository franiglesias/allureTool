package main

import (
	. "allureTool/config"
	. "allureTool/reader"
	. "allureTool/report"
	"io/ioutil"
	"log"
)

func main() {
	c := GetConfig()

	NewCsvFile(c.OutputFile()).
		Write(
			FilterReport(
				EmptyReport().BuildWith(aggregatedDataIn(c.PathToReports())),
				usingFiltersIn(c.FiltersFile()),
			).ToRaw(),
		)
}

func aggregatedDataIn(folder string) [][]string {
	files := filesInDir(folder)
	var raw [][]string
	for _, file := range files {
		rawContents := NewCsvFile(folder + file).Read()
		raw = append(raw, rawContents...)
	}
	return raw
}

func filesInDir(dir string) []string {
	var fPaths []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fPaths = append(fPaths, file.Name())
	}

	return fPaths
}

func usingFiltersIn(filtersFile string) []string {
	var flat []string
	for _, line := range NewCsvFile(filtersFile).Read() {
		flat = append(flat, line[0])
	}
	return flat
}

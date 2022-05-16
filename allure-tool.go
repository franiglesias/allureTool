package main

import (
	. "allureTool/config"
	. "allureTool/report"
	. "allureTool/source"
	"github.com/spf13/afero"
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
	files := FilesInDir(folder)
	var aggregated [][]string
	for _, file := range files {
		aggregated = append(aggregated, NewCsvFile(folder+file).Read()...)
	}
	return aggregated
}

func filtersIn(filtersFile string) []string {
	file := NewDataFile(filtersFile, afero.NewOsFs())

	return file.ReadLines()
}

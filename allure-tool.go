package main

import (
	. "allureTool/config"
	"allureTool/project"
	. "allureTool/report"
	. "allureTool/source"
	"github.com/spf13/afero"
)

func main() {
	c := getConfig()

	_ = project.GetProjects(c)

	NewCsvFile(c.OutputFile()).
		Write(
			MakeEmptyReport().
				BuildWith(aggregatedDataIn(c.PathToReports())).
				FilterWith(filtersIn(c.FiltersFile())).
				ToRaw(),
		)
}

func getConfig() Config {
	c := Config{
		Env:  ".env",
		Conf: "config.yml",
		Fs:   afero.NewOsFs(),
	}

	return c.Get()
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

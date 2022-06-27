package main

import (
	"allureTool/auto"
	. "allureTool/config"
	"allureTool/project"
	. "allureTool/report"
	. "allureTool/source"
	"github.com/spf13/afero"
	"time"
)

func main() {
	c := getConfig()

	_ = project.GetProjects(*c)

	NewCsvFile(c.OutputFile()).
		Write(
			MakeEmptyReport().
				BuildWith(aggregatedDataIn(c.PathToReports())).
				FilterWith(filtersIn(c.FiltersFile())).
				ToRaw(),
		)
}

func getConfig() *Config {
	fs := afero.NewOsFs()
	c := Config{
		Env:  ".env",
		Conf: "config.yml",
		Fs:   fs,
	}

	subfolder := time.Now().Format("2006-01-02-15-04-05")
	config, _ := auto.SelfProject(fs, subfolder, c.Get())
	return config
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

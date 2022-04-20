package config

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	output  string
	reports string
	baseDir string
	filters string
}

func GetConfig() Config {
	outputFile := flag.String("output", "output.csv", "File to generate results report")
	reports := flag.String("source", "allure", "Folder where report files are stored")
	filters := flag.String("filters", "filters.csv", "List of labels we want to find")
	baseDir := flag.String("base", "./data/", "Base folder for working")
	flag.Parse()

	conf := Config{
		output:  *outputFile,
		reports: *reports,
		baseDir: strings.TrimSuffix(*baseDir, string(os.PathSeparator)),
		filters: *filters,
	}
	return conf
}

func (c Config) PathToReports() string {
	return filepath.Join(c.baseDir, c.reports) + string(os.PathSeparator)
}

func (c Config) OutputFile() string {
	return filepath.Join(c.baseDir, c.output)
}

func (c Config) FiltersFile() string {
	return filepath.Join(c.baseDir, c.filters)
}

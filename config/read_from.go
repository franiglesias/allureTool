package config

import "allureTool/source"

func ReadFrom(file string) []string {
	var flat []string
	for _, line := range source.NewCsvFile(file).Read() {
		flat = append(flat, line[0])
	}
	return flat
}

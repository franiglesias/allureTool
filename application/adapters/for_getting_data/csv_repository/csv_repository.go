package csv_repository

import (
	"allureTool/application/domain"
	"fmt"
	"github.com/spf13/afero"
	"log"
	"strconv"
	"strings"
)

type CSVRepository struct {
	projects map[string]ProjectFile
	fs       afero.Fs
}

func MakeCSVRepositoryFromFiles(fs afero.Fs, projectFiles ...ProjectFile) CSVRepository {
	r := CSVRepository{
		fs:       fs,
		projects: map[string]ProjectFile{},
	}

	for _, p := range projectFiles {
		r.projects[p.project] = p
	}

	return r
}

func (r CSVRepository) Retrieve(name string) domain.ExecutionData {
	data := r.readProject(name)

	return domain.ExecutionData{
		Tests: convertRawDataToTests(data),
	}
}

func (r CSVRepository) readProject(name string) [][]string {
	return r.projects[name].ReadFromFs(r.fs)
}

func convertRawDataToTests(data [][]string) []domain.Test {
	var foundTests []domain.Test

	for _, datum := range data {
		if datum[0] == "epic" {
			continue
		}

		test := domain.Test{
			Epic:    normalizeLabel(datum[0]),
			Feature: normalizeLabel(datum[1]),
			Story:   normalizeLabel(datum[2]),
			Failed:  strToInt(datum[3]),
			Broken:  strToInt(datum[4]),
			Passed:  strToInt(datum[5]),
			Skipped: strToInt(datum[6]),
			Unknown: strToInt(datum[7]),
		}

		foundTests = append(foundTests, test)
	}
	return foundTests
}

func normalizeLabel(label string) string {
	if len(label) == 0 {
		return label
	}

	parts := strings.SplitN(label, ":", 2)

	if len(parts) != 2 {
		return label
	}

	return strings.ReplaceAll(parts[0], " ", "") + ":" + parts[1]
}

func strToInt(data string) int {
	converted, err := strconv.Atoi(data)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Conversion of %s to int failed", data), err)
	}
	return converted
}

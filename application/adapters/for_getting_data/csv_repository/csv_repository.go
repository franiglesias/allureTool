package csv_repository

import (
	"allureTool/application/domain"
	"github.com/spf13/afero"
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

		foundTests = append(foundTests, MakeTestFromRawData(datum))
	}
	return foundTests
}

package csv_repository

import (
	"allureTool/application/domain"
	"github.com/spf13/afero"
	"reflect"
	"testing"
)

func TestObtainDataFromCSVFilesRepository(t *testing.T) {
	project := ProjectFile{
		project: "myProject",
		path:    "path/behaviors.csv",
	}

	data := [][]string{
		{"epic", "feature", "story", "failed", "broken", "passed", "skipped", "unknown"},
		{"EP-002", "FT-003", "US-005", "0", "0", "1", "0", "0"},
	}

	fs := afero.NewMemMapFs()

	populateFileWithExampleData(fs, "path/behaviors.csv", data)

	r := MakeCSVRepositoryFromFiles(fs, project)

	got := r.Retrieve("myProject")

	if len(got.Tests) != 1 {
		t.Errorf("No data added to repository. Wanted: %v, got: %v", 1, got.Tests)
	}

	want := domain.Test{
		Epic:    "EP-002",
		Feature: "FT-003",
		Story:   "US-005",
		Failed:  0,
		Broken:  0,
		Passed:  1,
		Skipped: 0,
		Unknown: 0,
	}

	if !reflect.DeepEqual(got.Tests[0], want) {
		t.Errorf("Read or processing error. Wanted: %v, got: %v", want, got.Tests[0])

	}
}

func populateFileWithExampleData(fs afero.Fs, pathToFile string, data [][]string) {
	file := CSVFile{
		Fs:   fs,
		Path: pathToFile,
	}

	_ = file.Write(data)
}

package zip_repository

import (
	"allureTool/application/domain"
	"archive/zip"
	"encoding/csv"
	"github.com/spf13/afero"
	"reflect"
	"testing"
)

func TestObtainDataFromAllureZipArchive(t *testing.T) {
	data := [][]string{
		{"epic", "feature", "story", "failed", "broken", "passed", "skipped", "unknown"},
		{"EP-002", "FT-003", "US-005", "0", "0", "1", "0", "0"},
	}

	fs := afero.NewMemMapFs()

	makeSampleZipArchive(fs, "myProject.zip", data)

	zipArchive := MakeZipArchive(fs, "myProject.zip")

	r := MakeZipRepositoryFromArchive(zipArchive)

	got := r.Retrieve("myProject")

	if len(got.Tests) != 1 {
		t.Errorf("No data added to repository. Wanted: %v, got: %v", 1, got.Tests)
	}

	want := domain.MakePassedTest("EP-002", "FT-003", "US-005")

	if !reflect.DeepEqual(got.Tests[0], want) {
		t.Errorf("Read or processing error. Wanted: %v, got: %v", want, got.Tests[0])
	}
}

func makeSampleZipArchive(fs afero.Fs, file string, data [][]string) {
	archive, _ := fs.Create(file)
	defer archive.Close()

	writer := zip.NewWriter(archive)

	zipWriter, _ := writer.Create("data/behaviors.csv")
	csvWriter := csv.NewWriter(zipWriter)
	csvWriter.WriteAll(data)
	csvWriter.Flush()
	writer.Flush()
	writer.Close()
}

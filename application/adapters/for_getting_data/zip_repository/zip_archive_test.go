package zip_repository

import (
	"github.com/spf13/afero"
	"testing"
)

func TestZipArchiveDoesNotExist(t *testing.T) {
	archive := ZipArchive{
		Fs:   afero.NewMemMapFs(),
		Path: "/nonexistant.zip",
		Tmp:  "/tmp",
	}

	err := archive.ExtractFile("some_file.dat")

	if err == nil {
		t.Fatal("Error was expected, but not returned.")
	}
}

func TestZipArchiveIsNotValid(t *testing.T) {
	fs := afero.NewMemMapFs()
	_, _ = fs.Create("/not_valid.zip")

	archive := ZipArchive{
		Fs:   fs,
		Path: "/not_valid.zip",
		Tmp:  "/tmp",
	}

	err := archive.ExtractFile("some_file.dat")

	if err == nil {
		t.Fatal("Error was expected, but not returned.")
	}
}

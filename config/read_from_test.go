package config

import (
	"github.com/spf13/afero"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestReadFrom(t *testing.T) {

	fs := afero.NewMemMapFs()

	exampleFile, err := simulateFileInFS(fs)
	if err != nil {
		return
	}

	dataFile := DataFile{
		Path: exampleFile,
		Fs:   fs,
	}
	got := dataFile.ReadLines()

	want := dataForExample()

	if len(got) == 0 {
		t.Logf("No contents read from file %s", exampleFile)
		t.Fail()
	}
	for i, line := range want {
		if line != got[i] {
			t.Logf("Lines doesn't match at line %d. Expected %s, got %s", i+1, line, got[i])
			t.Fail()
		}
	}
}

func TestDuplicateTo(t *testing.T) {
	fs := afero.NewMemMapFs()

	exampleFile, err := simulateFileInFS(fs)
	if err != nil {
		return
	}

	dataFile := DataFile{
		Path: exampleFile,
		Fs:   fs,
	}

	duplicated, err := dataFile.DuplicateTo("another/path/copy.data")

	got := duplicated.ReadLines()
	want := dataFile.ReadLines()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Copy was wrong.")
	}
}

func simulateFileInFS(fs afero.Fs) (string, error) {
	var contents = []byte(strings.Join(dataForExample(), "\n"))

	exampleFile := "example.dat"

	err := afero.WriteFile(fs, exampleFile, contents, os.ModeAppend)
	if err != nil {
		return "", err
	}
	return exampleFile, nil
}

func dataForExample() []string {
	return []string{
		"line 1",
		"line 2",
		"line 3",
	}
}

package auto

import (
	"allureTool/config"
	"github.com/spf13/afero"
	"os"
	"reflect"
	"testing"
)

func TestPrepareANewProject(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "data/projects.csv", []byte("example"), os.ModeAppend)
	_ = afero.WriteFile(fs, "data/filters-epics.csv", []byte("US-150"), os.ModeAppend)

	expectedFs := []string{
		"data/subfolder/allure",
		"data/subfolder/filters-epics.csv",
		"data/subfolder/projects.csv",
	}

	conf := config.Config{
		Env:      "",
		Conf:     "",
		Fs:       fs,
		BaseDir:  "data",
		BaseUrl:  "",
		Server:   "",
		Password: "",
		Username: "",
		Reports:  "allure",
		Projects: "projects.csv",
		Filters:  "filters-epics.csv",
	}

	c, err := SelfProject(fs, "subfolder", conf)

	if err != nil {
		t.Errorf(err.Error())
	}

	files, err := afero.Glob(fs, "**/subfolder/*")
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(files) == 0 {
		t.Errorf("No files found")
	}

	if !reflect.DeepEqual(expectedFs, files) {
		t.Errorf("Expected %#v, got %#v", expectedFs, files)
	}

	verifyEquals(t, "data/subfolder", c.BaseDir)
	verifyEquals(t, "data/subfolder/allure/", c.PathToReports())
	verifyEquals(t, "data/subfolder/projects.csv", c.ProjectsFile())
	verifyEquals(t, "data/subfolder/filters-epics.csv", c.FiltersFile())
}

func verifyEquals(t *testing.T, want string, got string) {
	if got != want {
		t.Errorf("Expected base dir to be %s, got %s", want, got)
	}
}

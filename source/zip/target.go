package zip

import (
	"github.com/spf13/afero"
	"os"
	"path/filepath"
)

type Target struct {
	Project     string
	Destination string
	Prefix      string
	Extension   string
	Subfolder   string
}

func NewTarget(project, destination string) Target {
	return Target{
		Project:     project,
		Destination: destination,
		Prefix:      "allure-report",
		Extension:   ".csv",
		Subfolder:   "/data/behaviors",
	}
}

func (t Target) destinationFile() string {
	return t.Destination + t.Project + t.Extension
}

func (t Target) behaviorsFile() string {
	return t.Prefix + t.Subfolder + t.Extension
}

func (t Target) temporalBehaviorsFile() string {
	temporal, _ := filepath.Abs("/tmp")

	return temporal + string(os.PathSeparator) + t.behaviorsFile()
}

func (t Target) temporalFolder() string {
	temporal, _ := filepath.Abs("/tmp")

	return temporal + string(os.PathSeparator) + t.Prefix
}

func (t Target) IsNotNamed(name string) bool {
	return name != t.destinationFile()
}

func (t Target) MoveToDestinationInFs(fs afero.Fs) error {
	return fs.Rename(t.temporalBehaviorsFile(), t.destinationFile())
}

func (t Target) RemoveFromFs(fs afero.Fs) error {
	return fs.RemoveAll(t.temporalFolder())
}

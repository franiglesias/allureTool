package auto

import (
	"allureTool/config"
	"github.com/spf13/afero"
	"log"
	"os"
)

type MyFs afero.Fs

func SelfProject(fs afero.Fs, subfolder string, c config.Config) (*config.Config, error) {
	sf := "data" + string(os.PathSeparator) + subfolder

	createDirectory(fs, sf)
	createDirectory(fs, sf+string(os.PathSeparator)+c.Reports)

	f := config.NewDataFile(c.FiltersFile(), c.Fs)
	_, err := f.DuplicateTo(sf + string(os.PathSeparator) + c.Filters)
	if err != nil {
		return nil, err
	}

	p := config.NewDataFile(c.ProjectsFile(), c.Fs)
	_, err = p.DuplicateTo(sf + string(os.PathSeparator) + c.Projects)
	if err != nil {
		return nil, err
	}

	c.BaseDir = sf

	return &c, nil
}

func createDirectory(fs afero.Fs, path string) {
	_, err := fs.Stat(path)
	if err == nil {
		return
	}

	err = fs.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Panicf("I cannot create the directory %s because: %s", path, err.Error())
	}
}

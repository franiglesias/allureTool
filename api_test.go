package main

import (
	"allureTool/config"
	"allureTool/project"
	"github.com/spf13/afero"
	"testing"
)

func MakeConfigForTests() config.Config {
	c := config.Config{
		Env:  ".env",
		Conf: "config.yml",
		Fs:   afero.NewOsFs(),
	}

	return c.Get()
}

func TestApiGetProject(t *testing.T) {
	conf := MakeConfigForTests()

	projects := config.NewDataFile(conf.ProjectsFile(), conf.Fs).ReadLines()

	for _, name := range projects {
		p := project.Project{
			Name:   name,
			Config: conf,
		}
		err := p.GetData()
		if err != nil {
			t.Logf(err.Error())
			t.Fail()
		}
	}
}

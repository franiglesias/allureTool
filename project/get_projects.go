package project

import "allureTool/config"

func GetProjects(conf config.Config) error {
	projects := config.NewDataFile(conf.ProjectsFile(), conf.Fs).ReadLines()

	for _, name := range projects {
		p := Project{
			Name:   name,
			Config: conf,
		}
		err := p.GetData()
		if err != nil {
			return err
		}
	}

	return nil
}

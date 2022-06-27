package project

import "allureTool/config"

func GetProjects(conf config.Config) error {
	projectsFile := conf.ProjectsFile()
	projects := config.NewDataFile(projectsFile, conf.Fs).ReadLines()

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

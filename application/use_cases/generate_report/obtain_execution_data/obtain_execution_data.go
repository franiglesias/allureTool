package obtain_execution_data

import (
	"allureTool/application/domain"
	"allureTool/application/ports/for_getting_data"
)

type ObtainExecutionData struct {
	Repository for_getting_data.ProjectRepository
}

func MakeObtainExecutionData(repository for_getting_data.ProjectRepository) ObtainExecutionData {
	return ObtainExecutionData{
		Repository: repository,
	}
}

func (oed ObtainExecutionData) FromProjects(projects []string, filters []string) (domain.ExecutionData, error) {
	data := domain.MakeEmptyExecutionData()
	for _, project := range projects {
		for _, test := range oed.Repository.Retrieve(project).Tests {
			if test.Referencing(filters) {
				data.Append(test)
			}
		}
	}

	return data, nil
}

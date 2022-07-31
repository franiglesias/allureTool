package for_getting_data

import "allureTool/application/domain"

type ProjectRepository interface {
	Retrieve(name string) domain.ExecutionData
}

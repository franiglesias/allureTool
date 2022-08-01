package memory_repository

import "allureTool/application/domain"

type MemoryRepository struct {
	Data map[string]domain.ExecutionData
}

func MakeEmptyMemoryRepository() MemoryRepository {
	return MemoryRepository{
		Data: map[string]domain.ExecutionData{},
	}
}

func (r MemoryRepository) Retrieve(project string) domain.ExecutionData {
	return r.Data[project]
}

func (r *MemoryRepository) AddTest(project string, test domain.Test) {
	if _, ok := r.Data[project]; !ok {
		r.Data[project] = domain.ExecutionData{Tests: []domain.Test{}}
	}

	data := r.Data[project]
	data.Append(test)
	r.Data[project] = data
}

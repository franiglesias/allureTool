package obtain_execution_data

import (
	"allureTool/application/adapters/for_getting_data/memory_repository"
	"allureTool/application/domain"
	"reflect"
	"testing"
)

func TestObtainExecutionDataFromEmptyProject(t *testing.T) {
	projects := []string{"myProject"}

	emptyRepository := memory_repository.MemoryRepository{
		Data: map[string]domain.ExecutionData{"myProject": {}},
	}

	want := domain.ExecutionData{
		Tests: []domain.Test{},
	}

	gather := MakeObtainExecutionData(emptyRepository)

	got, err := gather.FromProjects(projects, []string{})
	if err != nil {
		t.Errorf("ObtainFromProjects() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ObtainFromProjects() got = %v, want %v", got, want)
	}
}

func TestObtainExecutionDataFromOneProject(t *testing.T) {
	projects := []string{"myProject"}

	repository := memory_repository.MemoryRepository{
		Data: map[string]domain.ExecutionData{
			"myProject": {
				Tests: []domain.Test{
					{
						Epic:    "epic",
						Feature: "feature",
						Story:   "US-001",
						Failed:  0,
						Broken:  0,
						Passed:  1,
						Skipped: 0,
						Unknown: 0,
					},
				},
			},
		},
	}

	want := domain.ExecutionData{
		Tests: []domain.Test{
			{
				Epic:    "epic",
				Feature: "feature",
				Story:   "US-001",
				Failed:  0,
				Broken:  0,
				Passed:  1,
				Skipped: 0,
				Unknown: 0,
			},
		},
	}

	gather := MakeObtainExecutionData(repository)

	got, err := gather.FromProjects(projects, []string{})
	if err != nil {
		t.Errorf("ObtainFromProjects() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ObtainFromProjects() got = %v, want %v", got, want)
	}
}

func TestObtainExecutionDataFromSeveralProjects(t *testing.T) {
	projects := []string{
		"myProject",
		"otherProject",
	}

	repository := memory_repository.MemoryRepository{
		Data: map[string]domain.ExecutionData{
			"myProject": {
				Tests: []domain.Test{
					{
						Epic:    "epic",
						Feature: "feature",
						Story:   "US-001",
						Failed:  0,
						Broken:  0,
						Passed:  1,
						Skipped: 0,
						Unknown: 0,
					},
				},
			},
			"otherProject": {
				Tests: []domain.Test{
					{
						Epic:    "EP-001",
						Feature: "FE-002",
						Story:   "US-003",
						Failed:  0,
						Broken:  0,
						Passed:  3,
						Skipped: 0,
						Unknown: 0,
					},
				},
			},
		},
	}

	want := domain.ExecutionData{
		Tests: []domain.Test{
			{
				Epic:    "epic",
				Feature: "feature",
				Story:   "US-001",
				Failed:  0,
				Broken:  0,
				Passed:  1,
				Skipped: 0,
				Unknown: 0,
			},
			{
				Epic:    "EP-001",
				Feature: "FE-002",
				Story:   "US-003",
				Failed:  0,
				Broken:  0,
				Passed:  3,
				Skipped: 0,
				Unknown: 0,
			},
		},
	}

	gather := MakeObtainExecutionData(repository)
	got, err := gather.FromProjects(projects, []string{})
	if err != nil {
		t.Errorf("ObtainFromProjects() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ObtainFromProjects() got = %v, want %v", got, want)
	}
}

func TestObtainExecutionDataFiltering(t *testing.T) {
	projects := []string{"myProject"}

	repository := memory_repository.MemoryRepository{
		Data: map[string]domain.ExecutionData{
			"myProject": {
				Tests: []domain.Test{
					{
						Epic:    "epic",
						Feature: "feature",
						Story:   "US-001",
						Failed:  0,
						Broken:  0,
						Passed:  1,
						Skipped: 0,
						Unknown: 0,
					},
				},
			},
		},
	}

	want := domain.ExecutionData{
		Tests: []domain.Test{
			{
				Epic:    "epic",
				Feature: "feature",
				Story:   "US-001",
				Failed:  0,
				Broken:  0,
				Passed:  1,
				Skipped: 0,
				Unknown: 0,
			},
		},
	}

	gather := MakeObtainExecutionData(repository)

	got, err := gather.FromProjects(projects, []string{"US-001"})
	if err != nil {
		t.Errorf("ObtainFromProjects() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ObtainFromProjects() got = %v, want %v", got, want)
	}
}

func TestObtainExecutionDataFilteringNotMatching(t *testing.T) {
	projects := []string{"myProject"}

	repository := memory_repository.MemoryRepository{
		Data: map[string]domain.ExecutionData{
			"myProject": {
				Tests: []domain.Test{
					{
						Epic:    "epic",
						Feature: "feature",
						Story:   "US-001",
						Failed:  0,
						Broken:  0,
						Passed:  1,
						Skipped: 0,
						Unknown: 0,
					},
				},
			},
		},
	}

	want := domain.ExecutionData{
		Tests: []domain.Test{},
	}

	gather := MakeObtainExecutionData(repository)

	got, err := gather.FromProjects(projects, []string{"US-002"})
	if err != nil {
		t.Errorf("ObtainFromProjects() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ObtainFromProjects() got = %v, want %v", got, want)
	}
}

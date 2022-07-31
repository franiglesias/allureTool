package generate_report

import (
	"allureTool/application/adapters/for_getting_data/memory_repository"
	"allureTool/application/domain"
	"allureTool/application/use_cases/generate_report/analyze_execution"
	"allureTool/application/use_cases/generate_report/obtain_execution_data"
	"allureTool/application/use_cases/generate_report/summarize_data"
	"reflect"
	"testing"
)

const projectName = "myProject"

type story struct {
	epic    string
	feature string
	story   string
}

func TestGenerateSuccessfulReportAllTrackedAreTested(t *testing.T) {
	stories := []story{
		{epic: "EP-001", feature: "FT-001", story: "US-001"},
		{epic: "EP-001", feature: "FT-002", story: "US-002"},
		{epic: "EP-002", feature: "FT-004", story: "US-003"},
	}

	repository := givenARepositoryWithData(stories)

	generateReport := GenerateReport{
		obtain:    obtain_execution_data.MakeObtainExecutionData(repository),
		analyze:   analyze_execution.AnalyzeExecution{},
		summarize: summarize_data.Summarize{},
	}

	got, err := generateReport.Execute(givenRequestForFilters("US-001", "US-002", "US-003"))

	if err != nil {
		t.Errorf("Execute raised an error %v", err.Error())
	}

	want := expectedReport(stories)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %#v, got %#v", want, got)
	}
}

func givenRequestForFilters(filters ...string) GenerateReportRequest {
	projects := []string{
		projectName,
	}

	return GenerateReportRequest{
		filters:  filters,
		projects: projects,
	}
}

func givenARepositoryWithData(stories []story) *memory_repository.MemoryRepository {
	repository := memory_repository.MakeEmptyMemoryRepository()
	for _, s := range stories {
		repository.AddTest(projectName, makePassingTest(s))
	}

	return &repository
}

func expectedReport(stories []story) GenerateReportResponse {
	summary := domain.ReportSummary{
		Tracked: len(stories),
		Found:   len(stories),
		Pct:     100.00,
	}

	details := domain.ReportDetails{
		Lines: []domain.DetailsLine{},
	}

	for _, s := range stories {
		details.Lines = append(details.Lines, makeDetailsName(s))
	}
	return GenerateReportResponse{
		summary: summary,
		details: details,
	}
}

func makePassingTest(s story) domain.Test {
	return domain.Test{
		Epic:    s.epic,
		Feature: s.feature,
		Story:   s.story,
		Failed:  0,
		Broken:  0,
		Passed:  1,
		Skipped: 0,
		Unknown: 0,
	}
}

func makeDetailsName(s story) domain.DetailsLine {
	return domain.DetailsLine{
		Epic:    s.epic,
		Feature: s.feature,
		Story:   s.story,
		Tested:  true,
	}
}

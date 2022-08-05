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

var FirstUserStory = story{epic: "EP-001", feature: "FT-001", story: "US-001"}
var SecondUserStory = story{epic: "EP-001", feature: "FT-002", story: "US-002"}
var ThirdUserStory = story{epic: "EP-002", feature: "FT-004", story: "US-003"}

func TestGenerateSuccessfulReportAllTrackedAreTested(t *testing.T) {
	stories := []story{FirstUserStory, SecondUserStory, ThirdUserStory}

	generateReport := buildGenerateReportWithStories(stories)

	got, _ := generateReport.Execute(GenerateReportRequest{
		filters:  []string{"US-001", "US-002", "US-003"},
		projects: []string{projectName},
	})

	want := expectedReport(stories)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %#v, got %#v", want, got)
	}
}

func TestGenerateSuccessfulReportFiltering(t *testing.T) {
	stories := []story{FirstUserStory, SecondUserStory, ThirdUserStory}

	generateReport := buildGenerateReportWithStories(stories)

	got, _ := generateReport.Execute(GenerateReportRequest{
		filters:  []string{"US-002"},
		projects: []string{projectName},
	})

	want := expectedReport([]story{SecondUserStory})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %#v, got %#v", want, got)
	}
}

func buildGenerateReportWithStories(stories []story) GenerateReport {
	repository := givenARepositoryWithData(stories)

	return GenerateReport{
		obtain:    obtain_execution_data.MakeObtainExecutionData(repository),
		analyze:   analyze_execution.AnalyzeExecution{},
		summarize: summarize_data.Summarize{},
	}
}

func givenARepositoryWithData(stories []story) *memory_repository.MemoryRepository {
	repository := memory_repository.MakeEmptyMemoryRepository()
	for _, s := range stories {
		repository.AddTest(projectName, domain.MakePassedTest(s.epic, s.feature, s.story))
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

func makeDetailsName(s story) domain.DetailsLine {
	return domain.DetailsLine{
		Epic:    s.epic,
		Feature: s.feature,
		Story:   s.story,
		Tested:  true,
	}
}

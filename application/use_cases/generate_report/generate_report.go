package generate_report

import (
	"allureTool/application/domain"
	"allureTool/application/use_cases/generate_report/analyze_execution"
	"allureTool/application/use_cases/generate_report/obtain_execution_data"
	"allureTool/application/use_cases/generate_report/summarize_data"
)

type GenerateReportRequest struct {
	filters  []string
	projects []string
}

type GenerateReportResponse struct {
	summary domain.ReportSummary
	details domain.ReportDetails
}

type GenerateReport struct {
	obtain    obtain_execution_data.ObtainExecutionData
	analyze   analyze_execution.AnalyzeExecution
	summarize summarize_data.Summarize
}

func (g GenerateReport) Execute(request GenerateReportRequest) (GenerateReportResponse, error) {
	testExecutionData, err := g.obtain.FromProjects(request.projects, request.filters)

	if err != nil {
		return GenerateReportResponse{}, err
	}

	details := g.analyze.ExecutionData(testExecutionData)

	return GenerateReportResponse{
		summary: g.summarize.Details(details),
		details: details,
	}, nil
}

package generate_report

import (
	"allureTool/application/domain"
	"allureTool/application/use_cases/generate_report/analyze_execution"
	"allureTool/application/use_cases/generate_report/obtain_execution_data"
	"allureTool/application/use_cases/generate_report/summarize_data"
)

type GenerateReportRequest struct {
	Filters  []string
	Projects []string
}

type GenerateReportResponse struct {
	Summary domain.ReportSummary
	Details domain.ReportDetails
}

type GenerateReport struct {
	Obtain    obtain_execution_data.ObtainExecutionData
	Analyze   analyze_execution.AnalyzeExecution
	Summarize summarize_data.Summarize
}

func (g GenerateReport) Execute(request GenerateReportRequest) (GenerateReportResponse, error) {
	testExecutionData, err := g.Obtain.FromProjects(request.Projects, request.Filters)

	if err != nil {
		return GenerateReportResponse{}, err
	}

	details := g.Analyze.ExecutionData(testExecutionData)

	return GenerateReportResponse{
		Details: details,
		Summary: g.Summarize.Details(details),
	}, nil
}

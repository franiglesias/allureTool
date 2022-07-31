package analyze_execution

import (
	"allureTool/application/domain"
)

type AnalyzeExecution struct {
}

func (AnalyzeExecution) ExecutionData(data domain.ExecutionData) domain.ReportDetails {
	report := domain.MakeEmptyReportDetails()

	for _, test := range data.Tests {
		report.AddTest(test)
	}

	return report
}

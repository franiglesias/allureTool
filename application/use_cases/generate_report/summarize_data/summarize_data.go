package summarize_data

import (
	"allureTool/application/domain"
)

type Summarize struct {
}

func (s Summarize) Details(details domain.ReportDetails) domain.ReportSummary {
	return details.Summarize()
}

package report

import "strings"

func FilterReport(report Report, filters []string) Report {
	filtered := EmptyReport()

	for _, test := range report.Tests {
		for _, filter := range filters {
			if test.IsRelatedTo(filter) {
				filtered = filtered.AddTest(test)
			}
		}
	}

	return filtered
}

func (test Test) IsRelatedTo(filter string) bool {
	return strings.HasPrefix(test.epic, filter) || strings.HasPrefix(test.feature, filter) || strings.HasPrefix(test.story, filter)
}

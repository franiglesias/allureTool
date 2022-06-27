package report

import "strings"

func (report Report) FilterWith(filters []string) Report {
	filtered := MakeEmptyReport()

	for _, test := range report.Tests {
		for _, filter := range filters {
			if filter == "" {
				continue
			}
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

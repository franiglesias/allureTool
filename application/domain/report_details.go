package domain

type DetailsLine struct {
	Symbol  string
	Project string
	Epic    string
	Feature string
	Story   string
	Tested  bool
}

func MakeDetailsLineOfTest(test Test) DetailsLine {
	return DetailsLine{
		Epic:    test.Epic,
		Feature: test.Feature,
		Story:   test.Story,
		Tested:  test.Tested(),
	}
}

type ReportDetails struct {
	Lines []DetailsLine
}

func MakeEmptyReportDetails() ReportDetails {
	return ReportDetails{
		Lines: []DetailsLine{},
	}
}

func (r ReportDetails) Summarize() ReportSummary {
	tracked := 0
	found := 0

	for _, detail := range r.Lines {
		tracked += 1
		if detail.Tested {
			found += 1
		}
	}

	return ReportSummary{
		Tracked: tracked,
		Found:   found,
		Pct:     float32(found/tracked) * 100,
	}
}

func (r *ReportDetails) AddTest(test Test) {
	line := MakeDetailsLineOfTest(test)

	r.Lines = append(r.Lines, line)
}

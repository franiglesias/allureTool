package summarize_data

import (
	domain2 "allureTool/application/domain"
	"reflect"
	"testing"
)

func TestSummarizeData(t *testing.T) {
	tests := []struct {
		name    string
		details domain2.ReportDetails
		want    domain2.ReportSummary
	}{
		{name: "All success", details: buildDetailsSuccess(), want: domain2.ReportSummary{
			Tracked: 3,
			Found:   3,
			Pct:     100.0,
		}},
		{name: "All failed", details: buildDetailsNoneTested(), want: domain2.ReportSummary{
			Tracked: 3,
			Found:   0,
			Pct:     0.0,
		}},
	}

	summarize := Summarize{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summarize.Details(tt.details); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Details() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildDetailsSuccess() domain2.ReportDetails {
	return domain2.ReportDetails{
		Lines: []domain2.DetailsLine{
			{
				Project: "myproject",
				Symbol:  "US-001",
				Epic:    "EP-001",
				Feature: "FT-001",
				Story:   "US-001",
				Tested:  true,
			},
			{
				Project: "myproject",
				Symbol:  "US-002",
				Epic:    "EP-001",
				Feature: "FT-002",
				Story:   "US-002",
				Tested:  true,
			},
			{
				Project: "myproject",
				Symbol:  "US-003",
				Epic:    "EP-002",
				Feature: "FT-004",
				Story:   "US-003",
				Tested:  true,
			},
		},
	}
}

func buildDetailsNoneTested() domain2.ReportDetails {
	return domain2.ReportDetails{
		Lines: []domain2.DetailsLine{
			{
				Project: "myproject",
				Symbol:  "US-001",
				Epic:    "EP-001",
				Feature: "FT-001",
				Story:   "US-001",
				Tested:  false,
			},
			{
				Project: "myproject",
				Symbol:  "US-002",
				Epic:    "EP-001",
				Feature: "FT-002",
				Story:   "US-002",
				Tested:  false,
			},
			{
				Project: "myproject",
				Symbol:  "US-003",
				Epic:    "EP-002",
				Feature: "FT-004",
				Story:   "US-003",
				Tested:  false,
			},
		},
	}
}

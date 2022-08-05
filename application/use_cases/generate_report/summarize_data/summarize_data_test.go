package summarize_data

import (
	"allureTool/application/domain"
	"reflect"
	"testing"
)

func TestSummarizeData(t *testing.T) {
	tests := []struct {
		name    string
		details domain.ReportDetails
		want    domain.ReportSummary
	}{
		{name: "All success", details: buildDetailsSuccess(), want: domain.ReportSummary{
			Tracked: 3,
			Found:   3,
			Pct:     100.0,
		}},
		{name: "All failed", details: buildDetailsNoneTested(), want: domain.ReportSummary{
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

func buildDetailsSuccess() domain.ReportDetails {
	return domain.ReportDetails{
		Lines: []domain.DetailsLine{
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

func buildDetailsNoneTested() domain.ReportDetails {
	return domain.ReportDetails{
		Lines: []domain.DetailsLine{
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

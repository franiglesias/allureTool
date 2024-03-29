package analyze_execution

import (
	"allureTool/application/domain"
	"reflect"
	"testing"
)

func TestAnalyzeTest_ExecutionData(t *testing.T) {
	type args struct {
		data domain.ExecutionData
	}
	tests := []struct {
		name string
		args args
		want domain.ReportDetails
	}{
		{
			name: "Test is added",
			args: args{
				data: domain.ExecutionData{
					Tests: []domain.Test{
						domain.MakePassedTest("EP-001", "FT-001", "US-001"),
					},
				},
			},
			want: domain.ReportDetails{Lines: []domain.DetailsLine{
				{
					Epic:    "EP-001",
					Feature: "FT-001",
					Story:   "US-001",
					Tested:  true,
				},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			an := AnalyzeExecution{}
			if got := an.ExecutionData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecutionData() = %v, want %v", got, tt.want)
			}
		})
	}
}

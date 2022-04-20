package report

import (
	"strings"
	"testing"
)

func TestBuildReport(t *testing.T) {
	t.Run("Converts raw data into a report", func(t *testing.T) {
		want := expectedTests()
		got := EmptyReport().BuildWith(rawData())

		if len(got.Tests) != 3 {
			t.Errorf("Got %d lines instead of %d", len(got.Tests), 3)
		}

		for i, line := range got.Tests {
			if line != want[i] {
				t.Errorf("Line %d failed to be converted", i+1)
			}
		}
	})

	t.Run("Converts Report to raw data", func(t *testing.T) {
		report := MakeReport(expectedTests())
		want := rawOutputData()

		got := report.ToRaw()

		if len(got) != 3 {
			t.Errorf("Got %d lines instead of %d", len(got), 3)
		}

		for i, line := range got {
			for j, cell := range line {
				if cell != want[i][j] {
					t.Errorf("Line %d failed to be converted", i+1)
				}
			}
		}
	})

	t.Run("Can filter epics", func(t *testing.T) {
		report := MakeReport(expectedTests())
		filters := []string{
			"US-454",
		}

		filtered := FilterReport(report, filters)

		if !strings.HasPrefix(filtered.Tests[0].epic, "US-454") {
			t.Errorf("This line should not be here: %s, expected: %s", filtered.Tests[0].epic, "US-454")
		}
	})

	t.Run("Can filter features", func(t *testing.T) {
		report := MakeReport(expectedTests())
		filters := []string{
			"US-35",
		}

		filtered := FilterReport(report, filters)

		if !strings.HasPrefix(filtered.Tests[0].feature, "US-35") {
			t.Errorf("This line should not be here: %s, expected: %s", filtered.Tests[0].feature, "US-35")
		}
	})

	t.Run("Can filter stories", func(t *testing.T) {
		report := MakeReport(expectedTests())
		filters := []string{
			"US-435",
		}

		filtered := FilterReport(report, filters)

		if !strings.HasPrefix(filtered.Tests[0].story, "US-435") {
			t.Errorf("This line should not be here: %s, expected: %s", filtered.Tests[0].story, "US-435")
		}
	})
}

func expectedTests() []Test {
	return []Test{
		{
			epic:    "US-114: This is an EPIC",
			feature: "US-183: Correctly store data",
			story:   "US-234: Some story",
			failed:  0,
			broken:  0,
			passed:  23,
			skipped: 0,
			unknown: 0,
		},
		{
			epic:    "US-114: This is an EPIC",
			feature: "US-35: Event store",
			story:   "US-435: Another story",
			failed:  0,
			broken:  0,
			passed:  56,
			skipped: 0,
			unknown: 0,
		},
		{
			epic:    "US-454: Another epic",
			feature: "US-880: Use tracking service",
			story:   "",
			failed:  0,
			broken:  0,
			passed:  2,
			skipped: 0,
			unknown: 0,
		},
	}
}

func rawData() [][]string {
	return [][]string{
		{"Epic", "Feature", "Story", "FAILED", "BROKEN", "PASSED", "SKIPPED", "UNKNOWN"},
		{"US-114: This is an EPIC", "US-183: Correctly store data", "US-234: Some story", "0", "0", "23", "0", "0"},
		{"US-114: This is an EPIC", "US-35: Event store", "US-435: Another story", "0", "0", "56", "0", "0"},
		{"US - 454: Another epic", "US-880: Use tracking service", "", "0", "0", "2", "0", "0"},
	}
}

func rawOutputData() [][]string {
	return [][]string{
		//{"Epic", "Feature", "Story", "FAILED", "BROKEN", "PASSED", "SKIPPED", "UNKNOWN"},
		{"US-114: This is an EPIC", "US-183: Correctly store data", "US-234: Some story", "0", "0", "23", "0", "0"},
		{"US-114: This is an EPIC", "US-35: Event store", "US-435: Another story", "0", "0", "56", "0", "0"},
		{"US-454: Another epic", "US-880: Use tracking service", "", "0", "0", "2", "0", "0"},
	}
}

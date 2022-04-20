package source

import (
	"os"
	"testing"
)

func TestFileReader(t *testing.T) {

	const csvFile = "../data/example.csv"

	cleanExampleFile(t, csvFile)

	NewCsvFile(csvFile).Write(theWantedThing())

	t.Run("Can read a CSV file given path", func(t *testing.T) {
		got := NewCsvFile(csvFile).Read()
		want := theWantedThing()

		failIfSomeCellDiverge(t, want, got)
	})
}

func cleanExampleFile(t *testing.T, csvFile string) {
	_, err := os.Stat(csvFile)
	if err == nil {
		err := os.Remove(csvFile)
		if err != nil {
			t.Errorf("Example file could not be cleaned")
		}
	}
}

func failIfSomeCellDiverge(t *testing.T, want [][]string, got [][]string) {
	for i, row := range want {
		for j := range row {
			if got[i][j] != want[i][j] {
				t.Errorf("Not the expected data at (%d, %d). Wanted: %s, got: %s", i, j, want[i][j], got[i][j])
			}
		}
	}
}

func theWantedThing() [][]string {
	return [][]string{
		{"header 1", "header 2"},
		{"line 1 1", "line 1 2"},
		{"line 2 1", "line 2 2"},
	}
}

package source

import (
	"github.com/spf13/afero"
	"testing"
)

var TestFS = afero.NewMemMapFs()

func TestCsvFileWrite(t *testing.T) {
	const csvFile = "example.csv"

	NewCsvFileWithFS(csvFile, TestFS).Write(theWantedThing())

	got, _ := afero.ReadFile(TestFS, csvFile)

	want := theExpected()

	if string(got) != string(want) {
		t.Errorf("Want: %s\nGot:  %s\n", want, got)
	}
}

func TestFileReader(t *testing.T) {

	const csvFile = "example.csv"

	NewCsvFileWithFS(csvFile, TestFS).Write(theWantedThing())

	got := NewCsvFileWithFS(csvFile, TestFS).Read()
	want := theWantedThing()

	failIfSomeCellDiverge(t, want, got)

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

func theExpected() []byte {
	var e []byte
	for _, lines := range theWantedThing() {
		e = append(e, []byte(lines[0]+","+lines[1]+"\n")...)
	}
	return e
}

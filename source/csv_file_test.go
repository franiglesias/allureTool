package source

import (
	"github.com/spf13/afero"
	"os"
	"testing"
)

var TestFS = afero.NewMemMapFs()

func TestCsvFileWrite(t *testing.T) {
	const csvFile = "example.csv"

	NewCsvFileWithFS(csvFile, TestFS).Write(theWantedThing())

	got, _ := afero.ReadFile(TestFS, csvFile)

	want := theExpected()

	if string(got) != string(want) {
		t.Logf("Want: %s\nGot:  %s\n", want, got)
		t.Fail()
	}
}

func TestFileReader(t *testing.T) {
	const csvFile = "example.csv"

	err := afero.WriteFile(TestFS, csvFile, theExpected(), os.ModePerm)
	if err != nil {
		return
	}
	got := NewCsvFileWithFS(csvFile, TestFS).Read()
	want := theWantedThing()

	failIfSomeCellDiverge(t, want, got)

}

func failIfSomeCellDiverge(t *testing.T, want [][]string, got [][]string) {
	for i, row := range want {
		for j := range row {
			if got[i][j] != want[i][j] {
				t.Logf("Not the expected data at (%d, %d). Wanted: %s, got: %s", i, j, want[i][j], got[i][j])
				t.Fail()
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

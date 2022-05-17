package config

import (
	"github.com/spf13/afero"
	"os"
	"strings"
	"testing"
)

func TestGetFilesInDir(t *testing.T) {
	tests := []struct {
		name   string
		folder string
		files  []string
		want   []string
	}{
		{
			name:   "Get all files",
			folder: "example",
			files:  []string{"first.csv", "second.csv", "third.csv"},
			want:   []string{"first.csv", "second.csv", "third.csv"},
		},
		{
			name:   "Ignore invisibles",
			folder: "example",
			files:  []string{".first.csv", "second.csv", "third.csv"},
			want:   []string{"second.csv", "third.csv"},
		},
		{
			name:   "Ignore parent dirs",
			folder: "example",
			files:  []string{".", "..", "third.csv"},
			want:   []string{"third.csv"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			directory := Directory{
				path: test.folder,
				fs:   afero.NewMemMapFs(),
			}

			for _, file := range test.files {
				_, _ = directory.fs.Create(test.folder + string(os.PathSeparator) + file)
			}

			got := directory.Files()
			if !equalSlices(got, test.want) {
				t.Logf("Expected [%s], got [%s]", strings.Join(test.want, ", "), strings.Join(got, ", "))
				t.Fail()
			}
		})
	}
}

func equalSlices(got []string, want []string) bool {
	if len(got) != len(want) {
		return false
	}

	for i, element := range got {
		if element != want[i] {
			return false
		}
	}

	return true
}

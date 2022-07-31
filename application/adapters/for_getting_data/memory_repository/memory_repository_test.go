package memory_repository

import (
	"allureTool/application/domain"
	"testing"
)

func TestAddingTestsToMemoryRepository(t *testing.T) {
	r := MakeEmptyMemoryRepository()

	test := domain.Test{
		Epic:    "EP-002",
		Feature: "FT-003",
		Story:   "US-005",
		Failed:  0,
		Broken:  0,
		Passed:  1,
		Skipped: 0,
		Unknown: 0,
	}

	r.AddTest("myProject", test)
	got := r.Retrieve("myProject")

	if len(got.Tests) != 1 {
		t.Errorf("No data added to repository. Wanted: %v, got: %v", 2, got.Tests)
	}
}

package memory_repository

import (
	"allureTool/application/domain"
	"testing"
)

func TestAddingTestsToMemoryRepository(t *testing.T) {
	r := MakeEmptyMemoryRepository()

	test := domain.MakePassedTest("EP-002", "FT-003", "US-005")

	r.AddTest("myProject", test)
	got := r.Retrieve("myProject")

	if len(got.Tests) != 1 {
		t.Errorf("No data added to repository. Wanted: %v, got: %v", 2, got.Tests)
	}
}

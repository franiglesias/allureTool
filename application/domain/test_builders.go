package domain

func MakePassedTest(epic, feature, story string) Test {
	return Test{
		Epic:    epic,
		Feature: feature,
		Story:   story,
		Failed:  0,
		Broken:  0,
		Passed:  1,
		Skipped: 0,
		Unknown: 0,
	}
}

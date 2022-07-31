package domain

type Test struct {
	Epic    string
	Feature string
	Story   string
	Failed  int
	Broken  int
	Passed  int
	Skipped int
	Unknown int
}

func (t Test) Tested() bool {
	return t.Passed > 0
}

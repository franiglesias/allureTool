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

func (t Test) Referencing(filters []string) bool {
	if len(filters) == 0 {
		return true
	}
	for _, filter := range filters {
		if t.Epic == filter || t.Feature == filter || t.Story == filter {
			return true
		}
	}

	return false
}

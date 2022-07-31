package domain

type ExecutionData struct {
	Tests []Test
}

func MakeEmptyExecutionData() ExecutionData {
	return ExecutionData{
		Tests: []Test{},
	}
}

func (d *ExecutionData) Append(test Test) {
	d.Tests = append(d.Tests, test)
}

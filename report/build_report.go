package report

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Test struct {
	epic    string
	feature string
	story   string
	failed  int
	broken  int
	passed  int
	skipped int
	unknown int
}

type Report struct {
	Tests      []Test
	skipHeader bool
}

func EmptyReport() Report {
	return Report{Tests: []Test{}, skipHeader: true}
}

func MakeReport(tests []Test) Report {
	return Report{Tests: tests, skipHeader: true}
}

func (report Report) AddTest(test Test) Report {
	return MakeReport(append(report.Tests, test))
}

func (report Report) BuildWith(raw [][]string) Report {
	filled := EmptyReport()

	for _, row := range raw {
		if row[0] == "Epic" && report.skipHeader {
			continue
		}

		filled = filled.AddTest(Test{
			epic:    normalizeLabel(row[0]),
			feature: normalizeLabel(row[1]),
			story:   normalizeLabel(row[2]),
			failed:  strToInt(row[3]),
			broken:  strToInt(row[4]),
			passed:  strToInt(row[5]),
			skipped: strToInt(row[6]),
			unknown: strToInt(row[7]),
		})
	}

	return filled
}

func (report Report) ToRaw() [][]string {
	var raw [][]string

	for _, test := range report.Tests {
		line := []string{
			test.epic,
			test.feature,
			test.story,
			strconv.Itoa(test.failed),
			strconv.Itoa(test.broken),
			strconv.Itoa(test.passed),
			strconv.Itoa(test.skipped),
			strconv.Itoa(test.unknown),
		}
		raw = append(raw, line)
	}

	return raw
}

func normalizeLabel(label string) string {
	if len(label) == 0 {
		return label
	}

	parts := strings.SplitN(label, ":", 2)

	if len(parts) != 2 {
		return label
	}

	return strings.ReplaceAll(parts[0], " ", "") + ":" + parts[1]
}

func strToInt(data string) int {
	converted, err := strconv.Atoi(data)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Conversion of %s to int failed", data), err)
	}
	return converted
}

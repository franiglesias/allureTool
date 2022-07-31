package csv_repository

import (
	"allureTool/application/domain"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func MakeTestFromRawData(datum []string) domain.Test {
	return domain.Test{
		Epic:    normalizeLabel(datum[0]),
		Feature: normalizeLabel(datum[1]),
		Story:   normalizeLabel(datum[2]),
		Failed:  strToInt(datum[3]),
		Broken:  strToInt(datum[4]),
		Passed:  strToInt(datum[5]),
		Skipped: strToInt(datum[6]),
		Unknown: strToInt(datum[7]),
	}
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

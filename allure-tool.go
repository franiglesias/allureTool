package main

import (
	"allureTool/application/adapters/for_getting_data/csv_repository"
	"allureTool/application/adapters/for_getting_data/zip_repository"
	"allureTool/application/use_cases/generate_report"
	"allureTool/application/use_cases/generate_report/analyze_execution"
	"allureTool/application/use_cases/generate_report/obtain_execution_data"
	"allureTool/application/use_cases/generate_report/summarize_data"
	"flag"
	"fmt"
	"github.com/spf13/afero"
	"log"
	"os"
)

func main() {

	generateReport := MakeGenerateReport()

	result, err := generateReport.Execute(generate_report.GenerateReportRequest{
		Filters:  []string{""},
		Projects: []string{"backend"},
	})

	if err != nil {
		log.Fatalf("Something failed. %#v\n", err)
	}

	fmt.Printf("%#v\n", result)
}

func MakeGenerateReport() generate_report.GenerateReport {
	zipCmd := flag.NewFlagSet("zip", flag.ExitOnError)
	zipArchive := zipCmd.String("archive", "", "Zip archive containing report")
	zipProject := zipCmd.String("project", "anon", "Project")

	csvCmd := flag.NewFlagSet("csv", flag.ExitOnError)
	csvFile := csvCmd.String("file", "", "File with tests data")
	csvProject := csvCmd.String("project", "anon", "Project")

	switch os.Args[1] {
	case "zip":
		zipCmd.Parse(os.Args[2:])
		return makeZip(*zipArchive, *zipProject)
	default:
		csvCmd.Parse(os.Args[2:])
		return makeCsv(*csvFile, *csvProject)
	}
}

func makeZip(zipFile, project string) generate_report.GenerateReport {
	archive := zip_repository.ZipArchive{
		Fs:   afero.NewOsFs(),
		Path: zipFile,
		Tmp:  "/tmp",
	}

	repository := zip_repository.MakeZipRepositoryFromArchive(archive)

	return generate_report.GenerateReport{
		Obtain:    obtain_execution_data.MakeObtainExecutionData(repository),
		Analyze:   analyze_execution.AnalyzeExecution{},
		Summarize: summarize_data.Summarize{},
	}
}

func makeCsv(csvFile, project string) generate_report.GenerateReport {
	repository := csv_repository.MakeCSVRepositoryForProjectAndFile(
		afero.NewOsFs(),
		project,
		csvFile,
	)

	return generate_report.GenerateReport{
		Obtain:    obtain_execution_data.MakeObtainExecutionData(repository),
		Analyze:   analyze_execution.AnalyzeExecution{},
		Summarize: summarize_data.Summarize{},
	}
}

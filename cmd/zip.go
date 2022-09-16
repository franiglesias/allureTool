/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"allureTool/application/adapters/for_getting_data/zip_repository"
	"allureTool/application/use_cases/generate_report"
	"allureTool/application/use_cases/generate_report/analyze_execution"
	"allureTool/application/use_cases/generate_report/obtain_execution_data"
	"allureTool/application/use_cases/generate_report/summarize_data"
	"fmt"
	"github.com/spf13/afero"
	"log"

	"github.com/spf13/cobra"
)

var zipArchive string

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Analyzing Allure ZIP file")
		fmt.Println("-------------------------")

		generateReport := makeZip(zipArchive, project)

		result, err := generateReport.Execute(generate_report.GenerateReportRequest{
			Filters:  []string{""},
			Projects: []string{project},
		})

		if err != nil {
			log.Fatalf("Something failed. %#v\n", err)
		}

		fmt.Printf("%#v\n", result)
	},
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

func init() {
	rootCmd.AddCommand(zipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	zipCmd.Flags().StringVar(&zipArchive, "archive", "", "Zip Archive to analyse")
}

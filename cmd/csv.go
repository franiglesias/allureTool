package cmd

import (
	"allureTool/application/adapters/for_getting_data/csv_repository"
	"allureTool/application/use_cases/generate_report"
	"allureTool/application/use_cases/generate_report/analyze_execution"
	"allureTool/application/use_cases/generate_report/obtain_execution_data"
	"allureTool/application/use_cases/generate_report/summarize_data"
	"fmt"
	"github.com/spf13/afero"
	"log"

	"github.com/spf13/cobra"
)

var csvFile string

// csvCmd represents the csv command
var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Analyzing Allure CSV file")
		fmt.Println("-------------------------")

		generateReport := makeCsv(csvFile, project)

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

func init() {
	rootCmd.AddCommand(csvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	zipCmd.Flags().StringVar(&csvFile, "file", "", "CSV file to analyse")

}

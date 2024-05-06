package cmd

import (
	"fmt"
	"os"

	"github.com/MalteHerrmann/ginkgo-parser/report"
	"github.com/spf13/cobra"
)

const (
	ExportTypeMD   = "md"
	ExportTypeHTML = "html"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ginkgo-parser REPORT_PATH EXPORT_PATH [flags]",
	Short: "Parses the generated JSON report from a Ginkgo run and exports the test information to Markdown or HTML.",
	Long: `This tool parses the JSON reports, that can be generated using the Ginkgo BDD
testing binary.

By default, the tool generates a Markdown document, that lists the ran
specs in a nested and sorted way.

Optionally, an HTML document can be generated.`,
	Args: cobra.ExactArgs(2),
	RunE: func(_ *cobra.Command, args []string) error {
		jsonFile := args[0]
		if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", jsonFile)
		}

		exportPath := os.Args[1]
		switch exportType {
		case ExportTypeMD:
			return report.ConvertGinkgoReportToMarkdown(jsonFile, exportPath)
		default:
			return fmt.Errorf("unsupported export type: %s", exportType)
		}
	},
}

// exportType defines which type of output should be generated with this tool.
var exportType string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(
		&exportType,
		"type",
		ExportTypeMD,
		fmt.Sprintf("The type of the exported file. Can be %q or %q.", ExportTypeMD, ExportTypeHTML),
	)
}

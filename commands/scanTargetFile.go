package commands

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pomcom/bagoScan/pkg/services"
	"github.com/spf13/cobra"
)

var (
	scanFileTargets = &cobra.Command{
		Use:   "scan-file-targets",
		Short: "Scans targets from a file with all activated modules.",
		Long: `Scans targets from a file with all activated modules. 
		If no config.yml is provided, all tools with default flag get executed against the target.`,
		Run: startScanFileTargets,
	}
)

func startScanFileTargets(cmd *cobra.Command, args []string) {
	// var targetFile string

	// cmd.Flags().StringVarP(&targetFile, "target-file", "f", "", "File containing the targets to scan")

	if targetFile == "" {
		fmt.Println("Please provide a target file using the --target-file flag.")
		return
	}

	targets, err := readTargetsFromFile(targetFile)
	if err != nil {
		fmt.Println("Error reading targets from file:", err)
		return
	}

	pentestService := services.NewTestRunnerService("config.yaml")
	pentestService.RunAllTools(targets)
}

func readTargetsFromFile(filePath string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	fileContent := string(fileBytes)
	targets := strings.Split(fileContent, "\n")
	return targets, nil
}

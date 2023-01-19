package commands

import (
	"fmt"

	"github.com/pomcom/bagoScan/pkg/services"
	"github.com/spf13/cobra"
)

var (
	scan = &cobra.Command{
		Use:   "scan",
		Short: "Scans a target with all activated modules.",
		Long:  `If no config.yml is provided, all tools with default flag get executed against the target.`,
		Run:   startScan,
	}
)

func startScan(cmd *cobra.Command, args []string) {

	targets, _ := cmd.Flags().GetStringSlice("target")
	fmt.Println("Targets:", targets)
	pentestService := services.NewTestRunnerService("config.yaml")
	pentestService.RunAllTools(targets)
}

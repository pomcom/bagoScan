package commands

import (
	"github.com/pomcom/bagoScan/pkg/services"
	"github.com/spf13/cobra"
)

var (
	runNmap = &cobra.Command{
		Use:   "nmap",
		Short: "Runs only nmap scan against target",
		Long:  `Runs an default nmap scan against the target. Uses the provided flags from config.yml, if present. Else default flags will be used.`,
		Run:   startNmap,
	}
)

func startNmap(cmd *cobra.Command, args []string) {

	target, _ := cmd.Flags().GetString("target")
	pentestService := services.NewTestRunnerService("config.yaml")
	pentestService.RunSingleTool("nmap", target)
}

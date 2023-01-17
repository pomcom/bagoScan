package commands

import (
	"github.com/pomcom/bagoScan/pkg/services"
	"github.com/spf13/cobra"
)

var (
	runTestssl = &cobra.Command{
		Use:   "testssl",
		Short: "Runs only testssl scan on target",
		Long:  `Runs an default testssl scan against the target. Uses the provided flags from config.yml, if present. Else default flags will be used.`,
		Run:   startNmap,
	}
)

func startTestssl(cmd *cobra.Command, args []string) {

	target, _ := cmd.Flags().GetString("target")
	pentestService := services.NewPentestService("config.yaml")
	pentestService.RunSingleTool("testssl", target)
}

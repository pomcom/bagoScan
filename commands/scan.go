package commands

import (
	"github.com/pomcom/bagoScan/pkg/config"
	"github.com/pomcom/bagoScan/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	scan = &cobra.Command{
		Use:   "scan",
		Short: "Scans a target with all activated modules.",
		Long:  `Scan currently supports testssl and nmap`,
		Run:   startScan,
	}
)

func startScan(cmd *cobra.Command, args []string) {
	target, _ := cmd.Flags().GetString("target")

	runner := utils.NewRunner()

	cfg := config.NewConfigHandler("config.yml")
	tools := cfg.GetTools()
	for _, tool := range tools {
		runner.AddTool(tool)
	}

	// for _, tool := range utils.GetSupportedTools() {
	// 	runner.AddTool(tool)
	// }

	// runner.AddTool(testssl.Testssl{})
	// runner.AddTool(nmap.Nmap{})
	runner.Run(target)
}

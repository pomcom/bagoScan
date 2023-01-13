package commands

import (
	"github.com/pomcom/bagoScan/tools"
	"github.com/pomcom/bagoScan/tools/testssl"
	"github.com/pomcom/bagoScan/utils"
	"github.com/spf13/cobra"
)

var (
	scan = &cobra.Command{
		Use:   "scan",
		Short: "Scans a target with all activated modules.",
		Long:  `Scan currently only checks for SSL configuration via testssl.sh.`,
		Run:   startScan,
	}
)

func startScan(cmd *cobra.Command, args []string) {
	target, _ := cmd.Flags().GetString("target")

	var tools []tools.Tool
	fileHandler := utils.Filehandler{}

	tools = append(tools, testssl.Testssl{})
	runner := utils.Runner{Tools: tools, Filehandler: fileHandler}
	runner.Run(target)
}

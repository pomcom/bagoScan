package commands

import (
	"github.com/pomcom/bagoScan/tools"
	"github.com/pomcom/bagoScan/tools/testssl"
	"github.com/pomcom/bagoScan/utils"
	"github.com/spf13/cobra"
)

var (
	scan = &cobra.Command{
		Use:   "bagoScan",
		Short: "bagoScan automatates some pentesting steps - work in progress.",
		Long:  `bagoScan currently only checks for SSL configuration via testssl.sh.`,
		Run:   startScan,
	}
)

func startScan(cmd *cobra.Command, args []string) {
	var tools []tools.Tool
	fileHandler := utils.Filehandler{}

	tools = append(tools, testssl.Testssl{})
	runner := utils.Runner{Tools: tools, Filehandler: fileHandler}
	runner.Run()
}

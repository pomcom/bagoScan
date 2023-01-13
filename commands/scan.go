package commands

import (
	"github.com/pomcom/bagoScan/tools"
	"github.com/pomcom/bagoScan/tools/nmap"
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

	runner := utils.NewRunner([]tools.Tool{
		testssl.Testssl{},
		nmap.Nmap{},
	}, utils.Filehandler{})
	runner.Run(target)
}

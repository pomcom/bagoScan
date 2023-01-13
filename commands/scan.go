package commands

import (
	"github.com/pomcom/bagoScan/tools/nmap"
	"github.com/pomcom/bagoScan/tools/testssl"
	"github.com/pomcom/bagoScan/utils"
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

	r := utils.NewRunner()
	r.AddTool(testssl.Testssl{})
	r.AddTool(nmap.Nmap{})
	r.Run(target)
}

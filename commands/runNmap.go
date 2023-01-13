package commands

import (
	"github.com/pomcom/bagoScan/pkg/tools/nmap"
	"github.com/pomcom/bagoScan/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	runNmap = &cobra.Command{
		Use:   "nmap",
		Short: "Runs only nmap scan on target",
		Long:  `Runs an default Nmap scan against the target.`,
		Run:   startNmap,
	}
)

func startNmap(cmd *cobra.Command, args []string) {
	target, _ := cmd.Flags().GetString("target")

	r := utils.NewRunner()
	r.AddTool(nmap.Nmap{})
	r.Run(target)
}

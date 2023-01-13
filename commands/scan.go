package commands

import (
	"fmt"

	"github.com/pomcom/bagoScan/pkg/config"
	"github.com/pomcom/bagoScan/pkg/tools/nmap"
	"github.com/pomcom/bagoScan/pkg/tools/testssl"
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

	config := config.NewConfigReader("config.yml")
	err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error reading config in startScan:", err)
		return
	}

	tools := config.GetTools()

	print("************************")
	print("tools in config:")

	fmt.Printf("%v", tools)

	print("************************")

	runner := utils.NewRunner()
	runner.AddTool(testssl.Testssl{})
	runner.AddTool(nmap.Nmap{})
	runner.Run(target)
}

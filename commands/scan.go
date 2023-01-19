package commands

import (
	"strings"

	"github.com/pomcom/bagoScan/pkg/services"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	scan = &cobra.Command{
		Use:   "scan",
		Short: "Scans a target with all activated modules.",
		Long: `If no config.yml is provided, all tools with default flag get executed against the target. Pass 
		a single target, or multiple targets using scan -t target1 -t target2. --target-file reads in a file with
		multiple targets line by line.`,
		Run: startScan,
	}
)

func startScan(cmd *cobra.Command, args []string) {
	target, _ := cmd.Flags().GetStringSlice("target")
	targetFile, _ := cmd.Flags().GetString("target-file")

	checkFlags()

	target = getTargets(targetFile, target)

	utils.Logger.Info("targets", zap.String("target", strings.Join(target, ",")))

	pentestService := services.NewTestRunnerService("config.yaml")
	pentestService.RunAllTools(target)
}

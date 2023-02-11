package commands

import (
	"strings"

	"github.com/pomcom/bagoScan/pkg/services"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	runDalfoxGauTest = &cobra.Command{
		Use:   "dalfoxgau",
		Short: "Runs dalfox xss and gau against target",
		Long:  `Runs dalfox and gau against the target.`,
		Run:   startDalfoxGau,
	}
)

func startDalfoxGau(cmd *cobra.Command, args []string) {

	target, _ := cmd.Flags().GetStringSlice("target")
	targetFile, _ := cmd.Flags().GetString("target-file")

	checkFlags()

	target = getTargets(targetFile, target)

	utils.Logger.Info("targets", zap.String("target", strings.Join(target, ",")))

	pentestService := services.NewTestRunnerService("config.yaml")
	pentestService.RunSingleTool("dalfoxGauTest", target)
}

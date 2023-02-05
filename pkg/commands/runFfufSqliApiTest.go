package commands

import (
	"strings"

	"github.com/pomcom/bagoScan/pkg/services"
	utils "github.com/pomcom/bagoScan/pkg/utils/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	runFfufSqliApiTest = &cobra.Command{
		Use:   "ffufApi",
		Short: "Runs ffuf api sql injection test against target",
		Long:  `Runs an api sqli injection  scan against the target. Uses the provided flags from config.yml, if present. Else default flags will be used.`,
		Run:   startFfufApiSqlTest,
	}
)

func startFfufApiSqlTest(cmd *cobra.Command, args []string) {

	target, _ := cmd.Flags().GetStringSlice("target")
	targetFile, _ := cmd.Flags().GetString("target-file")

	checkFlags()

	target = getTargets(targetFile, target)

	utils.Logger.Info("targets", zap.String("target", strings.Join(target, ",")))

	pentestService := services.NewTestRunnerService("config.yaml")
	pentestService.RunSingleTool("ffufSqliApiTest", target)
}

package commands

import (
	"github.com/spf13/cobra"
)

var target []string
var targetFile string

var (
	rootCmd = &cobra.Command{
		Use:           "bagoScan",
		Short:         "bagoScan executes various tools against a target.",
		Long:          `bagoScan executes various tools against a target. Work in Progress.`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {

	scan.Flags().StringSliceVarP(&target, "target", "t", []string{}, "The target to scan")
	scan.Flags().StringVarP(&targetFile, "target-file", "f", "", "File containing the targets to scan")

	runNmap.Flags().StringSliceVarP(&target, "target", "t", []string{}, "The target to scan")
	runNmap.Flags().StringVarP(&targetFile, "target-file", "f", "", "File containing the targets to scan")

	runSqlMap.Flags().StringSliceVarP(&target, "target", "t", []string{}, "The target to scan")
	runSqlMap.Flags().StringVarP(&targetFile, "target-file", "f", "", "File containing the targets to scan")

	runTestssl.Flags().StringSliceVarP(&target, "target", "t", []string{}, "The target to scan")
	runTestssl.Flags().StringVarP(&targetFile, "target-file", "f", "", "File containing the targets to scan")

	rootCmd.AddCommand(scan)
	rootCmd.AddCommand(runNmap)
	rootCmd.AddCommand(runSqlMap)
	rootCmd.AddCommand(runTestssl)

}

func Execute() error {
	return rootCmd.Execute()
}
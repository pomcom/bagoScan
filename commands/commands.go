package commands

import (
	"github.com/spf13/cobra"
)

var target []string

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

	scan.MarkFlagRequired("target")
	rootCmd.AddCommand(scan)

	runNmap.Flags().StringSliceVarP(&target, "target", "t", []string{}, "The target to scan")
	runNmap.MarkFlagRequired("target")

	runTestssl.Flags().StringSliceVarP(&target, "target", "t", []string{}, "The target to scan")
	runTestssl.MarkFlagRequired("target")

	rootCmd.AddCommand(runNmap)
	rootCmd.AddCommand(runTestssl)
}

func Execute() error {
	return rootCmd.Execute()
}

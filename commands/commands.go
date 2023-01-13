package commands

import (
	"github.com/spf13/cobra"
)

var target string

var (
	rootCmd = &cobra.Command{
		Use:           "bagoScan",
		Short:         "bagoScan runs mulitple security scans on a target.",
		Long:          `bagoScan runs multiple modular security scans on a given target. Work in Progress.`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	scan.Flags().StringVarP(&target, "target", "t", "", "The target to scan")
	scan.MarkFlagRequired("target")
	rootCmd.AddCommand(scan)

	runNmap.Flags().StringVarP(&target, "target", "t", "", "The target to scan")
	runNmap.MarkFlagRequired("target")
	rootCmd.AddCommand(runNmap)
}

func Execute() error {
	return rootCmd.Execute()
}

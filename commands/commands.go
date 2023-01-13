package commands

// root cmd - gets executed when no flags provided
import (
	"github.com/spf13/cobra"
)

var target string

var (
	rootCmd = &cobra.Command{
		Use:           "bagoScan",
		Short:         "bagoScan automatates some pentesting steps - work in progress.",
		Long:          `bagoScan currently only checks for SSL configuration via testssl.sh.`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	scan.Flags().StringVarP(&target, "target", "t", "", "The target to scan")
	scan.MarkFlagRequired("target")
	rootCmd.AddCommand(scan)
}

func Execute() error {
	return rootCmd.Execute()
}

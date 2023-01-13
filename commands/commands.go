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
	rootCmd.Flags().StringVarP(&target, "target", "t", "", "The target to scan")
	rootCmd.MarkFlagRequired("target")
}

func Execute() error {
	return rootCmd.Execute()
}

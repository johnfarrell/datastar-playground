package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dumpCmd)
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dumps current configuration and exits",
	Long:  "Used for validating current configuration values. Exits immediately, used more for testing.",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("DEBUG")
		logger.Info("INFO")
		logger.Warn("WARN")
		logger.Error("ERROR")
	},
}

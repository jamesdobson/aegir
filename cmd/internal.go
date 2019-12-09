package cmd

import (
	"github.com/spf13/cobra"
)

// internalCmd represents the internal command
var internalCmd = &cobra.Command{
	Use:   "internal",
	Short: "⛔️  Run internal commands",
}

func init() {
	rootCmd.AddCommand(internalCmd)
}

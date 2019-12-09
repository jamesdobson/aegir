package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "⚡️  Run a single job from inside the container",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go called")
	},
}

func init() {
	internalCmd.AddCommand(goCmd)
}

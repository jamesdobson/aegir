package cmd

import (
	"log"

	"github.com/jamesdobson/aegir/job"

	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "⚡️  Run a single job from inside the container",
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func init() {
	internalCmd.AddCommand(goCmd)
}

func execute() {
	task, err := job.MakeTask()
	if err != nil {
		log.Fatalf("Unable to start task: %v", err)
	}

	err = task.Execute()
	if err != nil {
		log.Fatalf("Task failed: %v", err)
	}
}

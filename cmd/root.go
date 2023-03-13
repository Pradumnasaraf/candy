package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "candy",
	Short: "Do all your tedious tasks with a single command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

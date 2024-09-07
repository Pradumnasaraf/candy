package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const CliVersion = "v1.8.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Know the installed version of candy",
	Long:  `This command will help you to know the installed version of candy`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("candy version:", CliVersion)
	},
}

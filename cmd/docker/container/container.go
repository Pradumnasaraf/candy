package container

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	runnnigContainer = []string{}
)

// containerCmd represents the container command
var ContainerCmd = &cobra.Command{
	Use:   "container [command] [flags]",
	Short: "Container related commands",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

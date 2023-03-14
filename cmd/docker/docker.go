package docker

import (
	"github.com/spf13/cobra"
)

// DockerCmd is the command docker related commands
var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker related commands. Like generating a Dockerfile for a language.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

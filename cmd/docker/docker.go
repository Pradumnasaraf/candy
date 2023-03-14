package docker

import (
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker related commands. Like generating a Dockerfile for a language.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

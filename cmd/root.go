package cmd

import (
	"log"
	"os"

	"github.com/Pradumnasaraf/candy/cmd/docker"
	"github.com/Pradumnasaraf/candy/cmd/kubernetes"
	"github.com/spf13/cobra"
)

// rootCmd is the root command for candy
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

func init() {

	// Subcommands for the root command
	rootCmd.AddCommand(jsonToYaml)
	rootCmd.AddCommand(yamlToJsonCmd)
	rootCmd.AddCommand(textToJsonCmd)
	rootCmd.AddCommand(docker.DockerCmd)
	rootCmd.AddCommand(kubernetes.KubernetesCmd)
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
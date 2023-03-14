package docker

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	language string
)

// dockerfileCmd is the command for generating a Dockerfile for a language
var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Generates a Dockerfile for a language.",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case language == "go" || language == "golang":
			createDockerfile(Golang)
		case language == "python" || language == "py":
			createDockerfile(Python)
		case language == "node" || language == "nodejs":
			createDockerfile(Node)
		case language == "java":
			createDockerfile(Java)
		case language == "ruby":
			createDockerfile(Ruby)
		default:
			fmt.Println("Sorry, we don't have a Dockerfile for that language yet.")
		}
	},
}

func createDockerfile(lang string) {
	file, err := os.Create("Dockerfile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(lang)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dockerfile generated successfully.")
}

func init() {
	DockerCmd.AddCommand(dockerfileCmd)

	dockerfileCmd.Flags().StringVarP(&language, "lang", "l", "", "Programming language to generate Dockerfile for.")
	dockerfileCmd.MarkFlagRequired("lang")
}

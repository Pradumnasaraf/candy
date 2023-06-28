package docker

import (
	"log"
	"os"
	"strings"

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

		language = strings.ToLower(language)
		switch {
		case language == "go" || language == "golang":
			createDockerfile(golang)
		case language == "python" || language == "py":
			createDockerfile(python)
		case language == "node" || language == "nodejs":
			createDockerfile(node)
		case language == "java":
			createDockerfile(java)
		case language == "ruby":
			createDockerfile(ruby)
		default:
			log.Print("Currently we don't support Dockerfile generation for " + language + ".")
		}
	},
}

func createDockerfile(lang string) {
	file, err := os.Create("Dockerfile")
	checkNilErr(err)
	defer file.Close()

	_, err = file.WriteString(lang)
	checkNilErr(err)

	log.Print("Dockerfile created successfully.")
}

func init() {
	dockerfileCmd.Flags().StringVarP(&language, "lang", "l", "", "Programming language to generate Dockerfile for.")
	err := dockerfileCmd.MarkFlagRequired("lang")
	checkNilErr(err)
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

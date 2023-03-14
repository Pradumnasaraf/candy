package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inputJsonFile  string
	outputYamlFile string
)

// yamlToJsonCmd represents the yamlToJson command
var JsonToYaml = &cobra.Command{
	Use:   "JTY",
	Short: "Converts a JSON into YAML.",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the JSON file
		vp := viper.New()
		vp.SetConfigFile(inputJsonFile)
		err := vp.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		// Write the YAML file
		if outputYamlFile == "" {
			outputYamlFile = "output.yaml"
		}
		vp.SetConfigFile(outputYamlFile)

		err = vp.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}

		if outputYamlFile == "" {
			fmt.Println("Opearation completed successfully. Check the output.yaml file.")
		} else {
			fmt.Println("Opearation completed successfully. Check the " + outputYamlFile + " file.")
		}
	},
}

func init() {
	rootCmd.AddCommand(JsonToYaml)

	JsonToYaml.Flags().StringVarP(&outputYamlFile, "output", "o", "", "Output YAML file name (default is output.yaml)")
	JsonToYaml.Flags().StringVarP(&inputJsonFile, "file", "f", "", "Input the JSON file name")
	JsonToYaml.MarkFlagRequired("file")
}

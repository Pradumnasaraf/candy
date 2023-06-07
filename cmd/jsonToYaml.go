package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inputJsonFile  string
	outputYamlFile string
)

// jsonToYaml is the command for converting JSON to YAML
var jsonToYaml = &cobra.Command{
	Use:   "JTY [flags]",
	Short: "Converts a JSON into YAML.",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the JSON file
		vp := viper.New()
		vp.SetConfigFile(inputJsonFile)
		err := vp.ReadInConfig()
		checkNilErr(err)

		// Write the YAML file
		if outputYamlFile == "" {
			outputYamlFile = "output.yaml"
		}
		vp.SetConfigFile(outputYamlFile)

		err = vp.WriteConfig()
		checkNilErr(err)

		if outputYamlFile == "" {
			log.Print("Operation completed successfully. Check the output.yaml file.")
		} else {
			log.Print("Operation completed successfully. Check the " + outputYamlFile + " file.")
		}
	},
}

func init() {

	// Flags for the JYT command
	jsonToYaml.Flags().StringVarP(&outputYamlFile, "output", "o", "", "Output YAML file name (default is output.yaml)")
	jsonToYaml.Flags().StringVarP(&inputJsonFile, "file", "f", "", "Input the JSON file name")
	err := jsonToYaml.MarkFlagRequired("file")
	checkNilErr(err)
}

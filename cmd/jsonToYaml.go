package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inputJsonFile  string
	outputYamlFile string
)

// jsonToYamlCmd is the command for converting JSON to YAML
var jsonToYamlCmd = &cobra.Command{
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
			fmt.Println("Operation completed successfully. Check the output.yaml file.")
		} else {
			fmt.Println("Operation completed successfully. Check the " + outputYamlFile + " file.")
		}
	},
}

func init() {

	// Flags for the JYT command
	jsonToYamlCmd.Flags().StringVarP(&outputYamlFile, "output", "o", "", "Output YAML file name (default is output.yaml)")
	jsonToYamlCmd.Flags().StringVarP(&inputJsonFile, "file", "f", "", "Input the JSON file name")
	err := jsonToYamlCmd.MarkFlagRequired("file")
	checkNilErr(err)
}

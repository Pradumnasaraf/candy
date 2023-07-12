package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inputYamlFile  string
	outputJsonFile string
)

// yamlToJsonCmd is the command for converting YAML to JSON
var yamlToJsonCmd = &cobra.Command{
	Use:   "YTJ [flags]",
	Short: "Converts a YAML into JSON.",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the YAML file
		vp := viper.New()
		vp.SetConfigFile(inputYamlFile)
		err := vp.ReadInConfig()
		checkNilErr(err)

		// Write the JSON file
		if outputJsonFile == "" {
			outputJsonFile = "output.json"
		}
		vp.SetConfigFile(outputJsonFile)
		err = vp.WriteConfig()
		checkNilErr(err)

		if outputJsonFile == "" {
			fmt.Println("Operation completed successfully. Check the output.json file.")
		} else {
			fmt.Println("Operation completed successfully. Check the " + outputJsonFile + " file.")
		}
	},
}

func init() {

	// Flags for the YTJ command
	yamlToJsonCmd.Flags().StringVarP(&outputJsonFile, "output", "o", "", "Output JSON file name (default is output.json)")
	yamlToJsonCmd.Flags().StringVarP(&inputYamlFile, "file", "f", "", "Input the YAML file name")

	err := yamlToJsonCmd.MarkFlagRequired("file")
	checkNilErr(err)

}

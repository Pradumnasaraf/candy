package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inputYamlFile  string
	outputJsonFile string
)

// yamlToJsonCmd represents the yamlToJson command
var yamlToJsonCmd = &cobra.Command{
	Use:   "YTJ",
	Short: "Converts a YAML into JSON.",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the YAML file
		vp := viper.New()
		vp.SetConfigFile(inputYamlFile)
		err := vp.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		// Write the JSON file
		if outputJsonFile == "" {
			outputJsonFile = "output.json"
		}
		vp.SetConfigFile(outputJsonFile)
		err = vp.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}

		if outputJsonFile == "" {
			fmt.Println("Opearation completed successfully. Check the output.json file.")
		} else {
			fmt.Println("Opearation completed successfully. Check the " + outputJsonFile + " file.")
		}
	},
}

func init() {
	rootCmd.AddCommand(yamlToJsonCmd)

	yamlToJsonCmd.Flags().StringVarP(&outputJsonFile, "output", "o", "", "Output JSON file name (default is output.json)")
	yamlToJsonCmd.Flags().StringVarP(&inputYamlFile, "file", "f", "", "Input the YAML file name")
	yamlToJsonCmd.MarkFlagRequired("file")
}

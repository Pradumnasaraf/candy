package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inputJsonFile string
)

// yamlToJsonCmd represents the yamlToJson command
var JsonToYaml = &cobra.Command{
	Use:   "JTY",
	Short: "Converts a JSON file to YAML",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the JSON file
		vp := viper.New()
		vp.SetConfigFile(inputJsonFile)
		err := vp.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		// Write the YAML file
		vp.SetConfigFile("output.yaml")
		err = vp.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("JSON file converted to YAML successfully. Check the output.yaml file.")
	},
}

func init() {
	rootCmd.AddCommand(JsonToYaml)

	JsonToYaml.Flags().StringVarP(&inputJsonFile, "file", "f", "", "Input the JSON file name")
	JsonToYaml.MarkFlagRequired("file")

}

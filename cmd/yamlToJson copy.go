package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inputYamlFile string
)

// yamlToJsonCmd represents the yamlToJson command
var yamlToJsonCmd = &cobra.Command{
	Use:   "YTJ",
	Short: "Converts a YAML file to JSON",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the YAML file
		vp := viper.New()
		vp.SetConfigFile(inputYamlFile)
		err := vp.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		// Write the JSON file
		vp.SetConfigFile("output.json")
		err = vp.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("YAML file converted to JSON successfully. Check the output.json file.")
	},
}

func init() {
	rootCmd.AddCommand(yamlToJsonCmd)

	yamlToJsonCmd.Flags().StringVarP(&inputYamlFile, "file", "f", "", "Input the YAML file name")
	yamlToJsonCmd.MarkFlagRequired("file")
}

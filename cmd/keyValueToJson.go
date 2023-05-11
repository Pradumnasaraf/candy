package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	inputTextFile   string
	outputJsonFile1 string
)

// textToJsonCmd represents the aa command
var keyValueToJson = &cobra.Command{
	Use:   "KVTJ [flags]",
	Short: "Converts Key-Value (text) to JSON.",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the input file
		content, err := os.ReadFile(inputTextFile)
		checkNilErr(err)

		// Check if the input file is empty
		if len(content) == 0 {
			checkNilErr(err)
		}

		// Convert the input file to JSON
		entries := strings.Split(string(content), "\n")
		m := make(map[string]string)
		for _, e := range entries {
			parts := strings.Split(e, "=")
			m[parts[0]] = parts[1]
		}
		jsonString, _ := json.MarshalIndent(m, "", "  ")

		// Print the output to the console
		if printOutput, _ := cmd.Flags().GetBool("print"); printOutput {
			log.Println(string(jsonString))
			return
		}

		if outputJsonFile1 == "" {
			outputJsonFile1 = "output.json"
		}

		// Write the output file
		file, err := os.Create(outputJsonFile1)
		checkNilErr(err)
		
		defer file.Close()

		_, err = file.WriteString(string(jsonString))
		checkNilErr(err)
	},
}

func init() {

	// Flags for the TTJ command
	keyValueToJson.Flags().StringVarP(&inputTextFile, "file", "f", "", "Input the text file name. Eg: keys.txt or .env")
	keyValueToJson.MarkFlagRequired("file")

	keyValueToJson.Flags().StringVarP(&outputJsonFile1, "output", "o", "", "Output JSON file name (default is output.json)")
	keyValueToJson.Flags().BoolP("print", "p", false, "Print the output to the console")
}

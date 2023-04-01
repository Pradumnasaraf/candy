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

// aaCmd represents the aa command
var textToJsonCmd = &cobra.Command{
	Use:   "TTJ",
	Short: "Convert Key-Value text to JSON.",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the input file
		content, err := os.ReadFile(inputTextFile)
		if err != nil {
			log.Fatal("Error reading input file: ", err)
		}

		// Check if the input file is empty
		if len(content) == 0 {
			log.Fatal("Input file is empty")
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
		if err != nil {
			log.Fatal("Error creating output file: ", err)
		}
		defer file.Close()

		_, err = file.WriteString(string(jsonString))
		if err != nil {
			log.Fatal("Error writing to output file: ", err)
		}
	},
}

func init() {

	// Flags for the TTJ command
	textToJsonCmd.Flags().StringVarP(&inputTextFile, "file", "f", "", "Input the text file name. Eg: keys.txt or .env")
	textToJsonCmd.MarkFlagRequired("file")

	textToJsonCmd.Flags().StringVarP(&outputJsonFile1, "output", "o", "", "Output JSON file name (default is output.json)")
	textToJsonCmd.Flags().BoolP("print", "p", false, "Print the output to the console")
}

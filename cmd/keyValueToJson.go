package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	inputTextFile   string
	outputJsonFile1 string
	printOutput     bool
)

// textToJsonCmd represents the aa command
var keyValueToJsonCmd = &cobra.Command{
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

			// Skip empty lines, comments and lines without "="
			if e == "" {
				continue
			}

			// Skip comments
			if strings.HasPrefix(e, "#") || strings.HasPrefix(e, "//") {
				continue
			}

			// Skip lines without "="
			if !strings.Contains(e, "=") {
				continue
			}

			// Split the line by "="
			parts := strings.Split(e, "=")

			var value string

			// If the value contains "=" then join the parts otherwise take the second part
			if len(parts) > 2 {
				value = strings.Join(parts[1:], "=")
			} else {
				value = parts[1]
			}

			// Remove " and ' from the value string
			value = strings.Trim(value, "\"")
			value = strings.Trim(value, "'")

			// Add the key-value pair to the map
			m[strings.TrimSpace(parts[0])] = strings.TrimSpace(value)
		}
		jsonString, _ := json.MarshalIndent(m, "", "  ")

		// Print the output to the console
		if printOutput {
			fmt.Println(string(jsonString))
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

		fmt.Println("Operation completed successfully. Check the", outputJsonFile1, "file.")
	},
}

func init() {

	// Flags for the TTJ command
	keyValueToJsonCmd.Flags().StringVarP(&inputTextFile, "file", "f", "", "Input the text file name. Eg: keys.txt or .env")
	err := keyValueToJsonCmd.MarkFlagRequired("file")
	checkNilErr(err)

	keyValueToJsonCmd.Flags().StringVarP(&outputJsonFile1, "output", "o", "", "Output JSON file name (default is output.json)")
	keyValueToJsonCmd.Flags().BoolVarP(&printOutput, "print", "p", false, "Print the output to the console")
}

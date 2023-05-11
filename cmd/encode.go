package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	isDecode bool
)

var encodeCmd = &cobra.Command{
	Use:   "encode [string to encode]",
	Short: "It encodes and decodes a string to base64 and vice versa.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		stringToEncode := args[0]

		if isDecode {
			decoded, err := base64.StdEncoding.DecodeString(stringToEncode)
			checkNilErr(err)
			fmt.Println(string(decoded))
			return
		}

		encoded := base64.StdEncoding.EncodeToString([]byte(stringToEncode))
		fmt.Println(encoded)
	},
}

func init() {
	encodeCmd.Flags().BoolVarP(&isDecode, "decode", "d", false, "Decode the string")
}

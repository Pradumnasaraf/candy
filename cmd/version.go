package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// This command gives the current version of the CLI tool.
var version = &cobra.Command{
	Use:   "version",
	Short: "Current version for the CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		//Reading the version.yml file.
		vp := viper.New()
		vp.SetConfigName("version")
		vp.SetConfigType("yml")
		vp.AddConfigPath("./")
		err := vp.ReadInConfig()
		checkNilErr(err)

		// Fetching the value of version.
		version := vp.GetString("version")
		fmt.Println(version)

	},
}

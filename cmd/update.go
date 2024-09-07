package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update candy to the latest version",
	Long:  `This command will help you to update candy to the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		update()
	},
}

func update() {
	cmd := exec.Command("go", "install", "github.com/Pradumnasaraf/candy@latest")
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	fmt.Printf("CLI updated successfully to the latest version (If any). Current version is: %s\n", CliVersion)
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcnksm/go-latest"
)

var (
	checkLatest bool
)

const (
	CLI_VERSION = "1.7.0"
	OWNER       = "Pradumnasaraf"
	REPO        = "candy"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "outputs the cli version",
	Run: func(cmd *cobra.Command, args []string) {

		if checkLatest {
			checkForNewVersion()
		} else {
			fmt.Println(CLI_VERSION)
		}

	},
}

func checkForNewVersion() {

	githubTag := &latest.GithubTag{
		Owner:             OWNER,
		Repository:        REPO,
		FixVersionStrFunc: latest.DeleteFrontV(),
	}

	res, err := latest.Check(githubTag, CLI_VERSION)

	if err != nil {
		fmt.Println("Unable to check for latest version. Check your internet connection")
		return
	}

	if res.Outdated {
		fmt.Printf("The latest version of candy is %s.\nPlease update to the latest version by running go get -u github.com/Pradumnasaraf/candy@latest", res.Current)
		return
	}

	fmt.Println("You are using the latest version of candy")

}

func init() {
	// Flags for the version command
	versionCmd.Flags().BoolVarP(&checkLatest, "latest", "l", false, "Check if the latest version is installed")
}

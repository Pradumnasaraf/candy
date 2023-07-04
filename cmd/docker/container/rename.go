package container

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a container",
	Run: func(cmd *cobra.Command, args []string) {
		renameContainer()
	},
}

func renameContainer() {
	ctx, cli := dockerClient()
	containerList := runnnigContainerList(cli, ctx)

	if len(containerList) == 0 {
		fmt.Println("No container running")
		return
	}

	for _, container := range containerList {
		runnnigContainer = append(runnnigContainer, container.Names[0][1:]+" - "+container.ID[:6])

	}

	// Prompt the user to select a container
	fmt.Println("CONTAINER NAME - CONTAINER ID")
	prompt := promptui.Select{
		Label: "Select a container to rename",
		Items: runnnigContainer,
	}

	_, result, err := prompt.Run()
	checkErr(err)

	// Get new name from the command line
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter new name: ")
	scanner.Scan()
	newName := scanner.Text()

	newName = strings.TrimSpace(newName)

	if newName == "" {
		fmt.Println("New name cannot be empty")
		return
	}

	slpit := strings.Split(result, " - ")

	// Rename the container
	err = cli.ContainerRename(ctx, slpit[0], newName)
	checkErr(err)

	fmt.Printf("Container %s renamed to %s\n", slpit[0], newName)

}

func init() {
	ContainerCmd.AddCommand(renameCmd)

}

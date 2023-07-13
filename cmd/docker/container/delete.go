package container

import (
	"fmt"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	deleteAll bool
)

// deleteContainerCmd represents the delete command
var deleteContainerCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a running container",
	Run: func(cmd *cobra.Command, args []string) {
		deleteContainer()
	},
}

func deleteContainer() {
	ctx, cli := dockerClient()
	containerList := runnnigContainerList(cli, ctx)

	if len(containerList) == 0 {
		fmt.Println("No container running")
		return
	}

	// If the flag --all is set, delete all containers
	if deleteAll {

		prompt := promptui.Select{
			Label: "Are you sure you want to delete all containers?",
			Items: []string{"Yes", "No"},
		}

		_, option, err := prompt.Run()
		checkErr(err)

		if option == "Yes" {

			for _, container := range containerList {
				fmt.Printf("Deleting container %s (%s)...\n", container.Names[0][1:], container.ID[:6])
				err = cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{Force: true})
				checkErr(err)
			}

			fmt.Println("Deleted all containers")
			return
		} else {
			fmt.Println("Containers not removed")
			return
		}

	}

	// If the flag --all is not set, delete a specific container
	for _, container := range containerList {
		runnnigContainer = append(runnnigContainer, container.Names[0][1:]+" - "+container.ID[:6])

	}

	prompt := promptui.Select{
		Label: "Select a containe to delete",
		Items: runnnigContainer,
	}

	_, conSelection, err := prompt.Run()
	checkErr(err)

	prompt = promptui.Select{
		Label: "Are you sure you want to delete " + conSelection + " ?",
		Items: []string{"Yes", "No"},
	}

	_, option, err := prompt.Run()
	checkErr(err)

	if option == "Yes" {
		slpit := strings.Split(conSelection, " - ")

		err = cli.ContainerRemove(ctx, slpit[1], types.ContainerRemoveOptions{Force: true})
		checkErr(err)

		fmt.Printf("Container %s deleted successfully\n", slpit[0])
	} else {
		fmt.Println("Container not removed")
		return
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	ContainerCmd.AddCommand(deleteContainerCmd)
	deleteContainerCmd.Flags().BoolVarP(&deleteAll, "all", "a", false, "Delete all containers")
}

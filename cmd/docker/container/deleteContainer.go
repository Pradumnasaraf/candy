package container

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var runnnigContainer = []string{}

// dc is the command for deleting a container
var deleteContainerCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a running container",
	Run: func(cmd *cobra.Command, args []string) {
		deleteContainer()
	},
}

func deleteContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	checkErr(err)

	defer cli.Close()

	containerList, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	checkErr(err)

	for _, container := range containerList {
		runnnigContainer = append(runnnigContainer, container.Names[0][1:]+" - "+container.ID[:6])

	}

	if len(runnnigContainer) == 0 {
		fmt.Println("No container running")
		return
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

		err = cli.ContainerRemove(ctx, slpit[0], types.ContainerRemoveOptions{Force: true})
		checkErr(err)

		fmt.Println("Container removed")
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
}

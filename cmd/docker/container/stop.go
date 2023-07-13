package container

import (
	"fmt"
	"log"
	"strings"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	stopall bool
)

// deleteallcontainerCmd represents the deleteallcontainer command
var stopAllContainersCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a running container",
	Run: func(cmd *cobra.Command, args []string) {
		stopContainers()
	},
}

func stopContainers() {
	ctx, cli := dockerClient()
	containerList := runnnigContainerList(cli, ctx)

	if len(containerList) == 0 {
		fmt.Println("No container running")
		return
	}

	// If the flag --all is set, delete all containers
	if stopall {

		prompt := promptui.Select{
			Label: "Are you sure you want to stop all containers?",
			Items: []string{"Yes", "No"},
		}

		_, option, err := prompt.Run()

		if err != nil {
			log.Fatal(err)
		}

		if option == "Yes" {

			for _, container := range containerList {
				fmt.Printf("Stopping container %s (%s)...\n", container.Names[0][1:], container.ID[:6])
				noWaitTimeout := 5 // to not wait for the container to exit gracefully
				err := cli.ContainerStop(ctx, container.ID, containertypes.StopOptions{Timeout: &noWaitTimeout})
				if err != nil {
					log.Fatal(err)
				}

			}
			fmt.Println("Stopped all containers")
			return

		} else {
			fmt.Println("Aborted")
			return
		}

	}

	// If the flag --all is not set, delete a specific container
	for _, container := range containerList {
		runnnigContainer = append(runnnigContainer, container.Names[0][1:]+" - "+container.ID[:6])
	}

	// Prompt the user to select a container
	fmt.Println("CONTAINER NAME - CONTAINER ID")
	prompt := promptui.Select{
		Label: "Select a container to stop",
		Items: runnnigContainer,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	prompt = promptui.Select{
		Label: "Are you sure you want to stop " + result + "?",
		Items: []string{"Yes", "No"},
	}

	_, option, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if option == "Yes" {

		split := strings.Split(result, " - ")

		noWaitTimeout := 5 // to not wait for the container to exit gracefully
		err := cli.ContainerStop(ctx, split[1], containertypes.StopOptions{Timeout: &noWaitTimeout})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Container + " + split[0] + " stopped successfully")

	}

}

func init() {

	ContainerCmd.AddCommand(stopAllContainersCmd)

	stopAllContainersCmd.Flags().BoolVarP(&stopall, "all", "a", false, "Delete all containers")

}

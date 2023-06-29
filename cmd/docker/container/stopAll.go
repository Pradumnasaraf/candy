package container

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// deleteallcontainerCmd represents the deleteallcontainer command
var stopAllContainersCmd = &cobra.Command{
	Use:   "stopall",
	Short: "Stop all running containers",

	Run: func(cmd *cobra.Command, args []string) {
		stopAllContainers()
	},
}

func stopAllContainers() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	if len(containers) == 0 {
		fmt.Println("No container running")
		return
	}

	for _, container := range containers {
		fmt.Printf("Stopping container %s (%s)...\n", container.Names[0][1:], container.ID[:6])
		noWaitTimeout := 5 // to not wait for the container to exit gracefully
		if err := cli.ContainerStop(ctx, container.ID, containertypes.StopOptions{Timeout: &noWaitTimeout}); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Stopped all containers")
}

func init() {

	ContainerCmd.AddCommand(stopAllContainersCmd)

}

package container

import (
	"context"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// dockerClient returns a docker client
func dockerClient() (context.Context, *client.Client) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	checkErr(err)
	defer func() { _ = cli.Close() }()

	return ctx, cli
}

// runnnigContainerList returns a list of running containers
func runnnigContainerList(cli *client.Client, ctx context.Context) []containertypes.Summary {

	containerList, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	checkErr(err)
	return containerList

}

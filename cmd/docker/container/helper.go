package container

import (
	"context"

	containertypes "github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

// dockerClient returns a docker client
func dockerClient() (context.Context, *client.Client) {
	ctx := context.Background()
	cli, err := client.New(client.FromEnv)
	checkErr(err)
	defer func() { _ = cli.Close() }()

	return ctx, cli
}

// runnnigContainerList returns a list of running containers
func runnnigContainerList(cli *client.Client, ctx context.Context) []containertypes.Summary {

	containerList, err := cli.ContainerList(ctx, client.ContainerListOptions{})
	checkErr(err)
	return containerList.Items

}

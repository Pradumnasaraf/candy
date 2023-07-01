package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func dockerClient() (context.Context, *client.Client) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	checkErr(err)
	defer cli.Close()

	return ctx, cli
}

func runnnigContainerList(cli *client.Client, ctx context.Context) []types.Container {

	containerList, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	checkErr(err)
	return containerList

}

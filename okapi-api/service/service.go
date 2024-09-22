package service

import (
	"context"
	"fmt"
	"io"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
)

func GetRunningDockerContainers() ([]types.Container, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.43"))
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	return containers, nil
}

func ExecPythonInContainer(containerId, command string) (string, error) {
    apiClient, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.43"))
    if err != nil {
        return "", fmt.Errorf("fail to create docker client: %w", err)
    }

    execConfig := container.ExecOptions{
        Cmd:          strslice.StrSlice([]string{"python3", "-c", command}),
        AttachStdout: true,
        AttachStderr: true,
    }

    execIDResp, err := apiClient.ContainerExecCreate(context.Background(), containerId, execConfig)
    if err != nil {
        return "", fmt.Errorf("fail to execute docker command: %w", err)
    }

    resp, err := apiClient.ContainerExecAttach(context.Background(), execIDResp.ID, container.ExecStartOptions{})
    if err != nil {
        return "", fmt.Errorf("fail to execute docker attachement: %w", err)
    }
    defer resp.Close()

    output, err := io.ReadAll(resp.Reader)
    if err != nil {
        return "", fmt.Errorf("fail to read output: %w", err)
    }

    return string(output), nil
}
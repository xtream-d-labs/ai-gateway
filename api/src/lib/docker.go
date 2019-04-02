package lib

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/log"
)

func init() {
	v, err := dockerServerVersion()
	if err != nil {
		log.Fatal("Cannot connect to the Docker daemon", err, nil)
	}
	log.Info("docker", nil, &log.Map{
		"version":         v.Version,
		"api-version":     v.APIVersion,
		"min-api-version": v.MinAPIVersion,
	})
}

func dockerServerVersion() (types.Version, error) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return types.Version{}, err
	}
	defer cli.Close()
	return cli.ServerVersion(context.Background())
}

const buildErrorResponse = "returned a non-zero code"

func buildDockerImage(ctx context.Context, config []byte, options types.ImageBuildOptions) error {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	reader, err := cli.ImageBuild(ctx, bufio.NewReader(bytes.NewReader(config)), options)
	if err != nil {
		return err
	}
	defer reader.Body.Close()

	// io.Copy(ioutil.Discard, reader.Body)
	buf := bytes.Buffer{}
	buf.ReadFrom(reader.Body) // wait for its done
	if strings.Contains(buf.String(), buildErrorResponse) {
		return fmt.Errorf(buf.String())
	}
	return nil
}

// RunAndWaitDockerContainer run containers and wait for its stopping
func RunAndWaitDockerContainer(ctx context.Context, name string,
	config *container.Config, host *container.HostConfig,
	net *network.NetworkingConfig) (*string, error) {
	ID, err := runDockerContainer(ctx, name, config, host, net)
	if err != nil {
		return nil, err
	}
	return waitforExitDockerContainer(ctx, swag.StringValue(ID))
}

func runDockerContainer(ctx context.Context, name string,
	config *container.Config, host *container.HostConfig,
	net *network.NetworkingConfig) (*string, error) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	res, err := cli.ContainerCreate(ctx, config, host, net, name)
	if err != nil {
		return nil, err
	}
	options := types.ContainerStartOptions{}
	err = cli.ContainerStart(ctx, res.ID, options)
	if err != nil {
		removeDockerContainer(ctx, cli, res.ID)
		return nil, err
	}
	return swag.String(res.ID), nil
}

func waitforExitDockerContainer(ctx context.Context, ID string) (*string, error) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	resultC, errC := cli.ContainerWait(ctx, ID, container.WaitConditionNotRunning)
	select {
	case <-resultC:
	case err = <-errC:
		removeDockerContainer(ctx, cli, ID)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	out, err := LogsOfDockerContainer(ctx, cli, ID)
	if err != nil {
		removeDockerContainer(ctx, cli, ID)
		return nil, err
	}
	removeDockerContainer(ctx, cli, ID)
	return out, nil
}

var (
	linebreak = regexp.MustCompile("[\x10-\x11]")
	printable = regexp.MustCompile("[^\x20-\x7E]")
)

// LogsOfDockerContainer retrieves logs from a specified container
func LogsOfDockerContainer(ctx context.Context, cli *docker.Client, ID string) (*string, error) {
	out, err := cli.ContainerLogs(ctx, ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return nil, err
	}
	defer out.Close()

	buf := bytes.Buffer{}
	io.Copy(&buf, out)

	log := linebreak.ReplaceAll(buf.Bytes(), []byte("|"))
	log = printable.ReplaceAll(log, []byte(" "))
	return swag.String(string(log)), nil
}

func removeDockerContainer(ctx context.Context, cli *docker.Client, ID string) error {
	options := types.ContainerRemoveOptions{Force: true}
	return cli.ContainerRemove(ctx, ID, options)
}

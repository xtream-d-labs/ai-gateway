package lib

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/log"
)

var (
	jobDockerFile []byte
	trimDockerTag = regexp.MustCompile(`:.*`)
	trimDockerLib = regexp.MustCompile(`library/`)
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
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("'GOPATH' should be set in advance.", nil, nil)
	}
	dir := fmt.Sprintf("%s/src/%s/templates", strings.TrimRight(gopath, "/"), config.ProjectPath)
	jobDockerFile, err = ioutil.ReadFile(filepath.Join(dir, "job_dockerfile"))
	if err != nil {
		log.Fatal("Could not read 'Dockerfile'.", err, &log.Map{
			"gopath": gopath,
			"dir":    dir,
		})
	}
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

func buildDockerImage(ctx context.Context, reader io.Reader, options types.ImageBuildOptions) error {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	builder, err := cli.ImageBuild(ctx, reader, options)
	if err != nil {
		return err
	}
	defer builder.Body.Close()

	// io.Copy(ioutil.Discard, reader.Body)
	buf := bytes.Buffer{}
	buf.ReadFrom(builder.Body) // wait for its done
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

// BuildJobImage builds a docker image for the job
func BuildJobImage(ctx context.Context, jobID, builder string, fqdnToPush bool) (*string, error) {
	job, err := db.GetJob(jobID)
	if err != nil {
		return nil, err
	}
	pkg, _, _, python, _, _, err := DetectImageContent(ctx, job.DockerImage)
	if err != nil {
		return nil, err
	}
	version := PyVersion(ctx, job.DockerImage, python)
	lib := fmt.Sprintf("/workspace/lib/python%s/site-packages", version)

	setup := "echo 'Unknown OS.'"
	switch pkg {
	case Apt:
		setup = "apt update && apt-get install -y bash"
	case Yum:
		setup = "yum install -y bash"
	case Apk:
		setup = "apk --no-cache add bash"
	}

	workspace := filepath.Join(config.Config.WorkspaceContainerDir, job.Workspaces[0])
	err = ioutil.WriteFile(filepath.Join(workspace, "Dockerfile"), []byte(fmt.Sprintf(
		string(jobDockerFile),
		job.DockerImage,
		builder,
		lib,
		setup,
		"\""+strings.Join(job.Commands, "\",\"")+"\"")), 0644)
	if err != nil {
		return nil, err
	}
	defer os.Remove(filepath.Join(workspace, "Dockerfile"))
	dir, err := ioutil.TempDir("", job.ID)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	err = archive(workspace, filepath.Join(dir, "build.tar.gz"))
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filepath.Join(dir, "build.tar.gz"))
	if err != nil {
		return nil, err
	}

	name := trimDockerTag.ReplaceAllString(job.DockerImage, "")
	name = trimDockerLib.ReplaceAllString(name, "")
	if fqdnToPush {
		if !strings.HasPrefix(name, config.Config.DockerRegistryHostName) {
			name = fmt.Sprintf(
				"%s/%s/%s",
				config.Config.DockerRegistryHostName,
				builder,
				name,
			)
		}
	}
	name = fmt.Sprintf(
		"%s:%s-%s",
		name,
		config.Config.JobImageTagPrefix,
		time.Now().Format("060102150405"),
	)
	options := types.ImageBuildOptions{
		Tags: []string{name},
		Labels: map[string]string{
			"com.rescale.scaleshift.name":            config.ProjectName,
			"com.rescale.scaleshift.image.built-as":  "job-excutable",
			"com.rescale.scaleshift.image.built-api": config.Config.APIVersion,
			"com.rescale.scaleshift.image.built-at":  time.Now().Format(time.RFC3339),
			"com.rescale.scaleshift.image.built-by":  builder,
			"com.rescale.scaleshift.image.built-on":  job.DockerImage,
		},
		SuppressOutput: true,
		NoCache:        true,
		ForceRemove:    true,
	}
	if err := buildDockerImage(ctx, file, options); err != nil {
		return nil, err
	}
	return swag.String(name), nil
}

// PushJobImage pushes a docker image for the job
func PushJobImage(ctx context.Context, imageName, authConfig string) error {
	cli, e := docker.NewEnvClient()
	if e != nil {
		return e
	}
	defer cli.Close()

	options := types.ImagePushOptions{}
	if strings.HasPrefix(imageName, config.Config.DockerRegistryHostName) {
		converted, err := json.Marshal(&types.AuthConfig{
			ServerAddress: config.Config.DockerRegistryHostName,
			Username:      config.Config.DockerRegistryUserName,
			Password:      authConfig,
		})
		if err != nil {
			return err
		}
		options.RegistryAuth = base64.URLEncoding.EncodeToString(converted)
	}
	if strings.HasPrefix(imageName, config.Config.NgcRegistryHostName) {
		converted, err := json.Marshal(&types.AuthConfig{
			ServerAddress: config.Config.NgcRegistryHostName,
			Username:      config.Config.NgcRegistryUserName,
			Password:      authConfig,
		})
		if err != nil {
			return err
		}
		options.RegistryAuth = base64.URLEncoding.EncodeToString(converted)
	}
	reader, err := cli.ImagePush(ctx, imageName, options)
	if err != nil {
		return err
	}
	defer reader.Close()

	// io.Copy(ioutil.Discard, reader.Body)
	buf := bytes.Buffer{}
	buf.ReadFrom(reader) // wait for its done

	if strings.Contains(buf.String(), buildErrorResponse) {
		return fmt.Errorf(buf.String())
	}
	return nil
}

// DeleteImage deletes a docker image
func DeleteImage(ctx context.Context, imageName string) error {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	options := types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	}
	_, err = cli.ImageRemove(ctx, imageName, options)
	if err != nil {
		return err
	}
	return nil
}

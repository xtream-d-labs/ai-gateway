package lib

import (
	"context"
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
	"docker.io/go-docker/api/types/mount"
	"docker.io/go-docker/api/types/strslice"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/log"
)

const (
	sFileName = "Singularity"
)

var (
	sFile  []byte
	sImgNm = regexp.MustCompile("[^a-zA-Z0-9.]")
)

func init() {
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("'GOPATH' should be set in advance.", nil, nil)
	}
	dir := fmt.Sprintf("%s/src/%s/templates", strings.TrimRight(gopath, "/"), config.ProjectPath)
	var err error
	sFile, err = ioutil.ReadFile(filepath.Join(dir, sFileName))
	if err != nil {
		log.Fatal("Could not read 'Singularity'.", err, &log.Map{
			"gopath": gopath,
			"dir":    dir,
		})
	}
}

// BuildSingularityImage builds singularity image
func BuildSingularityImage(jobID, authConfig, builder string) (*string, error) {
	job, err := db.GetJob(jobID)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	pullSingularityImage(ctx)

	dir := filepath.Join(config.Config.SingImgContainerDir, job.ID)
	if err = os.MkdirAll(dir, 0755); err != nil { // nolint
		return nil, err
	}
	pkg, _, _, python, _, _, err := DetectImageContent(ctx, job.DockerImage)
	if err != nil {
		return nil, err
	}
	version := PyVersion(ctx, job.DockerImage, python)
	lib := fmt.Sprintf("/workspace/lib/python%s/site-packages", version)

	setup := "echo 'Unknown OS is detected.'"
	switch pkg {
	case Apt:
		setup = "apt update && apt-get install -y bash"
	case Yum:
		setup = "yum install -y bash"
	case Apk:
		setup = "apk --no-cache add bash"
	}
	sfile := fmt.Sprintf(
		string(sFile),
		job.DockerImage,
		builder,
		lib,
		setup,
		strings.Join(job.Commands, " "))
	log.Debug("Singularityfile", nil, &log.Map{
		"content": sfile,
	})
	err = ioutil.WriteFile(filepath.Join(dir, sFileName), []byte(sfile), 0644)
	if err != nil {
		return nil, err
	}
	sImgName := fmt.Sprintf("%s-%d.simg", sImgNm.ReplaceAllString(job.DockerImage, "-"), time.Now().Unix())
	name := fmt.Sprintf("build-singularity-image-%d", time.Now().Unix())
	cmds := []string{"build", sImgName, sFileName}
	envs := []string{"SINGULARITY_CACHEDIR=/home/.cache"}
	// envs := []string{}
	if strings.HasPrefix(job.DockerImage, config.Config.DockerRegistryHostName) {
		envs = append(envs, fmt.Sprintf("SINGULARITY_DOCKER_USERNAME=%s", config.Config.DockerRegistryUserName))
		envs = append(envs, fmt.Sprintf("SINGULARITY_DOCKER_PASSWORD=%s", authConfig))
	}
	cfg := &container.Config{
		Image:      config.Config.SingImg,
		Cmd:        strslice.StrSlice(cmds),
		Env:        envs,
		WorkingDir: "/home/singularity/image",
	}
	mounts := []mount.Mount{
		mount.Mount{
			Type:   mount.TypeBind,
			Source: filepath.Join(config.Config.SingImgHostPath, job.ID),
			Target: "/home/singularity/image",
		},
		// FIXME
		// mount.Mount{
		// 	Type:   mount.TypeBind,
		// 	Source: filepath.Join(config.Config.SingImgHostPath, ".base"),
		// 	Target: "/home/.base",
		// },
		// mount.Mount{
		// 	Type:   mount.TypeBind,
		// 	Source: filepath.Join(config.Config.SingImgHostPath, ".cache"),
		// 	Target: "/home/.cache",
		// },
	}
	if len(job.Workspaces) > 0 {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: filepath.Join(config.Config.WorkspaceHostDir, job.Workspaces[0]),
			Target: "/var/workspace",
		})
	}
	host := &container.HostConfig{
		Privileged: true,
		Mounts:     mounts,
	}
	logs, err := RunAndWaitDockerContainer(ctx, name, cfg, host, nil)
	if err != nil {
		return nil, err
	}
	log.Debug(fmt.Sprintf("Run&Wait@buildSingularityImage: %s", swag.StringValue(logs)), nil, nil)
	return swag.String(filepath.Join(dir, sImgName)), nil
}

func pullSingularityImage(ctx context.Context) error {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	reader, err := cli.ImagePull(ctx, config.Config.SingImg, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	io.Copy(ioutil.Discard, reader) // wait for its done
	return nil
}

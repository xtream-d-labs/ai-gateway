package lib

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/strslice"
	docker "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/go-openapi/swag"
	"github.com/xtream-d-labs/ai-gateway/api/src/config"
	"github.com/xtream-d-labs/ai-gateway/api/src/log"
)

const (
	pipConfigName         = "pip.conf"
	ipythonConfigName     = "ipython_config.py"
	jupyterDockerfileName = "jupyter_dockerfile"
	jupyterConfigName     = "jupyter_notebook_config.py"
	jupyterPythonSample   = "python-get-started.ipynb"
	jupyterBashWithApt    = "bash-scripts-apt.ipynb"
	jupyterBashWithYum    = "bash-scripts-yum.ipynb"
	jupyterBashWithApk    = "bash-scripts-apk.ipynb"
)

var (
	pipConfig          []byte
	ipythonConfig      []byte
	jupyterDockerfile  []byte
	jupyterConfig      []byte
	jupyterPythonNote  []byte
	jupyterAptNote     []byte
	jupyterYumNote     []byte
	jupyterApkNote     []byte
	minimalJupyterPort uint16 = 30000
)

// PackageManager like apt / yum
type PackageManager int

// PackageManager values
const (
	Unknown PackageManager = iota
	Apk
	Apt
	Yum
)

var (
	nonascii = regexp.MustCompile("[^a-zA-Z0-9.]")
	pyver    = regexp.MustCompile(".*ython ")
)

func init() {
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("'GOPATH' should be set in advance.", nil, nil)
	}
	dir := fmt.Sprintf("%s/src/%s/templates", strings.TrimRight(gopath, "/"), config.ProjectPath)
	var err error
	pipConfig, err = ioutil.ReadFile(filepath.Join(dir, pipConfigName))
	if err != nil {
		log.Fatal("Could not read 'pip.conf'.", err, &log.Map{
			"gopath": gopath,
			"dir":    dir,
		})
	}
	ipythonConfig, _ = ioutil.ReadFile(filepath.Join(dir, ipythonConfigName))
	jupyterDockerfile, _ = ioutil.ReadFile(filepath.Join(dir, jupyterDockerfileName))
	jupyterConfig, _ = ioutil.ReadFile(filepath.Join(dir, jupyterConfigName))
	jupyterPythonNote, _ = ioutil.ReadFile(filepath.Join(dir, jupyterPythonSample))
	jupyterAptNote, _ = ioutil.ReadFile(filepath.Join(dir, jupyterBashWithApt))
	jupyterYumNote, _ = ioutil.ReadFile(filepath.Join(dir, jupyterBashWithYum))
	jupyterApkNote, _ = ioutil.ReadFile(filepath.Join(dir, jupyterBashWithApk))
	minimalJupyterPort = config.Config.JupyterMinimumPort
}

// WrapWithJupyterNotebook builds wrapper docker images
func WrapWithJupyterNotebook(ctx context.Context, id, image, builder string) (*string, error) {
	pkg, setup, python, lib, workdir, err := DetectImageContent(ctx, image, false)
	if err != nil {
		return nil, err
	}
	config, err := makeJupyterNotebookBuildContext(pkg, image, workdir, lib, setup, python)
	if err != nil {
		return nil, err
	}
	return buildJupyterNotebookImage(ctx, config, id, image, builder)
}

// DetectImageContent detects image conditions
func DetectImageContent(ctx context.Context, image string, short bool) (PackageManager, string, string, string, string, error) {
	pkg := Unknown
	setupScripts := ""

	python := checkPython(ctx, image)
	if python == nil {
		return pkg, "", "", "", "", fmt.Errorf("Python is not installed on the image: %s", image)
	}
	version := PyVersion(ctx, image, swag.StringValue(python))

	if which(ctx, image, "apt-get") {
		pkg = Apt
		setupScripts = "apt update && apt-get install -y bash wget jq " +
			"&& wget -qO /sbin/tini https://github.com/krallin/tini/releases/download/v0.19.0/tini " +
			"&& chmod +x /sbin/tini"
	}
	if which(ctx, image, "yum") {
		pkg = Yum
		setupScripts = "yum install -y epel-release bash wget jq " +
			"&& wget -qO /sbin/tini https://github.com/krallin/tini/releases/download/v0.19.0/tini " +
			"&& chmod +x /sbin/tini"
	}
	if which(ctx, image, "apk") {
		pkg = Apk
		pip := checkPip(ctx, image)
		if pip == nil {
			return pkg, "", "", "", "", fmt.Errorf("pip is not installed on the image: %s", image)
		}
		setupScripts = fmt.Sprintf(
			"apk --no-cache add bash wget tini build-base linux-headers musl-dev jq && %s install cython",
			swag.StringValue(pip),
		)
	}
	if short {
		return pkg, "", swag.StringValue(python), "", "", nil
	}
	workdir := DetectImageWorkDir(ctx, image)
	lib := fmt.Sprintf("%s/workspace/lib/python%s/site-packages", workdir, version)
	return pkg, setupScripts, swag.StringValue(python), lib, workdir, nil
}

// DetectImageWorkDir detects image working directory
func DetectImageWorkDir(ctx context.Context, image string) string {
	if inspected, err := inspect(ctx, image); err == nil {
		if _, ok := inspected.ContainerConfig.Labels["com.nvidia.volumes.needed"]; ok {
			return "/workspace"
		}
	}
	return "/root/notebook"
}

func checkPython(ctx context.Context, image string) *string {
	if which(ctx, image, "python") {
		return swag.String("python")
	}
	if which(ctx, image, "python3") {
		return swag.String("python3")
	}
	return nil
}

func checkPip(ctx context.Context, image string) *string {
	if which(ctx, image, "pip") {
		return swag.String("pip")
	}
	if which(ctx, image, "pip3") {
		return swag.String("pip3")
	}
	return nil
}

func which(ctx context.Context, image, cmd string) bool {
	name := fmt.Sprintf("check-command-existence-%d", time.Now().Unix())
	logs, _ := RunAndWaitDockerContainer(ctx, name, &container.Config{
		Image:      image,
		Entrypoint: strslice.StrSlice([]string{"which"}),
		Cmd:        strslice.StrSlice([]string{cmd}),
	}, nil, nil)
	return strings.Contains(swag.StringValue(logs), cmd)
}

func inspect(ctx context.Context, image string) (*types.ImageInspect, error) {
	cli, err := docker.NewClientWithOpts(docker.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	inspect, _, err := cli.ImageInspectWithRaw(ctx, image)
	if err != nil {
		return nil, err
	}
	return &inspect, nil
}

// PyVersion detect the python version inside the image
func PyVersion(ctx context.Context, image, python string) string {
	name := fmt.Sprintf("check-python-version-%d", time.Now().Unix())
	logs, _ := RunAndWaitDockerContainer(ctx, name, &container.Config{
		Image:      image,
		Entrypoint: strslice.StrSlice([]string{python}),
		Cmd:        strslice.StrSlice([]string{"-V"}),
	}, nil, nil)
	version := pyver.ReplaceAllString(swag.StringValue(logs), "")
	if len(version) > 3 {
		version = version[0:3]
	}
	return version
}

func makeJupyterNotebookBuildContext(pkg PackageManager, image, workdir, lib, setup, python string) ([]byte, error) {
	contents := []*tarContent{}
	pipCfg := fmt.Sprintf(string(pipConfig), workdir+"/workspace")
	content := []byte(pipCfg)
	contents = append(contents, &tarContent{
		body: content,
		header: &tar.Header{
			Name: pipConfigName,
			Mode: 0666,
			Size: int64(len(content)),
		},
	})
	ipythonCfg := fmt.Sprintf(string(ipythonConfig), lib)
	content = []byte(ipythonCfg)
	contents = append(contents, &tarContent{
		body: content,
		header: &tar.Header{
			Name: ipythonConfigName,
			Mode: 0666,
			Size: int64(len(content)),
		},
	})
	path := ""
	if len(workdir) > 0 {
		lib = fmt.Sprintf("%s:%s", workdir, lib)
		path = fmt.Sprintf("%s/workspace/bin:", workdir)
	}
	dockerfile := fmt.Sprintf(string(jupyterDockerfile), image, workdir, lib, path, setup, python, python, python)
	log.Debug(fmt.Sprintf("MakeJupyterNotebook -> Dockerfile: %s", dockerfile), nil, nil)
	content = []byte(dockerfile)
	contents = append(contents, &tarContent{
		body: content,
		header: &tar.Header{
			Name: "Dockerfile",
			Mode: 0666,
			Size: int64(len(content)),
		},
	})
	jupyterCfg := fmt.Sprintf(string(jupyterConfig), workdir)
	content = []byte(jupyterCfg)
	contents = append(contents, &tarContent{
		body: content,
		header: &tar.Header{
			Name: "jupyter_notebook_config.py",
			Mode: 0666,
			Size: int64(len(content)),
		},
	})
	contents = append(contents, &tarContent{
		body: jupyterPythonNote,
		header: &tar.Header{
			Name: "python-get-started.ipynb",
			Mode: 0666,
			Size: int64(len(jupyterPythonNote)),
		},
	})
	switch pkg {
	case Apt:
		contents = append(contents, &tarContent{
			body: jupyterAptNote,
			header: &tar.Header{
				Name: "bash-scripts.ipynb",
				Mode: 0666,
				Size: int64(len(jupyterAptNote)),
			},
		})
	case Yum:
		contents = append(contents, &tarContent{
			body: jupyterYumNote,
			header: &tar.Header{
				Name: "bash-scripts.ipynb",
				Mode: 0666,
				Size: int64(len(jupyterYumNote)),
			},
		})
	case Apk:
		contents = append(contents, &tarContent{
			body: jupyterApkNote,
			header: &tar.Header{
				Name: "bash-scripts.ipynb",
				Mode: 0666,
				Size: int64(len(jupyterApkNote)),
			},
		})
	}
	return targz(contents)
}

func buildJupyterNotebookImage(ctx context.Context, cfg []byte, id, image, builder string) (*string, error) {
	name := fmt.Sprintf("%s/%s", config.Config.JupyterImageNamespace, image)
	options := types.ImageBuildOptions{
		Tags: []string{name},
		Labels: map[string]string{
			"com.aigateway.name":            config.ProjectName,
			"com.aigateway.image.original":  id,
			"com.aigateway.image.built-as":  "jupyter-notebook",
			"com.aigateway.image.built-api": config.Config.APIVersion,
			"com.aigateway.image.built-at":  time.Now().Format(time.RFC3339),
			"com.aigateway.image.built-by":  builder,
			"com.aigateway.image.built-on":  image,
		},
		NoCache:     true,
		ForceRemove: true,
	}
	reader := bufio.NewReader(bytes.NewReader(cfg))
	if err := buildDockerImage(ctx, io.Reader(reader), options); err != nil {
		return nil, err
	}
	return swag.String(name), nil
}

// RunJupyterNotebook runs the specified image as a jupter notebook
func RunJupyterNotebook(ctx context.Context, originalImage, wrappedImage, workdirHost, workdir string) (*string, error) {
	cli, err := docker.NewClientWithOpts(docker.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	port, notebooks, err := availablePortAndNotebooks(ctx, cli)
	if err != nil {
		return nil, err
	}
	exposed, bindings, err := nat.ParsePortSpecs([]string{fmt.Sprintf("0.0.0.0:%s:8888/tcp", port)})
	if err != nil {
		return nil, err
	}
	identity := fmt.Sprintf("%d", time.Now().Unix())

	img := strings.TrimPrefix(originalImage, fmt.Sprintf("%s/", config.Config.JupyterImageNamespace))
	if swag.IsZero(workdirHost) {
		workdirHost = fmt.Sprintf("%s-%s", nonascii.ReplaceAllString(img, "-"), identity)
		if err = os.MkdirAll(filepath.Join(config.Config.WorkspaceContainerDir, workdirHost), 0755); err != nil { // nolint
			return nil, err
		}
	}
	cfg := &container.Config{
		Image:        wrappedImage,
		ExposedPorts: exposed,
		WorkingDir:   workdir,
		Labels: map[string]string{
			"com.aigateway.container.run-api": config.Config.APIVersion,
			"com.aigateway.container.started": time.Now().Format(time.RFC3339),
			"com.aigateway.container.publish": port,
			"com.aigateway.container.mounted": workdirHost,
		},
	}
	host := &container.HostConfig{
		PortBindings: bindings,
		Mounts: []mount.Mount{{
			Type:   mount.TypeBind,
			Source: filepath.Join(config.Config.WorkspaceHostDir, workdirHost), // Host machine
			Target: fmt.Sprintf("%s/workspace", workdir),
		}},
	}
	gpus := int(config.Config.NvidiaGPUsPerContainer)
	if 0 < config.Config.NvidiaGPUs && (notebooks+1)*gpus <= int(config.Config.NvidiaGPUs) {
		cfg.Labels["com.aigateway.container.gpus"] = fmt.Sprintf("%d", gpus)

		capabilities := [][]string{}
		capabilities = append(capabilities, []string{"gpu"})
		host.Resources = container.Resources{
			DeviceRequests: []container.DeviceRequest{{
				Count:        gpus,
				Capabilities: capabilities,
			}},
		}
	}
	log.Debug(fmt.Sprintf("RunJupyterNotebook: %+v, %+v", cfg, host), nil, nil)

	net := &network.NetworkingConfig{}
	name := fmt.Sprintf("ipynb-%s", identity)
	ID, err := runDockerContainer(ctx, name, cfg, host, net)
	return ID, err
}

func availablePortAndNotebooks(ctx context.Context, cli *docker.Client) (string, int, error) {
	var available uint16
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Quiet: true,
		All:   true,
	})
	if err != nil {
		return "", 0, err
	}
	count := 0
	for _, container := range containers {
		for _, port := range container.Ports {
			if port.PublicPort >= available {
				available = port.PublicPort + 1
			}
		}
		if as, ok1 := container.Labels["com.aigateway.image.built-as"]; ok1 {
			if _, ok2 := container.Labels["com.aigateway.container.publish"]; ok2 {
				if strings.EqualFold(as, "jupyter-notebook") {
					count++
				}
			}
		}
	}
	if available < minimalJupyterPort {
		available = minimalJupyterPort
	}
	return fmt.Sprintf("%d", available), count, nil
}

// RunningNotebooks returns containers which was build by this app
func RunningNotebooks(ctx context.Context, cli *docker.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Quiet: true,
		All:   true,
	})
	if err != nil {
		return nil, err
	}
	result := []types.Container{}
	for _, container := range containers {
		if as, ok1 := container.Labels["com.aigateway.image.built-as"]; ok1 {
			if _, ok2 := container.Labels["com.aigateway.container.publish"]; ok2 {
				if strings.EqualFold(as, "jupyter-notebook") {
					result = append(result, container)
				}
			}
		}
	}
	return result, nil
}

// ContainerAttrs returns original image, published port & started time
func ContainerAttrs(labels map[string]string) (string, int64, int64, time.Time) {
	var image string
	var port int64
	var gpus int64
	started := time.Now()
	for key, value := range labels {
		switch key {
		case "com.aigateway.image.built-on":
			image = value
		case "com.aigateway.container.publish":
			if candidate, err := strconv.ParseInt(value, 10, 64); err == nil {
				port = candidate
			}
		case "com.aigateway.container.gpus":
			if candidate, err := strconv.ParseInt(value, 10, 64); err == nil {
				gpus = candidate
			}
		case "com.aigateway.container.started":
			if candidate, err := time.Parse(time.RFC3339, value); err == nil {
				started = candidate
			}
		}
	}
	return image, port, gpus, started
}

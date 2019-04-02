package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/mount"
	"docker.io/go-docker/api/types/strslice"
)

// IPython defines a notebook
type IPython struct {
	Meta struct {
		KernelSpec struct {
			Name string `json:"name"`
			Lang string `json:"language"`
		} `json:"kernelspec"`
		Lang struct {
			Name    string `json:"name"`
			FileExt string `json:"file_extension"`
		} `json:"language_info"`
	} `json:"metadata"`
}

// ParseIPython parses a notebook
func ParseIPython(path string) *IPython {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	ipy := &IPython{}
	err = json.Unmarshal(raw, ipy)
	if err != nil {
		return nil
	}
	return ipy
}

// ConvertNotebook converts Jupyter notebook to Python scripts
func ConvertNotebook(ctx context.Context, image, ipynb, to, workdir string, mounts []mount.Mount) error {
	_, err := RunAndWaitDockerContainer(ctx, "container-name", &container.Config{
		Image:      image,
		Entrypoint: strslice.StrSlice([]string{"jupyter"}),
		Cmd:        strslice.StrSlice([]string{"nbconvert", "--to", to, ipynb}),
		WorkingDir: fmt.Sprintf("%s/workspace", workdir),
	}, &container.HostConfig{Mounts: mounts}, nil)
	return err
}

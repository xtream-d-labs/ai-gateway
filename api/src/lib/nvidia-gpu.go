package lib

import (
	"context"
	"regexp"
	"strconv"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/go-openapi/swag"
	"github.com/xtream-d-labs/ai-gateway/api/src/config"
)

func init() {
	i, err := DeviceCount(context.Background())
	if err != nil {
		config.Config.NvidiaGPUs = 0
		return
	}
	config.Config.NvidiaGPUs = swag.Int64Value(i)
}

var space = regexp.MustCompile(`[^0-9]`)

// NvididSMI returns the result of nvidia-smi
func NvididSMI(ctx context.Context, args []string) (*string, error) {
	capabilities := [][]string{}
	capabilities = append(capabilities, []string{"gpu"})

	return RunAndWaitDockerContainer(ctx, "nvidia-smi-runner", &container.Config{
		Image:      config.Config.NvidiaSmiContainer,
		Entrypoint: strslice.StrSlice([]string{"nvidia-smi"}),
		Cmd:        args,
	}, &container.HostConfig{
		Resources: container.Resources{
			DeviceRequests: []container.DeviceRequest{{
				Count:        -1,
				Capabilities: capabilities,
			}},
		},
	}, nil)
}

// DeviceCount returns number of NVIDIA GPUS
func DeviceCount(ctx context.Context) (*int64, error) {
	logs, err := NvididSMI(ctx, strslice.StrSlice([]string{
		"--format=csv,noheader,nounits",
		"--query-gpu=count",
	}))
	if err != nil {
		return nil, err
	}
	candidate := space.ReplaceAllString(swag.StringValue(logs), "")
	i64, e := strconv.ParseInt(candidate, 10, 64)
	if e != nil {
		return nil, e
	}
	return swag.Int64(i64), nil
}

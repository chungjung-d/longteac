package process

import (
	"net"
	"os"

	"github.com/chungjung-d/longteac/config"
	"github.com/chungjung-d/longteac/process/mounts"
)

func settingMount(containerDirPath string, ociConfig *config.OCIConfig) error {
	for _, mount := range ociConfig.MountsConfig {
		if err := mounts.Mount(mount); err != nil {
			return err
		}
	}

	return nil
}

func settingIPCSocket(socketPath string) (net.Listener, error) {
	os.Remove(socketPath)

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, err
	}

	return l, nil
}

package process

import (
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

func settingOverlayFsMount(containerDirPath string) error {
	return mounts.OverlayFsMount(containerDirPath)
}

func settingPivotRoot(containerDirPath string) error {
	return mounts.PivotRoot(containerDirPath)
}

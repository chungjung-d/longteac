package mounts

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func PivotRoot(containerDirPath string) error {

	newRoot := containerDirPath + "/overlay" + "/merge"
	oldRoot := containerDirPath + "/overlay" + "/merge" + "/old_root"

	err := os.MkdirAll(oldRoot, 0755)
	if err != nil {
		return err
	}

	os.Chdir(newRoot)
	fmt.Println("PivotRoot", containerDirPath+"/rootfs", oldRoot)

	err = unix.PivotRoot(newRoot, oldRoot)
	if err != nil {
		return err
	}

	os.Chdir("/")

	err = unix.Unmount("/old_root", unix.MNT_DETACH)
	if err != nil {
		return err
	}

	err = os.Remove("/old_root")
	if err != nil {
		return err
	}

	return nil
}

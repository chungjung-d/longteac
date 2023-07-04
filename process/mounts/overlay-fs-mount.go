package mounts

import (
	"os"

	"golang.org/x/sys/unix"
)

func OverlayFsMount(containerDirPath string) error {

	os.MkdirAll(containerDirPath+"/overlay", 0755)

	os.MkdirAll(containerDirPath+"/overlay/work", 0755)
	os.MkdirAll(containerDirPath+"/overlay/container", 0755)
	os.MkdirAll(containerDirPath+"/overlay/merge", 0755)

	source := "overlay"
	target := containerDirPath + "/overlay/merge"
	fstype := "overlay"
	flags := 0
	data := "lowerdir=" + containerDirPath + "/rootfs,upperdir=" + containerDirPath + "/overlay/container,workdir=" + containerDirPath + "/overlay/work"

	err := unix.Mount(source, target, fstype, uintptr(flags), data)
	if err != nil {
		return err
	}

	return nil
}

package mounts

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/chungjung-d/longteac/config"
)

// TODO It not work on Bind mount (Bind mount will execute on longtea process)
func Mount(mount config.MountsConfig) error {

	if mount.Type == "bind" {
		prepareBindMount(mount)
	}
	if mount.Type != "bind" {
		prepareNonBindMount(mount)
	}
	mountFlag, mountOptions := getMountFlag(mount)
	return syscall.Mount(mount.Source, "rootfs"+mount.Destination, mount.Type, mountFlag, strings.Join(mountOptions, ","))
}

func getMountFlag(mount config.MountsConfig) (uintptr, []string) {

	fmt.Println("Mounting", mount.Source, "to", mount.Destination, "as", mount.Type)

	flag := 0
	var mountOptions []string

	if mount.Type == "bind" {
		flag = flag | syscall.MS_BIND
	}

	for _, option := range mount.Options {
		switch option {
		case "nosuid":
			flag = flag | syscall.MS_NOSUID
		case "nodev":
			flag = flag | syscall.MS_NODEV
		case "noexec":
			flag = flag | syscall.MS_NOEXEC
		case "ro":
			flag = flag | syscall.MS_RDONLY
		case "remount":
			flag = flag | syscall.MS_REMOUNT
		case "mand":
			flag = flag | syscall.MS_MANDLOCK
		case "dirsync":
			flag = flag | syscall.MS_DIRSYNC
		case "noatime":
			flag = flag | syscall.MS_NOATIME
		case "nodiratime":
			flag = flag | syscall.MS_NODIRATIME
		case "bind":
			flag = flag | syscall.MS_BIND
		case "rbind":
			flag = flag | syscall.MS_BIND | syscall.MS_REC
		case "unbindable":
			flag = flag | syscall.MS_UNBINDABLE
		case "runbindable":
			flag = flag | syscall.MS_UNBINDABLE | syscall.MS_REC
		case "private":
			flag = flag | syscall.MS_PRIVATE
		case "rprivate":
			flag = flag | syscall.MS_PRIVATE | syscall.MS_REC
		case "shared":
			flag = flag | syscall.MS_SHARED
		case "rshared":
			flag = flag | syscall.MS_SHARED | syscall.MS_REC
		case "slave":
			flag = flag | syscall.MS_SLAVE
		case "rslave":
			flag = flag | syscall.MS_SLAVE | syscall.MS_REC
		case "relatime":
			flag = flag | syscall.MS_RELATIME
		case "strictatime":
			flag = flag | syscall.MS_STRICTATIME
		case "active":
			flag = flag | syscall.MS_ACTIVE
		case "nouser":
			flag = flag | syscall.MS_NOUSER
		case "iversion":
			flag = flag | syscall.MS_I_VERSION
		default:
			mountOptions = append(mountOptions, option)

		}

	}

	return uintptr(flag), mountOptions
}

func prepareBindMount(mount config.MountsConfig) {

	sourceFile, _ := os.Stat(mount.Source)

	if sourceFile.IsDir() {
		os.MkdirAll("rootfs"+mount.Destination, os.ModePerm)
	} else {
		filenameSplit := strings.Split(mount.Destination, "/")
		filenameSplit = filenameSplit[:len(filenameSplit)-1]
		os.MkdirAll("rootfs/"+strings.Join(filenameSplit, "/"), os.ModePerm)
		os.Create("rootfs" + mount.Destination)
	}
}

func prepareNonBindMount(mount config.MountsConfig) {

	if _, err := os.Stat("rootfs" + mount.Destination); os.IsNotExist(err) {
		os.MkdirAll("rootfs"+mount.Destination, os.ModePerm)
	}
}

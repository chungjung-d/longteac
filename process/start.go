package process

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chungjung-d/longteac/config"
	ipchandler "github.com/chungjung-d/longteac/process/ipc-handler"
)

func StartContainerRootProcess(ctx context.Context) {
	fmt.Printf("New namespace setup complete, running as PID %d\n", os.Getpid())

	socketPath := ctx.Value(config.SocketPath).(string)
	containerDirPath := ctx.Value(config.ContainerDirPath).(string)

	err := os.Chdir(containerDirPath)
	if err != nil {
		fmt.Println("Error changing directory: ")
		log.Fatal(err)
	}

	ociConfig, err := config.GetOCIConfig(containerDirPath)
	if err != nil {
		fmt.Println("Error getting OCI config: ")
		log.Fatal(err)
	}

	err = settingMount(containerDirPath, ociConfig)
	if err != nil {
		fmt.Println("Error setting mount: ")
		log.Fatal(err)
	}

	listener, err := settingIPCSocket(socketPath)
	if err != nil {
		fmt.Println("Error setting IPC socket: ")
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ipchandler.HandleConnection(conn)
	}
}

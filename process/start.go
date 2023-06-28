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

	listener, err := settingIPCSocket(socketPath)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ipchandler.HandleConnection(conn)
	}
}

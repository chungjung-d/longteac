package process

import (
	"net"
	"os"
)

func settingIPCSocket(socketPath string) (net.Listener, error) {
	os.Remove(socketPath)

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, err
	}

	return l, nil
}

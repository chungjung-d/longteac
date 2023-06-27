package process

import (
	"fmt"
	"net"
	"os"
)

func CreateRootProcess() {
	fmt.Printf("New namespace setup complete, running as PID %d\n", os.Getpid())

	socketPath := "/tmp/mysocket.sock"
	os.Remove(socketPath)

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Here, you can add code to handle the received command and execute functions accordingly
	fmt.Printf("Reived %d bytes: %s\n", n, buf[:n])
}

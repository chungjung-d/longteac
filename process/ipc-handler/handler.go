package ipchandler

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
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

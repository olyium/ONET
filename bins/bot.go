package main

import (
	"fmt"
	"net"
)

var (
	IP   string = "127.0.0.1"
	PORT string = "271"
)

func main() {
	// still in works
	SERVER, _ := net.ResolveTCPAddr("TCP", IP+":"+PORT)
	CONNECTION, _ := net.DialTCP("TCP", nil, SERVER)

	for {
		RECEIVED := make([]byte, 1024)
		_, _ = CONNECTION.Read(RECEIVED)
		fmt.Println(RECEIVED)
	}

}

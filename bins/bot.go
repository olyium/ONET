package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	IP   string = "127.0.0.1"
	PORT string = "271"
)

func main() {

	CONNECTION, ERR := net.Dial("tcp", IP+":"+PORT)
	if ERR != nil {
		os.Exit(1)
	}

	ConnectionScanner := bufio.NewScanner(CONNECTION)
	for ConnectionScanner.Scan() {
		fmt.Println(ConnectionScanner.Text())
	}

}

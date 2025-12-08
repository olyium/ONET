package core

import (
	"fmt"
	"net"
)

func ConnectionHandlerBot(CONNECTION net.Conn) {
	fmt.Println("bot connected")
}

package core

import (
	"net"
)

func Help(CONNECTION net.Conn) string {

	HELP := "\033[33m\n1. .help - shows this menu?\n2. .methods - shows the methods onet provides\033[0m\r\n"
	return HELP

}

func Methods(CONNECTION net.Conn) string {

	METHODS := "\033[33m\n1. .udp - <ip> <port> <seconds> [udp flood method] \n2. .get <url> <seconds> [sends a very large amount of get requests to a site]\033[0m\r\n"
	return METHODS

}

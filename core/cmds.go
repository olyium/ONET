package core

import (
	"net"
)

func Help(CONNECTION net.Conn) string {

	return "\033[33m" + HELPSTRING + "\033[0m\n"

}

func Methods(CONNECTION net.Conn) string {

	return "\033[33m" + METHODSSTRING + "\033[0m\n"

}

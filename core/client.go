package core

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func ConnectionHandlerClient(CONNECTION net.Conn) {

}

func ConnectionHandlerClientLoggedIn(CONNECTION net.Conn) {
	OUTPUT := ""

	for {

		Writer := bufio.NewWriter(CONNECTION)
		BannerScanner := bufio.NewScanner(strings.NewReader(BANNER))
		UserInputScanner := bufio.NewScanner(CONNECTION)

		Writer.WriteString("\033[2J\033[H")

		for BannerScanner.Scan() {
			Writer.WriteString("\033[33m" + BannerScanner.Text() + "\033[0m\r\n")
		}

		Writer.WriteString("\n")
		Writer.WriteString(OUTPUT + "\n")
		Writer.Flush()
		Writer.WriteString(fmt.Sprintf(USERCLI, "hi") + "\033[?25l")
		Writer.Flush()

		for UserInputScanner.Scan() {

			INPUT := UserInputScanner.Text()

			if INPUT == ".help" {
				OUTPUT = Help(CONNECTION)
			}

			if INPUT == ".methods" {
				OUTPUT = Methods(CONNECTION)
			}
			break
		}

	}
}

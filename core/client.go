package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

type CLIENT struct {
	USERNAME string
	PASSWORD string
}

func ConnectionHandlerClient(CONNECTION net.Conn) {

	OUTPUT := ""

	for {

		Writer := bufio.NewWriter(CONNECTION)
		Reader := bufio.NewScanner(CONNECTION)
		BannerScanner := bufio.NewScanner(strings.NewReader(BANNER))

		Writer.WriteString("\033[2J\033[H")

		for BannerScanner.Scan() {
			Writer.WriteString("\033[33m" + BannerScanner.Text() + "\033[0m\r\n")
		}

		Writer.WriteString("\n")
		Writer.Flush()
		if OUTPUT != "" {
			Writer.WriteString(OUTPUT + "\n")
		}
		if _, ERR := os.Stat("./config/accounts.json"); ERR != nil {
			Writer.WriteString("\033[31mwe're experiencing difficulties retrieving our userbase.\033[0m")
			Writer.Flush()
			CONNECTION.Close()
		}

		Writer.WriteString("\r" + fmt.Sprintf(USERCLI, "login"))
		Writer.Flush()

		for Reader.Scan() {
			INPUT := Reader.Text()
			if len(strings.Split(INPUT, ":")) != 2 {
				OUTPUT = "\033[31m\rinvalid credentials format. do username:password\033[0m\n"
				break
			}
			if ValidateClient(CONNECTION, strings.Split(INPUT, ":")[0], strings.Split(INPUT, ":")[1]) {
				fmt.Println("\033[33m" + CONNECTION.RemoteAddr().String() + " logged in\033[0m")
				break
			} else {
				OUTPUT = "\033[31m\rinvalid credentials\033[0m\n"
				break
			}
		}

	}

}

func ConnectionHandlerClientLoggedIn(CONNECTION net.Conn, USERNAME string) {
	OUTPUT := ""

	for {

		Writer := bufio.NewWriter(CONNECTION)
		BannerScanner := bufio.NewScanner(strings.NewReader(BANNER))
		UserInputScanner := bufio.NewScanner(CONNECTION)

		Writer.WriteString("\033[2J\033[H")

		for BannerScanner.Scan() {
			Writer.WriteString("\033[33m" + BannerScanner.Text() + "\033[0m\r\n")
		}

		Writer.WriteString("\r\n")
		if OUTPUT != "" {
			Writer.WriteString(OUTPUT + "\r\n")
		}
		Writer.WriteString("\r" + fmt.Sprintf(USERCLI, USERNAME))
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

func ValidateClient(CONNECTION net.Conn, USERNAME string, PASSWORD string) bool {

	if _, ERR := os.Stat("./config/accounts.json"); ERR == nil {

		CONTENT, ERR := os.ReadFile("./config/accounts.json")
		if ERR != nil {
			fmt.Println("\033[31mfailed to read the accounts.json\033[0m")
			return false
		}

		var PAYLOAD []CLIENT
		ERR = json.Unmarshal(CONTENT, &PAYLOAD)
		if ERR != nil {
			fmt.Println("\033[31mfailed to read the json\033[0m")
			return false
		}

		for _, ACCOUNT := range PAYLOAD {
			if ACCOUNT.USERNAME == USERNAME && ACCOUNT.PASSWORD == PASSWORD {
				ConnectionHandlerClientLoggedIn(CONNECTION, USERNAME)
				return true
			}
		}

	}
	return false
}

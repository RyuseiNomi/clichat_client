package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "chat-server:8000", "http service address")

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/room"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial: ", err)
	}
	for {
		input := getStdin()
		if input == "exit" {
			c.Close()
			break
		}
	}
}

// getStdin ターミナルからの標準入力を受け取る
// "exit"を入力された場合にプログラムを終了する
func getStdin() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput := scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return stringInput
}

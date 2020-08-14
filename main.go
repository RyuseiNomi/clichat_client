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
		stringInput, byteInput := getStdin()
		if stringInput == "exit" {
			c.Close()
			break
		}
		c.WriteMessage(websocket.TextMessage, byteInput)
	}
}

// getStdin ターミナルからの標準入力を受け取る
// "exit"を入力された場合にプログラムを終了する
func getStdin() (string, []byte) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// stringはmain()で"exit"の判定を行うため。終了判定方法を変更した場合これは不要になる
	// byteはサーバへのメッセージ送信に利用するため
	stringInput := strings.TrimSpace(scanner.Text())
	byteInput := scanner.Bytes()

	return stringInput, byteInput
}

package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/RyuseiNomi/clichat_client/tracer"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "chat-server:8000", "http service address")

const messageBufferSize = 256

// user チャットに参加するユーザを表す
type user struct {
	input   chan []byte // ユーザが入力するメッセージを格納するチャネル
	receive chan []byte // サーバから受け取ったメッセージを格納するチャネル
	leave   chan bool
	tracer  tracer.Tracer
}

func main() {
	// ユーザ名の入力を受取
	log.Println("ユーザ名を入力：")
	username := getUserName()

	// URLにQueryStringを追加
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/room"}
	queryString := u.Query()
	queryString.Set("name", username)
	u.RawQuery = queryString.Encode()

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial: ", err)
	}

	usr := newUser()
	usr.tracer = tracer.New(os.Stdout)
	usr.run(c)
}

// run Clientが退出もしくはプログラムを終了するまで、以下の3つのチャネルを監視する
// 1. ユーザからの標準入力(メッセージ)
// 2. サーバからのWriteMessage
// 3. クライアントの退出フラグ("exit"という文字列)
func (u *user) run(c *websocket.Conn) {
	u.leave <- false
	for {
		go standByInput(u)
		go monitor(u, c)
		select {
		case input := <-u.input:
			// ターミナルからメッセージの入力を受けた時
			c.WriteMessage(websocket.TextMessage, input)
		case msg := <-u.receive:
			u.tracer.Trace(string(msg))
		case leave := <-u.leave:
			if leave == true {
				u.tracer.Trace("接続を終了します")
				c.Close()
				break
			}
		}
	}
}

// getUserName ユーザ名を受け取る
func getUserName() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	return name
}

// standByInput ユーザからのメッセージの入力を待機し、サーバに送信する
func standByInput(u *user) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// stringはmain()で"exit"の判定を行うため。終了判定方法を変更した場合これは不要になる
	stringInput := strings.TrimSpace(scanner.Text())
	if stringInput == "exit" {
		u.leave <- true
	} else {
		byteInput := scanner.Bytes()
		u.input <- byteInput
	}
}

// receive サーバからの受信を監視する
func monitor(u *user, c *websocket.Conn) {
	_, p, err := c.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	u.receive <- p
}

// newUser 空のuserモデルを返却する
func newUser() *user {
	return &user{
		input:   make(chan []byte, messageBufferSize),
		receive: make(chan []byte, messageBufferSize),
		leave:   make(chan bool, 1),
	}
}

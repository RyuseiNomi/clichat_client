package main

import (
	"flag"
	"log"
	"net/url"

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
	defer c.Close()
}

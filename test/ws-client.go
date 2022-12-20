package main

import (
	"fmt"
	"log"
	"github.com/tailscale/net/websocket"
)

func main() {
	origin := "http://100.99.127.71"
	url := "ws://100.99.14.179:80/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}
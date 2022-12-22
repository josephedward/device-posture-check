package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/tailscale/net/websocket"
	"godpc/cli"
	"io"
	"log"
	"net/http"
	// "tailscale.com/client/tailscale"
	"tailscale.com/tsnet"
)

var (
	hostname = flag.String("hostname", "PostureService", "hostname for the tailnet")
)

// // Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	//print out something very recognizable to the console
	cli.Success("server side")

	//return something very recognizable to the client
	io.WriteCloser(ws).Write([]byte("EchoServer"))
}

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	tsenv, err := cli.Env()
	cli.Success("tsenv : ", tsenv)
	cli.PrintIfErr(err)
	
	s := &tsnet.Server{
		Hostname: *hostname,
	}

	defer s.Close()

	ln, err := s.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	lc, err := s.LocalClient()
	if err != nil {
		log.Fatal(err)
	}
	cli.Success("lc : ", lc)

	http.Serve(ln, websocket.Handler(EchoServer))
	// http.Handle("/echo", websocket.Handler(EchoServer))

	// cli.Success(http.ListenAndServe(":12345", nil))

	// err = http.ListenAndServe(":12345", nil)
	// if err != nil {
	// 	panic("ListenAndServe: " + err.Error())
	// }
}

// // This example demonstrates a trivial echo server.
// func example() {
// 	zerolog.SetGlobalLevel(zerolog.DebugLevel)
// 	http.Handle("/echo", websocket.Handler(EchoServer))

// 	cli.Success(http.ListenAndServe(":12345", nil))

// 	err := http.ListenAndServe(":12345", nil)
// 	if err != nil {
// 		panic("ListenAndServe: " + err.Error())
// 	}

// }

// import (
// 	"websocket"
// 	"websocket/json"
// )

// type T struct {
// 	Msg string
// 	Count int
// }
// func main(){
// // create a new websocket
// ws, err := websocket.Dial("ws://localhost:12345/echo", "", "http://localhost/")

// // receive JSON type T
// var data T
// websocket.JSON.Receive(ws, &data)

// // send JSON type T
// websocket.JSON.Send(ws, data)

// }

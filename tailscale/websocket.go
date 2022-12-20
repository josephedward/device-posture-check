package tailscale 

import (
	"godpc/cli"
	"github.com/tailscale/net/websocket"
	"github.com/rs/zerolog"
	"tailscale.com/client/tailscale"
	"tailscale.com/tsnet"
	"flag"
	"log"
	"net/http"
	"io"
	// "html"
	// "github.com/tailscale/net/websocket"
	// "github.com/rs/zerolog"
	// "godpc/cli"
)

var (
	hostname = flag.String("hostname", "PostureService", "hostname for the tailnet")
)


// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func wsServer(statusQuery string, env cli.TsEnv) {
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

	http.Handle("/echo", websocket.Handler(EchoServer))
	
	cli.Success(http.ListenAndServe(":12345", nil))

	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}


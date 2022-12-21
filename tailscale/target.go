package tailscale

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/tailscale/net/websocket"
	"godpc/cli"
	"godpc/osq"
	"io"
	"net/http"
	"tailscale.com/tsnet"
	// "log"
	// "fmt"
	// "io/ioutil"
	// "strings"
)

var (
	tailnetServiceName = flag.String("hostname", "PostureService", "name of the service on the tailnet")
	// cqs                = osq.CurrentQueryStruct{}
)

// Echo the data received on the WebSocket.
func Echo(ws *websocket.Conn) {
	//print out something very recognizable to the console
	cli.Success("echo : ", ws)

	res := osq.GetResponse()
	//return something very recognizable to the client
	io.WriteCloser(ws).Write([]byte(res))
}

func WebSocketService(statusQuery string, env cli.TsEnv) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	tsenv, err := cli.Env()
	cli.Success("tsenv : ", tsenv)
	cli.PrintIfErr(err)

	s := &tsnet.Server{
		Hostname: *tailnetServiceName,
	}

	defer s.Close()

	ln, err := s.Listen("tcp", ":80")
	if err != nil {
		cli.Error(err)
	}

	defer ln.Close()

	lc, err := s.LocalClient()
	if err != nil {
		cli.Error(err)
	}
	cli.Success("lc : ", lc)

	osq.SetResponse(statusQuery)

	http.Serve(ln, websocket.Handler(Echo))
}

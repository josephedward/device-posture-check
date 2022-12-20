package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	
	"tailscale.com/tsnet"
	"github.com/rs/zerolog"
	"godpc/cli"
	//tailscale client
	
)

var (
	hostname = flag.String("hostname", "PostureService", "hostname for the tailnet")
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	
	statusQuery := cli.ReadFile("./current/response.json")

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


	log.Fatal(http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		who, err := lc.WhoIs(r.Context(), r.RemoteAddr)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprintf(w, "%s", string(statusQuery))

		//print IP address of the service node
		cli.Success("IP Address: ", html.EscapeString(r.Host))
		cli.WriteFile(".serviceip", html.EscapeString(r.Host))
	})))
}

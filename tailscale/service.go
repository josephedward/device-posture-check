package tailscale

import (
	"flag"
	"fmt"
	"godpc/cli"
	"html"
	"log"
	"net/http"
	"strings"
	"tailscale.com/tsnet"
)

var (
	hostname = flag.String("hostname", "hello", "hostname for the tailnet")
)

func CreateService(statusQuery string, env cli.TsEnv) {

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

	http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		who, err := lc.WhoIs(r.Context(), r.RemoteAddr)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// //print username
		fmt.Fprintf(w, "<p>Username: %s</p>", html.EscapeString(who.UserProfile.LoginName))
		fmt.Fprintf(w, "<p>Report for TS NODE @ IP ADDRESS : %s</p>", html.EscapeString(env.Ip))
		fmt.Fprintf(w, "<p>Query Results: %s</p>", strings.ReplaceAll(statusQuery, ",", ",<br/>"))

	}))

}

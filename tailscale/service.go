package tailscale

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"

	"tailscale.com/tsnet"
	"os/exec"
)

var (
	hostname = flag.String("hostname", "hello", "hostname for the tailnet")
)

func CreateService(statusQuery, tsIp string) {

	ip, err := exec.LookPath("tailscale ip -4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IP : ",ip)

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
		//print username
		fmt.Fprintf(w, "<p>Username: %s</p>", html.EscapeString(who.UserProfile.LoginName),)
		fmt.Fprintf(w, "<p>Report for TS NODE @ IP ADDRESS : %s</p>", html.EscapeString(tsIp),)
		fmt.Fprintf(w, "<p>Query Results: %s</p>", strings.ReplaceAll(statusQuery, ",", ",<br>"))

	})))
}

func firstLabel(s string) string {
	if hostname, _, ok := strings.Cut(s, "."); ok {
		return hostname
	}

	return s
}

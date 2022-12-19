package tailscale

import (
	"flag"
	"fmt"
	"godpc/cli"
	"html"
	"log"
	"net/http"
	"tailscale.com/tsnet"
	// "crypto/tls"
	// "os"
	// "strings"
	// "tailscale.com/client/tailscale"
)

var (
	hostname = flag.String("hostname", "PostureService", "hostname for the tailnet")
)

func CreateService(statusQuery string, env cli.TsEnv) string {

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

	serviceIp := ""

	http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		who, err := lc.WhoIs(r.Context(), r.RemoteAddr)
		cli.Success(who)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// //print username
		// fmt.Fprintf(w, "<p>Username: %s</p>", html.EscapeString(who.UserProfile.LoginName))
		// fmt.Fprintf(w, "<p>Report for TS NODE @ IP ADDRESS : %s</p>", html.EscapeString(env.Ip))
		fmt.Fprintf(w, "%s", string(statusQuery))

		//print IP address of the service node
		cli.Success("IP Address: ", html.EscapeString(r.Host))
		cli.WriteFile(".serviceip", html.EscapeString(r.Host))

	}))

	return serviceIp

}

// func CreateListener() {
// 	// os.Setenv("TAILSCALE_USE_WIP_CODE", "true")
// 	// TODO: comment out this line to avoid having to re-login each time you start this program
// 	// os.Setenv("TS_LOGIN", "1")

// 	os.Setenv("HOME", "./perm/postureservice")

// 	// hostname := flag.String("hostname", "josephedward.github", "tailscale hostname")
// 	allowedUser := flag.String("josephedward@github", "", "the name of a tailscale user to allow")
// 	flag.Parse()
// 	s := &tsnet.Server{
// 		Hostname: *hostname,
// 	}

// 	log.Printf("starting tailscale listener on hostname %s", *hostname)
// 	ln, err := s.Listen("tcp", ":443")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//create local client
// 	LocalClient, err := s.LocalClient()
// 	cli.PrintIfErr(err)

// 	ln = tls.NewListener(ln, &tls.Config{
// 		GetCertificate: LocalClient.GetCertificate,
// 	})

// 	defer ln.Close()

// 	http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		who, err := LocalClient.WhoIs(r.Context(), r.RemoteAddr)
// 		if err != nil {
// 			cli.Success("here")
// 			http.Error(w, err.Error(), 500)
// 			return
// 		}
// 		if who.UserProfile.LoginName != *allowedUser {
// 			cli.Success("here2")
// 			http.Error(w, "not allowed", 403)
// 			return
// 		}
// 		cli.Success("here3")
// 		fmt.Fprintf(w, "Hello, %s", html.EscapeString(who.UserProfile.LoginName))
// 	}))
// }

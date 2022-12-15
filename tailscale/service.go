package tailscale

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"godpc/cli"
	"log"

	"tailscale.com/tsnet"


	"html"
	"net/http"
	"strings"
	// "github.com/robertkrimen/otto"
	// v8 "rogchap.com/v8go"
	"os"
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
	fmt.Println(lc)

	cli.Success("tsenv.devId: ", env.ApiKey)
	sEnc := b64.StdEncoding.EncodeToString([]byte(env.DevId))
	cli.Success("base64 encoded : ", sEnc)
	authFunc := ReadJs("./scripts/authorize.js")

	test, err := fmt.Printf(authFunc, sEnc, env.DevId)
	fmt.Println(test)
	

	http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		who, err := lc.WhoIs(r.Context(), r.RemoteAddr)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		//print username
		fmt.Fprintf(w, "<p>Username: %s</p>", html.EscapeString(who.UserProfile.LoginName))
		fmt.Fprintf(w, "<p>Report for TS NODE @ IP ADDRESS : %s</p>", html.EscapeString(env.Ip))
		fmt.Fprintf(w, "<p>Query Results: %s</p>", strings.ReplaceAll(statusQuery, ",", ",<br>"))
		fmt.Fprintf(w, "<p>Query Results: %s</p>", strings.ReplaceAll(statusQuery, ",", ",<br>"))
	}))

}

func ReadJs(path string) string {
	//use the os package to read the file
	c, ioErr := os.ReadFile(path)
	js := string(c)
	cli.PrintIfErr(ioErr)
	return js
}





package main

import (
	// "encoding/json"
	// "context"
	"godpc/cli"
	// "fmt"
	// "github.com/hokaccha/go-prettyjson"
	// "github.com/tailscale/tailscale-client-go/tailscale"
	//tsnet
	// "github.com/tailscale/tailscale-client-go/tailscale/net"
	"godpc/osquery"
	"os"
	"github.com/rs/zerolog"
)	


func main() {
		//look through all os.Args and see if one is "debug"
		for _, arg := range os.Args {
			if arg == "prod" {
				zerolog.SetGlobalLevel(zerolog.InfoLevel)
				break
			}
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
	//welcome message
	cli.Welcome()
	//read the query
	cli.Success("Reading query")
	queryString := osquery.ReadQuery("query.sql")
	cli.Success("Query : "+queryString)
	//run the query
	//expose the results
	
}

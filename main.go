package main

import (
	"encoding/json"
	// "context"
	"godpc/cli"
	// "fmt"
	"github.com/hokaccha/go-prettyjson"
	// "github.com/tailscale/tailscale-client-go/tailscale"
	//tsnet
	// "github.com/tailscale/tailscale-client-go/tailscale/net"
	"godpc/osquery"
	"os"
	"github.com/rs/zerolog"
	// "fmt"
	// "log"
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
	cli.Success("Query : " + queryString)
	//run the query
	cli.Success("Running query")
	queryResponse := osquery.RunQuery("/var/osquery/osquery.em", queryString)
	cli.Success("Query Response : " + queryResponse)
	
	json, err := json.Marshal(queryResponse)
	cli.PrintIfErr(err)
	//pretty print the json
	pretty, err := prettyjson.Format(json)

	
	//expose the results
}

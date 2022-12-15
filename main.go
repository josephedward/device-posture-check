package main

import (
	"godpc/cli"
	"godpc/osquery"
	"godpc/tailscale"
	"os"
	"github.com/rs/zerolog"
	// "reflect"
	
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
	// cli.Welcome()


	tsenv, err := cli.Env()
	cli.Success("tsenv: ", tsenv)
	cli.PrintIfErr(err)

	//read the query
	cli.Success("Reading query")
	queryString := osquery.ReadQuery("query.sql")
	cli.Success("Query : " + queryString)
	//run the query
	cli.Success("Running query")
	queryResponse := osquery.RunQuery("/var/osquery/osquery.em", queryString)
	cli.Success("Query Response : ", queryResponse)
	// cli.Success("len(query)", len(queryResponse))

	//create the service
	cli.Success("Creating service")
	tailscale.CreateService(queryResponse,tsenv)


	}

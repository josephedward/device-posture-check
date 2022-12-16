package main

import (
	"godpc/cli"
	"godpc/osquery"
	// "godpc/tailscale"
	// "os"
	// "github.com/rs/zerolog"
	// "reflect"
	// "log"
	
	
)

var log = cli.ZeroLog()


func main() {
	bootstrap()
	query()
	// //create the service
	// cli.Success("Creating service")
	// tailscale.CreateService(queryResponse, tsenv)

}

func bootstrap(){
	cli.Welcome()
	tsenv, err := cli.Env()
	cli.Success("tsenv : ", tsenv)
	cli.PrintIfErr(err)
}

func query(){
	// read the query
	cli.Success("query : ")
	queryString := osquery.ReadQuery("query.sql")
	log.Info().Msg("query : " + queryString)
	//run the query
	cli.Success("run query")
	queryResponse := osquery.RunQuery("/var/osquery/osquery.em", queryString)
	cli.Success("query result : ", queryResponse)	
}


func service(){
	//create the service
	cli.Success("Creating service")
	// tailscale.CreateService(queryResponse, tsenv)
}
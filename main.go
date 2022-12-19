package main

import (
	"fmt"
	"godpc/cli"
	"godpc/osquery"
	"godpc/tailscale"
	"os"
)

var log = cli.ZeroLog()
var dpc *DevicePostureCheck

type DevicePostureCheck struct {
	tsenv cli.TsEnv
	osqst osquery.QueryStruct
}

func main() {
	dpc = &DevicePostureCheck{}

	dpc.tsenv = bootstrap()

	exit := false
	for !exit {
		Execute(dpc)
	}
}

func Execute(dpc *DevicePostureCheck) {

	options := []cli.PromptOptions{
		{
			Label: "Exit CLI",
			Key:   0,
		},
		{
			Label: "Display Current Query Information",
			Key:   1,
		},
		{
			Label: "Run New Query",
			Key:   2,
		},
		{
			Label: "Create Service",
			Key:   3,
		},
	}
	prompt := cli.Select("Welcome to GODPC - Please select an option: ", options)

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Option %d: %s\n", i+1, options[i].Label)

	switch options[i].Key {
	case 0:
		os.Exit(0)
	case 1:
		log.Info().Msg("Current Query : " + dpc.osqst.CurrentQuery)
		log.Info().Msg("Current Query Response String : " + dpc.osqst.CurrentQueryResponseStr)
	case 2:
		dpc.osqst = query()
	case 3:
		service(dpc.osqst.CurrentQueryResponseStr, dpc.tsenv)
	}

	Execute(dpc)
}

func bootstrap() cli.TsEnv {
	cli.Welcome()
	tsenv, err := cli.Env()
	cli.Success("tsenv : ", tsenv)
	cli.PrintIfErr(err)
	return tsenv
}

func query() osquery.QueryStruct {
	// read the query
	queryString := cli.PromptQuery()
	// cli.Success("query : ")
	// queryString := osquery.ReadQuery("query.sql")
	log.Info().Msg("query : " + queryString)
	//run the query
	cli.Success("run query")
	queryResponse := osquery.RunQuery("/var/osquery/osquery.em", queryString)
	log.Info().Msg("queryResponse" + queryResponse.CurrentQueryResponseStr)
	return queryResponse
}

func service(queryResponseStr string, tsenv cli.TsEnv) {
	//create the service
	cli.Success("Creating service")
	tailscale.CreateService(queryResponseStr, tsenv)
	// tailscale.CreateListener()
}

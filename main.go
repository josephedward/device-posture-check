package main

import (
	"fmt"
	"godpc/cli"
	"godpc/tailscale"
	"godpc/osq"
	"os"
)

var log = cli.ZeroLog()
var dpc *DevicePostureCheck

type DevicePostureCheck struct {
	tsenv cli.TsEnv
	osqst osq.QueryStruct
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
			Label: "Connect to Target Service Node",
			Key:   2,
		},
	}
	prompt := cli.Select("Please select an option: ", options)

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
		cli.Success("Current Query Information : ")
		cli.Success("Current Query : " + osq.GetQuery())
		cli.Success("Current Query Response String : " + osq.GetResponse())
	case 2:
		cli.Success("Connecting to Target Service Node : ")
		tailscale.ConnectService()

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

func query() osq.QueryStruct {
	// read the query
	queryString := cli.PromptQuery()
	// cli.Success("query : ")
	// queryString := osq.ReadQuery("query.sql")
	log.Info().Msg("query : " + queryString)
	//run the query
	cli.Success("run query")
	queryResponse := osq.RunQuery("/var/osquery/osquery.em", queryString)
	log.Info().Msg("queryResponse" + queryResponse.CurrentQueryResponseStr)
	return queryResponse
}

func service(queryResponseStr string, tsenv cli.TsEnv) {

	//create the service
	// cli.Success("Creating service")
	// tailscale.CreateService(queryResponseStr, tsenv)
	// tailscale.CreateListener()
}

// func visitServiceNode() {
// 	serviceIp := cli.ReadFile(".serviceip")
// 	cli.Success("serviceIp : ", serviceIp)
// 	deviceIp := cli.ReadFile(".deviceip")
// 	cli.Success("deviceIp : ", deviceIp)
// 	osq.CheckPosture(serviceIp, deviceIp)
// }

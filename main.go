package main

import (
	"fmt"
	"github.com/tailscale/net/websocket"
	"godpc/cli"
	"godpc/osq"
	"godpc/tailscale"
	"os"
	"os/exec"
	// "time"
)

var log = cli.ZeroLog()
var dpc *DevicePostureCheck

type DevicePostureCheck struct {
	tsenv  cli.TsEnv
	osqst  osq.QueryStruct
	wsconn *websocket.Conn
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
			Label: "Display Last Query Information",
			Key:   1,
		},
		{
			Label: "Connect to Target Service Node",
			Key:   2,
		},
		{
			Label: "Disconnect from Target Service Node",
			Key:   3,
		},
		{
			Label: "Run a New Query",
			Key:   4,
		},
		{
			Label: "Scripted Websocket Connection",
			Key:   5,
		},
		{
			Label: "Scripted Osquery Execution",
			Key:   6,
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
		dpc.displayQuery()
	case 2:
		dpc.wsconn = dpc.connectTarget()
	case 3:
		dpc.disconnectTarget()
	case 4:
		queryStr, err := cli.PromptString("Enter a query to run: ")
		cli.PrintIfErr(err)
		dpc.osqst = *dpc.newQuery(queryStr)
		
	case 5:
		dpc.scriptWebsocketConnection()
	case 6:
		dpc.osqst = *dpc.newQuery("select * from users;")
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

func (dpc *DevicePostureCheck) displayQuery() {
	query := osq.GetCurrentQueryStruct()
	cli.Success("current query : ", query.CurrentQuery)
	cli.Success("current query response : ", query.CurrentQueryResponseStr)
}

func (dpc *DevicePostureCheck) newQuery(queryString string) *osq.QueryStruct {
	query := osq.RunQuery("/var/osquery/osquery.em", queryString)
	cli.Success("query : ", query)
	return &query
}

func (dpc *DevicePostureCheck) connectTarget() *websocket.Conn {
	cli.Success("Connecting to Target Service Node...")
	originIp, err := cli.PromptString("Enter the IP address of origin ")
	cli.PrintIfErr(err)
	serviceIp, err := cli.PromptString("Enter the IP address of service ")
	cli.PrintIfErr(err)
	socketConnect := tailscale.ConnectWebSocket(originIp, serviceIp)
	return socketConnect
}

func (dpc *DevicePostureCheck) disconnectTarget() {
	cli.Success("Disconnecting from Target Service Node...")
	tailscale.DisconnectWebSocket(dpc.wsconn)
	cli.Success("Disconnected! ")
}

func (dpc *DevicePostureCheck) scriptWebsocketConnection() {
	//execute the /test/osquery.go script
	exec.Command("go run /test/osquery.go")

}

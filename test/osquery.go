package main

import (
	"encoding/json"
	"godpc/cli"
	"os"
	"time"
	"github.com/osquery/osquery-go"
	"github.com/rs/zerolog"
	// "strings"
	// "fmt"
	// "log"
)

func main() {
	//set current log level to debug
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	socketPath := ""
	query := ""
	outPath := ""

	//if there is no argument, use the default socket path
	if len(os.Args) < 2 {
		cli.Error("Using default socket path: %s", "/var/osquery/osquery.em")
		socketPath = "/var/osquery/osquery.em"
	} else {
		socketPath = os.Args[1]
	}
	client, err := osquery.NewClient(socketPath, 10*time.Second)

	if err != nil {
		cli.Error("Error creating Thrift client: %v", err)
	}
	defer client.Close()

	//if there is no argument, use the default query
	if len(os.Args) < 3 {
		cli.Error("Using default query in dir: %s", "./current/query.sql")
		query = cli.ReadFile("./current/query.sql")
	} else {
		query = os.Args[2]
	}

	resp, err := client.Query(query)
	if err != nil {
		cli.Error("Error communicating with osqueryd: %v", err)
	}
	if resp.Status.Code != 0 {
		cli.Error("osqueryd returned error: %s", resp.Status.Message)
	}

	cli.Success("Got results:\n%#v\n", resp.Response)

	if len(os.Args) < 4 {
		cli.Error("Writing query response to default location: %s", "./current/response.json")
		outPath = "./current/response.json"
	} else {
		outPath = os.Args[3]
	}
	f, err := os.Create(outPath)
	cli.PrintIfErr(err)
	response, err := json.Marshal(resp.Response)
	cli.PrintIfErr(err)
	l, err := f.Write([]byte(string(response)))
	cli.Success(l)
	cli.PrintIfErr(err)
	f.Close()

}

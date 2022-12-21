package osq

import (
	"encoding/json"
	"godpc/cli"
	"log"
	"strings"
	"time"
	"github.com/osquery/osquery-go"
	// "os"
	// "tailscale.com/client/tailscale"
	// "github.com/rs/zerolog"
	// "io"
	// "fmt"
	// "context"
	// "crypto/tls"
	// "fmt"
	// "net/http"


)

type QueryStruct struct {
	CurrentQuery            string
	CurrentQueryResponseMap map[string]interface{}
	CurrentQueryResponseStr string
}

var CurrentQueryStruct = &QueryStruct{}


func RunQuery(socketPath string, queryString string) QueryStruct {

	if socketPath == "" {
		log.Fatalf("Usage: %s SOCKET_PATH QUERY", socketPath)
	}

	client, err := osquery.NewClient(socketPath, 10*time.Second)
	if err != nil {
		log.Fatalf("Error creating Thrift client: %v", err)
	}
	defer client.Close()

	if queryString == "" {
		log.Fatalf("Bad Query: %s", queryString)
	}

	CurrentQueryStruct.CurrentQuery = strings.Trim(queryString, " ")

	resp, err := client.Query(queryString)
	if err != nil {
		log.Fatalf("Error communicating with osqueryd: %v", err)
	}
	if resp.Status.Code != 0 {
		log.Fatalf("osqueryd returned error: %s", resp.Status.Message)
	}

	response, err := json.Marshal(resp.Response)
	cli.PrintIfErr(err)
	CurrentQueryStruct.CurrentQueryResponseStr = string(response)
	cli.Success("returned json : " + CurrentQueryStruct.CurrentQueryResponseStr)
	CurrentQueryStruct.CurrentQueryResponseMap = QueryObject(CurrentQueryStruct.CurrentQueryResponseStr)
	return *CurrentQueryStruct
}

func QueryObject(queryResponse string) map[string]interface{} {
	// declare map and unmarshal json into it
	var myStoredVariable map[string]interface{}
	json.Unmarshal([]byte(queryResponse), &myStoredVariable)
	return myStoredVariable
}

func GetCurrentQueryStruct() QueryStruct {
	CurrentQueryStruct.CurrentQuery = GetQuery()
	CurrentQueryStruct.CurrentQueryResponseStr = GetResponse()
	CurrentQueryStruct.CurrentQueryResponseMap = QueryObject(CurrentQueryStruct.CurrentQueryResponseStr)
	return *CurrentQueryStruct
}


func GetQuery() string{
	return cli.ReadFile("./current/query.sql")
}

func GetResponse()string{
	return cli.ReadFile("./current/response.json")
}

func SetQuery(query string) {
	cli.WriteFile("./current/query.sql", query)
}

func SetResponse(response string) {
	cli.WriteFile("./current/response.json", response)
}

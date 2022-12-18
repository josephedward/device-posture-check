package osquery

import (
	// "fmt"
	"encoding/json"
	"godpc/cli"
	"log"
	"os"
	"strings"
	"time"

	"github.com/osquery/osquery-go"
)

type QueryStruct struct {
	CurrentQuery            string
	CurrentQueryResponseMap map[string]interface{}
	CurrentQueryResponseStr string
}

var CurrentQueryStruct = &QueryStruct{}

func ReadQuery(path string) string {
	//use the os package to read the file
	c, ioErr := os.ReadFile(path)
	sql := string(c)
	cli.PrintIfErr(ioErr)
	return sql
}

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


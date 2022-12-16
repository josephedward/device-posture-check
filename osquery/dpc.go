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

func ReadQuery(path string) string {
	//use the os package to read the file
	c, ioErr := os.ReadFile(path)
	sql := string(c)
	cli.PrintIfErr(ioErr)
	return sql
}

func RunQuery(socketPath string, queryString string) map[string]interface{}{

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
	resp, err := client.Query(queryString)
	if err != nil {
		log.Fatalf("Error communicating with osqueryd: %v", err)
	}
	if resp.Status.Code != 0 {
		log.Fatalf("osqueryd returned error: %s", resp.Status.Message)
	}

	response, err := json.Marshal(resp.Response)
	cli.PrintIfErr(err)
	strJson := strings.Trim(string(response), "[]")
	cli.Success("returned json : " + strJson)

	//declare map and unmarshal json into it
	var myStoredVariable map[string]interface{}
	json.Unmarshal([]byte(strJson), &myStoredVariable)
	cli.Success("myStoredVariable : " + myStoredVariable["name"].(string))

	return myStoredVariable
}

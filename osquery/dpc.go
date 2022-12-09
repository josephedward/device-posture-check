package osquery

import (
	// "fmt"
	"godpc/cli"
	"log"
	"os"
	"time"
	"github.com/osquery/osquery-go"
	"encoding/json"
)

func ReadQuery(path string) string{
	//use the os package to read the file
	c, ioErr := os.ReadFile(path)
	sql := string(c)
	cli.PrintIfErr(ioErr)
	return sql
}

func RunQuery(socketPath string, queryString string) string{

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
		log.Fatalf("Error communicating with osqueryd: %v",err)
	}
	if resp.Status.Code != 0 {
		log.Fatalf("osqueryd returned error: %s", resp.Status.Message)
	}

	cli.Success("Query Response", resp.Response)
	// return fmt.Sprintf("%#v", resp.Response)

	//send the response as json 
	json, err := json.Marshal(resp.Response)
	cli.PrintIfErr(err)
	return string(json)
}


package main

import (
	"fmt"
	"os"
	"time"
	"log"
	"github.com/osquery/osquery-go"
)

func main() {
	// fmt.Println("len(os.Args) : ", len(os.Args))
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s SOCKET_PATH QUERY", os.Args[0])
	}

	client, err := osquery.NewClient(os.Args[1], 10*time.Second)
	if err != nil {
		log.Fatalf("Error creating Thrift client: %v", err)
	}
	defer client.Close()

	resp, err := client.Query(os.Args[2])
	if err != nil {
		log.Fatalf("Error communicating with osqueryd: %v",err)
	}
	if resp.Status.Code != 0 {
		log.Fatalf("osqueryd returned error: %s", resp.Status.Message)
	}

	fmt.Printf("Got results:\n%#v\n", resp.Response)
}
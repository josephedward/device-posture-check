package main

import (
	"fmt"
	"github.com/osquery/osquery-go"
	"log"
	"time"
)

func main() {
	socketPath := "/var/osquery/osquery.em"
	client, err := osquery.NewClient(socketPath, 10*time.Second)
	if err != nil {
		log.Fatalf("Error creating Thrift client: %v", err)
	}
	defer client.Close()

	resp, err := client.Query("select * from users;")
	if err != nil {
		log.Fatalf("Error communicating with osqueryd: %v", err)
	}
	if resp.Status.Code != 0 {
		log.Fatalf("osqueryd returned error: %s", resp.Status.Code)
	}
	fmt.Printf("Got results:\n%v\n", resp.Response)
}

package main

import (
	"context"
	"fmt"
	"github.com/tailscale/tailscale-client-go/tailscale"
	"godpc/cli"
	"github.com/hokaccha/go-prettyjson"
)

func main() {

	tsenv, err := Env()
	cli.Success("tsenv: ", tsenv)
	cli.PrintIfErr(err)


	client, err := tailscale.NewClient(tsenv.apiKey, tsenv.tailnet)
	cli.PrintIfErr(err)

	// List all your devices
	devices, err := client.Devices(context.Background())
	cli.Success("devices :",devices)
	// loop over devices
	// for _, device := range devices {
	// 	fmt.Println(device)
	// }
	cli.PrintIfErr(err)
	s, _ := prettyjson.Marshal(devices)
	fmt.Println(string(s))
}
  
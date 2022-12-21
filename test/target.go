package main

import (
	"godpc/cli"
	"godpc/tailscale"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	tsenv, err := cli.Env()
	cli.PrintIfErr(err)
	currentResponse := cli.ReadFile("./current/response.json")
	tailscale.WebSocketService(currentResponse, tsenv)

}
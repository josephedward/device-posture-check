package main

import (
// .env import
// "github.com/joho/godotenv"
//tailscale import
// "tailscale.com/client/tailscale"
"fmt"

)

func main() {
// load .env file
tsenv, err := Env()
fmt.Println(tsenv)
if err != nil {
fmt.Println("Error loading .env file")
}

// tailscale - works
// curl 'https://api.tailscale.com/api/v2/tailnet/josephedward.github/devices' \
// -u "tskey-api-kfcWnv4CNTRL-nTmfmpo7zPfnHBuTgLTSTfBs4qZqKx2i"
}
#!/bin/bash

touch .tsip
tailscale ip -4 > .tsip
cat .tsip
tailscale status --json | jq -r '.Self.UserID' > .devid
cat .devid
task osquery
cd current && sudo caddy file-server --domain=$(cat ../.tsip) --browse
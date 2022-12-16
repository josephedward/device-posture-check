#!/bin/bash

echo TSIP=\"$(tailscale ip -4)\" >> .env
echo $TSIP
echo "y" | sudo apt install jq
echo "y" | sudo apt install gcc
echo "y" | sudo apt install g++
echo DEVID=\"$(tailscale status --json | jq -r '.Self.UserID'  )\" >> .env
echo $DEVID



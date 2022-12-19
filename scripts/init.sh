#!/bin/bash
TSIP=$(tailscale ip -4) 
echo TSIP : $TSIP
echo TSIP=\"$TSIP\" >> .env
# export TSIP : $TSIP

DEVID=$(tailscale status --json | jq -r '.Self.UserID')
echo DEVID : $DEVID
echo DEVID=\"$DEVID\" >> .env
# export TSIP : $TSIP

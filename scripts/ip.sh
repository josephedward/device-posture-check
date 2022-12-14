tailscale ip -4
tailscale status --json | jq -r '.Self.UserID'  
# tailscale status --json | jq -r '.Self.Hostinfo.Hostname'
# You can use the helper script:
sudo osqueryctl start

# Or, install the example config and launch daemon yourself:
sudo cp /var/osquery/osquery.example.conf /var/osquery/osquery.conf
sudo cp /var/osquery/io.osquery.agent.plist /Library/LaunchDaemons
sudo launchctl load /Library/LaunchDaemons/io.osquery.agent.plist
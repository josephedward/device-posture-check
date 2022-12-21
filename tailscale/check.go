package tailscale

import (
	"fmt"
	"godpc/cli"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tailscale/net/websocket"
	// "fmt"
	// "flag"
	// "io"
	// "log"
)

// needs to hit the tailnet for all devices
// find unauthed devices
// func FindUnauthedDevices(apiKey string) error {
// url := "https://api.tailscale.com/api/v2/devices"
// method := "GET"
func AuthorizeDevice(deviceID, apiKey string) error {

	//interpolate url
	url := "https://api.tailscale.com/api/v2/device/" + deviceID + "/authorized"
	method := "POST"

	payload := strings.NewReader(`{"authorized": true}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	//interpolate api key in header
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return err
}

// WIP - not working
// at a minimum, prompt for the origin (IP address of this device) and the URL (the IP address of the PostureService on the tailnet)
// origin, url string
func ConnectService() string {
	origin := "http://100.99.211.98"
	url := "ws://100.80.65.250:8888/"
	ws, err := websocket.Dial(url, "", origin)
	cli.PrintIfErr(err)
	// if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
	// 	cli.Error(err)
	// }
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		cli.PrintIfErr(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])

	return string(msg[:n])
}

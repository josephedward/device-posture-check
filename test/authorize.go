package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	//deviceID is first arg
	deviceID := os.Args[1]

	//api key is second arg
	apiKey := os.Args[2]

	//interpolate url
	url := "https://api.tailscale.com/api/v2/device/" + deviceID + "/authorized"
	method := "POST"
  println(url)

	payload := strings.NewReader(`{"authorized": true}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
  key := "Bearer "+apiKey
  fmt.Println(key)
  // os.Exit(1)

	//interpolate api key in header
	req.Header.Add("Authorization", key)
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

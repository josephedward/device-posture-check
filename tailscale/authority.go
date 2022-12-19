package tailscale

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

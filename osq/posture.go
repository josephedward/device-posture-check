package osq

import (

	"godpc/cli"
	"os"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"github.com/rs/zerolog"
	"tailscale.com/client/tailscale"
	"io"
	// "github.com/osquery/osquery-go"
	// "fmt"
	// "encoding/json"
	// "log"
	// "strings"
	// "time"

)

func CheckPosture(serviceIp string, deviceId string) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	serviceIP := ""
	deviceID := ""
	compliant := false
	tsenv, err := cli.Env()
	cli.PrintIfErr(err)

	//create tailscale client
	tsClient := tailscale.NewClient(tsenv.Tailnet, tailscale.APIKey(tsenv.ApiKey))

	cli.Success("tsClient : ", tsClient)
	ctx := context.Background()
	cli.Success("ctx : ", ctx)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//test if there is an argument
	if len(os.Args) < 2 {
		fmt.Println("No Service IP provided")
		os.Exit(1)
	} else {
		//first arg is ip address of tailscale device
		serviceIP = os.Args[1]
		cli.Success("serviceIP : ", serviceIP)
	}
	if len(os.Args) < 3 {
	fmt.Println("No Service IP provided")
	os.Exit(1)
	} else {
		// second arg is ip address of tailscale device
		deviceID = os.Args[2]
		cli.Success("deviceID : ", deviceID)
	}

	//visit url of tailscale device node ie
	visitNode := "http://" + serviceIP
	// + "/response.json"
	cli.Success("Visiting status service : ", visitNode)
	var client http.Client
	resp, err := client.Get(visitNode)
	cli.PrintIfErr(err)
	bodyString := ""
	// cli.Success("resp : ", resp.Body)
	// defer resp.Body.Close()
	
	// check if response is empty
	if resp == nil {
		fmt.Println("Response from ", visitNode, " is empty")
		os.Exit(1)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		cli.PrintIfErr(err)
		bodyString = string(bodyBytes)
	}

	cli.Success("response body string : ", bodyString)
	//check if response is not empty
	if resp != nil {
		fmt.Println("Response is not empty")
	}

	// //it is a specific length, then the device is compliant
	compliant = len(bodyString) > 1
	cli.Success("compliant : ", compliant)

	//if compliant, authorize
	if compliant == true {
		fmt.Println("Device is compliant")
		//use tailscale client to authorize device
		tsClient.AuthorizeDevice(ctx, deviceID)
		// tsClient.SetDeviceStatus(ctx, deviceID, "authorized")
		cli.Success("Device authorized")
		os.Exit(0)
	}

	// if not compliant, delete
	// if compliant == false {
	// 	fmt.Println("Device is not compliant")
	// 	//use tailscale client to delete device
	// 	tsClient.DeleteDevice(ctx, deviceID)
	// 	cli.Success("Device deleted")
	// 	os.Exit(0)
	// }

}


// /**
//  * SECOND HALF OF THE DPC PROCESS - needs to be separate app 
//  */
// func main() {

// 	tsenv, err := Env()
// 	cli.Success("tsenv: ", tsenv)
// 	cli.PrintIfErr(err)


// 	client, err := tailscale.NewClient(tsenv.apiKey, tsenv.tailnet)
// 	cli.PrintIfErr(err)

// 	// List all your devices
// 	devices, err := client.Devices(context.Background())
// 	cli.Success("device :",devices[0])
// 	// loop over devices
// 	for _, device := range devices {
// 		fmt.Println(device.Name)
// 	}
// 	cli.PrintIfErr(err)


	
// }

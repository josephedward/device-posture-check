
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

// 	// s, _ := prettyjson.Marshal(devices)
// 	// fmt.Println(string(s))	
	
// 	//open an ssh connection to a device
// 	conn, err := client.Dial(context.Background(), "device-name")
// 	cli.PrintIfErr(err)
// 	defer conn.Close()

// 	//run a command on the device
// 	cmd := exec.Command("ls", "-l")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.Stdin = conn
// 	err = cmd.Run()
// 	cli.PrintIfErr(err)

// 	//get the device's current status
// 	status, err := client.Status(context.Background())
// 	cli.PrintIfErr(err)
// 	fmt.Println(status)

	
// }

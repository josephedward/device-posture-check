package cli

import (
	"os"

	"github.com/joho/godotenv"
)

type TsEnv struct {
	ApiKey  string
	Tailnet string
	Ip      string
	DevId   string
}

func Env() (env TsEnv, err error) {
	err = godotenv.Load("./.env")
	apiKey := os.Getenv("APIKEY")
	tailnet := os.Getenv("TAILNET")
	tsip := os.Getenv("TSIP")
	devid := os.Getenv("DEVID")
	os.Setenv("TS_AUTHKEY", apiKey)
	os.Setenv("TS_NET", tailnet)
	os.Setenv("TS_IP", tsip)
	os.Setenv("TS_DEVID", devid)
	return TsEnv{apiKey, tailnet, tsip, devid}, err
}

func ReadFile(path string) string {
	//use the os package to read the file
	c, ioErr := os.ReadFile(path)
	readStr := string(c)
	PrintIfErr(ioErr)
	return readStr
}

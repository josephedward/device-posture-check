package cli

import (
	"os"
	"github.com/joho/godotenv"
)

type TsEnv struct {
	ApiKey          string
	Tailnet         string
	Ip			    string
}


func Env() (env TsEnv, err error){
	err = godotenv.Load("./.env")
	apiKey := os.Getenv("APIKEY")
	tailnet := os.Getenv("TAILNET")
	tsip := os.Getenv("TSIP")
	os.Setenv("TS_AUTHKEY", apiKey)
	os.Setenv("TS_NET", tailnet)
	os.Setenv("TS_IP", tsip)
	return TsEnv{apiKey, tailnet, tsip}, err
}
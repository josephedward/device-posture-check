package main

import (
	// "fmt"
	"os"

	"github.com/joho/godotenv"
)

type TsEnv struct {
	apiKey          string
	tailnet         string
}


func Env() (env TsEnv, err error){
	err = godotenv.Load("./.env")
	apiKey := os.Getenv("APIKEY")
	tailnet := os.Getenv("TAILNET")
	return TsEnv{apiKey, tailnet}, err
}

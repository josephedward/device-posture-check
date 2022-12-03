package main

import (
	"github.com/joho/godotenv"
	"os"
	// requests import


)

type TsEnv struct {
	API          string
}


func Env() (login TsEnv, err error){
	err = godotenv.Load("./.env")
	API := os.Getenv("TSAPI")
	return TsEnv{API}, err
}

// func LoadEnv() (login AwsEnv, err error) {
// 	//load env variables
// 	err = godotenv.Load("./.env")
// 	env := Env()
// 	return env, err
// }

// func LoadEnvPath(path string) (login AwsEnv, err error) {
// 	//load env variables
// 	err = godotenv.Load(path)
// 	env := Env()
// 	return env, err
// }

// func ArgEnv() (login AwsEnv, err error) {
// 	//set all needed vendor credentials from cli arg
// 	url := os.Args[1]
// 	username := os.Args[2]
// 	password := os.Args[3]
// 	key_id := os.Args[4]
// 	access_key := os.Args[5]

// 	return AwsEnv{url, username, password, key_id, access_key}, err
// }


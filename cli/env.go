package cli

import (
	"github.com/joho/godotenv"
	// "github.com/rs/zerolog"
	// "io/ioutil"
	"os"
	// "fmt"
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

func SetEnv(key string, val string) {
	os.Setenv(key, val)
}

func ReadFile(path string) string {
	//use the os package to read the file
	c, ioErr := os.ReadFile(path)
	readStr := string(c)
	PrintIfErr(ioErr)
	return readStr
}

func WriteFile(path string, content string) {
	//use the os package to write the file
	ioErr := os.WriteFile(path, []byte(content), 0644)
	PrintIfErr(ioErr)
}

// func fileOutput(fileName string) {
// 	// create a temp file
// 	tempFile, err := ioutil.TempFile("./logs", fileName)
// 	if err != nil {
// 		// Can we log an error before we have our logger? :)
// 		log.Error().Err(err).Msg("there was an error creating a temporary file for our log")
// 	}
// 	fileLogger := zerolog.New(tempFile).With().Logger()
// 	fileLogger.Info().Msg("This is an entry from my log")
// 	log.Error().Err(err).Msg("this is a test error message")

// }

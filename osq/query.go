package osq

import (
	"encoding/json"
	"github.com/osquery/osquery-go"
	"godpc/cli"
	"strings"
	"time"
)

type QueryStruct struct {
	CurrentQuery            string
	CurrentQueryResponseMap map[string]interface{}
	CurrentQueryResponseStr string
}

var CurrentQueryStruct = &QueryStruct{}

func RunQuery(socketPath string, queryString string) QueryStruct {
	if socketPath == "" {
		cli.Error("bad SOCKET_PATH QUERY : ", socketPath)
	}
	client, err := osquery.NewClient(socketPath, 10*time.Second)
	if err != nil {
		cli.Error("Error creating Thrift client : ", err)
	}
	defer client.Close()
	if queryString == "" {
		cli.Error("Bad Query : ", queryString)
	}

	CurrentQueryStruct.CurrentQuery = strings.Trim(queryString, " ")
	cli.Success("Running query : " + CurrentQueryStruct.CurrentQuery)
	resp, err := client.Query(CurrentQueryStruct.CurrentQuery)
	if err != nil {
		cli.Error("Error communicating with osqueryd : ", err)
	}
	if resp.Status.Code != 0 {
		cli.Error("osqueryd returned error : ", resp.Status.Message)
	}

	response, err := json.Marshal(resp.Response)
	cli.PrintIfErr(err)
	CurrentQueryStruct.CurrentQueryResponseStr = string(response)
	cli.Success("returned json : " + CurrentQueryStruct.CurrentQueryResponseStr)
	CurrentQueryStruct.CurrentQueryResponseMap = QueryObject(CurrentQueryStruct.CurrentQueryResponseStr)
	return *CurrentQueryStruct
}

func QueryObject(queryResponse string) map[string]interface{} {
	var myStoredVariable map[string]interface{}
	json.Unmarshal([]byte(queryResponse), &myStoredVariable)
	return myStoredVariable
}

func GetCurrentQueryStruct() QueryStruct {
	CurrentQueryStruct.CurrentQuery = GetQuery()
	CurrentQueryStruct.CurrentQueryResponseStr = GetResponse()
	CurrentQueryStruct.CurrentQueryResponseMap = QueryObject(CurrentQueryStruct.CurrentQueryResponseStr)
	return *CurrentQueryStruct
}

func GetQuery() string {
	return cli.ReadFile("./current/query.sql")
}

func GetResponse() string {
	return cli.ReadFile("./current/response.json")
}

func SetQuery(query string) {
	cli.WriteFile("./current/query.sql", query)
}

func SetResponse(response string) {
	cli.WriteFile("./current/response.json", response)
}

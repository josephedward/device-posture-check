.DEFAULT_GOAL := testrun

# BIN_FILE=myprogram

# buildandrun:
#  @go build -o "${BIN_FILE}"
#  ./"${BIN_FILE}"

testrun:
	@go run .
 
server:
	@go run server.go

osquery:
	@go run osquery.go
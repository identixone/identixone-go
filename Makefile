
test:
	go test -coverprofile=cover.txt ./...

html:
	go tool cover -html=cover.txt -o cover.html
all:
	go install -v ./...
	go test -c -v ./inventory
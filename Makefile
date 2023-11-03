.DEFAULT_GOAL := build 

clean:
	rm -f travtools

build: clean
	GOARCH=amd64 GOOS=linux go build -o travtools *.go

run:
	go run *.go
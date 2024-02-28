.DEFAULT_GOAL := build 

clean:
	rm -f travtools
	rm -f travtools.linux

build: clean
	templ generate
	go build -o travtools cmd/web/*.go

linuxbuild: clean
	templ generate
	GOARCH=amd64 GOOS=linux go build -o travtools.linux cmd/web/*.go

run:
	go run pkg/web/*.go
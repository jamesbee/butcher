.DEFAULT_GOAL := mac

all: clean mac linux

mac: main.go
	gox -tags="prod" -ldflags "-s -w" -osarch="darwin/amd64" -output="./bin/butcher"

linux: main.go
	gox -tags="prod" -ldflags "-s -w" -osarch="linux/amd64" -output="./bin/butcher-{{.OS}}-{{.Arch}}"

clean:
	rm -f ./bin/*
BINARY_SERVER = server
BINARY_AGENT = agent

default: build

build: build_server build_agent

build_server:
	go build -o ${BINARY_SERVER} cmd/server/main.go

build_agent:
	go build -o ${BINARY_AGENT} cmd/agent/main.go

test_server:
	go test ./... -timeout 60m --tags=server -v

test_agent:
	go test ./... -timeout 60m --tags=agent -v

clean:
	go clean
	rm -f $(BINARY_SERVER) $(BINARY_AGENT)

lint:
	go fmt ./...
	goimports -w $(find . -name '*.go')
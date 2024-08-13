.DEFAULT_GOAL := build

fmt:
		go fmt ./...

vet: fmt
		go vet ./...

clean:
		go clean

test: vet
		go test -v -cover ./...

build: test
		go build

run: build
		./url_shortener.exe

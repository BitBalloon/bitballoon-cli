.PHONY: fmt

build:
	go build -o bitballoon

fmt:
	gofmt -s -w -l bitballoon.go create.go deploy.go main.go update.go

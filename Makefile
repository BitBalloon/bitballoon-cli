.PHONY: all godep fmt

all: godep
	godep go install

godep:
	go get github.com/tools/godep

fmt:
	gofmt -s -w -l bitballoon.go create.go deploy.go main.go update.go

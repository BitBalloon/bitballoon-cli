.PHONY: all fmt lc

all:
	godep go install

fmt:
	gofmt -s -w -l bitballoon.go create.go deploy.go main.go update.go

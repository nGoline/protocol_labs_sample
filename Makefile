PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

install:
	go mod download
	go mod vendor

exchangeworker:	
	go build -o $(GOBIN)/exchangeworker ./cmd/exchangeworker/main.go || exit

build: exchangeworker

production:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/exchangeworker ./cmd/exchangeworker/main.go

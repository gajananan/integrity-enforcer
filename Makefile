GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=ie-server
LINTER=golangci-lint

all: test build

test:
		cd ./enforce; $(GOTEST) ./... -v

build:
		cd ./enforce;  $(GOBUILD) -o $(BINARY_NAME) -v

lint:
		cd ./enforce; $(LINTER) run

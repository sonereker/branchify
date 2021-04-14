GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=branchify
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_PATH=./cmd/branchify

all: get clean build build-linux
get:
	$(GOGET) ./cmd/branchify
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
	./$(BINARY_NAME)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(MAIN_PATH)

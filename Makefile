GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bookmaker.exe
BINARY_UNIX=$(BINARY_NAME)_unix

all: build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v ./cmd
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
		rm -rf ./output
test-json:
		$(GOBUILD) -o $(BINARY_NAME) -v ./cmd
		./$(BINARY_NAME) make -f ./2006.json

# Cross compilation
build-linux:
		GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd
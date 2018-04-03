# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=pro_binary
BINARY_UNIX=$(BINARY_NAME)_unix

all: run

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/gin-gonic/gin
	$(GOGET) github.com/go-sql-driver/mysql
	$(GOGET) github.com/go-redis/redis
	$(GOGET) github.com/jinzhu/gorm

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
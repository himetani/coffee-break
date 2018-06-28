.PHONY: all build build-linux install clean test

all: ;

NAME := coffee-time
REPOHOME := github.com/himetani/coffee-time
VERSION  := 0.0.9
REVISION  := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags="-s -w -X \"$(REPOHOME)/$(NAME)/cmd.Name=$(NAME)\" -X \"$(REPOHOME)/$(NAME)/cmd.Version=$(VERSION)\" -X \"$(REPOHOME)/$(NAME)/cmd.Revision=$(REVISION)\""

SRCS    := $(shell find . -path ./vendor -prune -o -name '*.go' -print)

bin/$(NAME): $(SRCS)
	@echo "=> go build $(LDFLAGS) -o bin/$(NAME)"
	@go build $(LDFLAGS) -o bin/$(NAME) && echo "=> Build success. Output: bin/$(NAME)"

bin/linux/$(NAME): $(SRCS)
	@echo "=> GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/linux/$(NAME)"
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/linux/$(NAME) && echo "=> Build success. Output: bin/linux/$(NAME)"

$$GOPATH/bin/$(NAME):
	go install $(LDFLAGS)

build: bin/$(NAME)

build-linux: bin/linux/$(NAME)

install: $$GOPATH/bin/$(NAME)

clean:
	rm -rf bin/*

test: 
	go test -v $(REPOHOME)/$(NAME)/... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

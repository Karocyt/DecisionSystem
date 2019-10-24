#******************************************************************************#
#*                                                                            *#
#*          ▄▄▄██▀▀▀      ███▄ ▄███▓         ██▓    ▄▄▄       ▄▄▄▄            *#
#*            ▒██        ▓██▒▀█▀ ██▒        ▓██▒   ▒████▄    ▓█████▄          *#
#*            ░██        ▓██    ▓██░        ▒██░   ▒██  ▀█▄  ▒██▒ ▄██         *#
#*         ▓██▄██▓       ▒██    ▒██         ▒██░   ░██▄▄▄▄██ ▒██░█▀           *#
#*          ▓███▒    ██▓ ▒██▒   ░██▒ ██▓    ░██████▒▓█   ▓██▒░▓█  ▀█▓         *#
#*          ▒▓▒▒░    ▒▓▒ ░ ▒░   ░  ░ ▒▓▒    ░ ▒░▓  ░▒▒   ▓▒█░░▒▓███▀▒         *#
#*          ▒ ░▒░    ░▒  ░  ░      ░ ░▒     ░ ░ ▒  ░ ▒   ▒▒ ░▒░▒   ░          *#
#*          ░ ░ ░    ░   ░      ░    ░        ░ ░    ░   ▒    ░    ░          *#
#*          ░   ░     ░         ░     ░         ░  ░     ░  ░ ░               *#
#*                    ░               ░                            ░          *#
#*                                                                            *#
#******************************************************************************#
                                   #* Makefile *#


SRC_FILES = engine.go \
			header.go \
			lexer.go \
			main.go \
			parser.go \
			infTree.go \
			utils.go

.PHONY: build get install run watch start stop restart fclean


GOPATH = $(shell pwd)
GOBIN = $(GOPATH)/bin
GOFILES = $(wildcard cmd/*.go)
GONAME = expert_system
TEST_FILE = example_input.txt

all:
	@echo "Building $(GOFILES) to ./bin"
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/$(GONAME) $(GOFILES)

get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

run: all
	@./bin/$(GONAME) ./examples/$(TEST_FILE)

fclean:
	@echo "Cleaning"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
	@rm -rf ./bin/

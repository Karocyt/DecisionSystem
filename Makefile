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

.PHONY: all get run test install fclean


GOPATH = $(shell pwd)
GOBIN = $(GOPATH)/bin
GOFILES = $(wildcard cmd/$(NAME)/*.go)
NAME = expert_system
TEST_FILE = ./examples/example_1.txt

all: $(NAME)

$(NAME): install
	

get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

install:
	@echo "Building $(GOFILES) to ./bin"
	@GOPATH=$(GOPATH) sh ./scripts/install.sh $(NAME)

run: all
	@./bin/$(NAME) $(TEST_FILE)

fclean:
	@echo "Cleaning"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
	@rm -rf ./bin/

test:
	sh ./scripts/tests.sh

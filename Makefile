SRC_DIR = cmd/

SRC_FILES = engine.go \
			header.go \
			lexer.go \
			main.go \
			parser.go \
			infTree.go

SRC = $(addprefix $(SRC_DIR), $(SRC_FILES))

all : $(SRC)
	go build -o expert_system $(SRC)
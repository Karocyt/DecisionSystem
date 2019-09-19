package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	openPar     string = "("
	closePar    string = ")"
	not         string = "!"
	and         string = "+"
	or          string = "|"
	xor         string = "^"
	imp         string = "=>"
	ioi         string = "<=>"
	com         string = "#"
	factSymbol  string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	factDeclar  string = "="
	queryDeclar string = "?"
)

type env struct {
	kb        []string
	trueFacts []string
	queries   []string
}

func main() {
	var env env
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		statement := strings.Trim(strings.Split(line, com)[0], " /t/n")
		if statement == "" {
			continue
		}
		if strings.HasPrefix(statement, factDeclar) {
			env.trueFacts = strings.Split(strings.TrimPrefix(statement, factDeclar), "")
		} else if strings.HasPrefix(statement, queryDeclar) {
			env.queries = strings.Split(strings.TrimPrefix(statement, queryDeclar), "")
		} else {
			env.kb = append(env.kb, statement)
		}
	}
	fmt.Printf("kb : %q\n", env.kb)
	fmt.Printf("trueFacts : %q\n", env.trueFacts)
	fmt.Printf("queries : %q\n", env.queries)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

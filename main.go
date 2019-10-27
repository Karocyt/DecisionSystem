package main

import (
	"os"
    "io/ioutil"
	"fmt"
	"internal/lexer"
	"errors"
)

const usage string = "Usage: ./%s {filename}\n"

func getInput() (string, error) {
	var s string

	if len(os.Args) == 2 {
		buf, e := ioutil.ReadFile(os.Args[1])
		s = string(buf)
		return s, e
	} else if len(os.Args) > 2 {
		return s, errors.New(usage)
	} else {
		s = "Bla bla bli blop bloup\n"
		fmt.Printf("Interactive mode ? Here you go.\n")
	}
	return s, nil
}

func main() {
	input, e := getInput()
	if e != nil {
		fmt.Println(e)
	} else {
		l := lexer.BeginLexing(input)
		fmt.Printf("Hi me! Test: %s", l.Input)
	}
}
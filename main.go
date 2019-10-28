package main

import (
	"os"
    "io/ioutil"
	"fmt"
	"internal/lexer"
	"errors"
)

func getInput() (string, error) {
	var s string

	if len(os.Args) == 2 {
		buf, e := ioutil.ReadFile(os.Args[1])
		s = string(buf)
		return s, e
	} else if len(os.Args) > 2 {
		return s, errors.New(fmt.Sprintf("Usage: %s {filename}", os.Args[0]))
	} else {
		return s, errors.New(fmt.Sprintf("Stdin mode ? Not Today.\nUsage: %s {filename}", os.Args[0]))
	}
	return s, nil
}

func main() {
	input, e := getInput()
	if e != nil {
		fmt.Println(e)
	} else {
		l := lexer.BeginLexing(input)
		for t := l.NextToken(); t!= nil; t = l.NextToken() {
			fmt.Printf("%s", t.Value)
		}
		fmt.Printf("Hi me!\nFull file was:\n%s\n{EOF}\n", l.Input)
	}
}
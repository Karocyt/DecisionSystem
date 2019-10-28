package main

import (
	"os"
    "io/ioutil"
	"fmt"
	"github.com/Karocyt/expertsystem/internal/lexer"
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
	if e == nil {
		l := lexer.BeginLexing(input, os.Args[1])

  		l.Emit(2)
  		l.Emit(2)
		for t := <-l.Tokens; l.Error == nil; t = <-l.Tokens {
			fmt.Printf("Token: %s\n", t.Value)
			l.Error = &lexer.LexingError{Lexer: l,Expected: "not that much",Got: "a lot",Line: l.Line,Pos: l.PosInLine()}
		}
		if l.Error != nil {
			fmt.Println(l.Error)
		}
	}
}
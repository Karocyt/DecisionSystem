package main

import (
	"fmt"
	"github.com/Karocyt/expertsystem/internal/lexer"
	"testing"
)

func TestTokenLoop(test *testing.T) {
	var flag int
	l := lexer.BeginLexing("", "ghostFile")
 	l.Emit(2)
	for t := <-l.Tokens; l.Error == nil; t = <-l.Tokens {
		fmt.Printf("Token: %s\n", t.Value)
		flag++
	}
	if l.Error != nil {
		fmt.Println(l.Error)
	} else if flag == 0 {
		fmt.Println(&lexer.LexingError{Lexer: l,Expected: "Token",Got: "nothing",Line: l.Line,Pos: l.PosInLine()})
	}
}
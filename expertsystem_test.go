package main

import (
	"github.com/Karocyt/expertsystem/internal/lexer"
	"testing"
)

func TestTokenLoop(test *testing.T) {
	var flag int
	l := lexer.BeginLexing("A+B=>C", "ghostFile")
 	l.Emit(2)
	for t := <-l.Tokens; l.Error == nil; t = <-l.Tokens {
		if t.Type != 99 { flag++ }
	}
	if l.Error != nil {
		test.Error(l.Error)
	} else if flag == 0 {
		test.Error(&lexer.LexingError{Lexer: l,Expected: "Token",Got: "nothing",Line: l.Line,Pos: l.PosInLine()})
	}
}
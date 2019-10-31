package main

import (
	"github.com/Karocyt/expertsystem/internal/lexer"
	"testing"
	"io/ioutil"
	"fmt"
)

func TestTokenLoop(test *testing.T) {
	var flag int
	if lexer.Debug {println("--- Start TestTokenLoop ---")}
	l := lexer.BeginLexing("A+B=>C", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 { flag++ }
	}
	if l.Error != nil {
		test.Error(l.Error)
	} else if flag == 0 {
		test.Error(&lexer.LexingError{Lexer: l,Expected: "Token",Got: "nothing",Line: l.Line,Pos: l.PosInLine()})
	}
}

func TestNoResult(test *testing.T) {
	var flag int
	if lexer.Debug {println("--- Start TestNoResult ---")}
	l := lexer.BeginLexing("A+B=>", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 { flag++ }
	}
	if l.Error == nil {
		test.Error()
	}
}

func TestNoOperator(test *testing.T) {
	var flag int
	if lexer.Debug {println("--- Start TestNoOperator ---")}
	l := lexer.BeginLexing("AB=>C", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 { flag++ }
	}
	if l.Error == nil {
		test.Error()
	}
}

func TestDoubleResult(test *testing.T) {
	var flag int
	if lexer.Debug {println("--- Start TestDoubleResult ---")}
	l := lexer.BeginLexing("A+!B=>CD", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 { flag++ }
	}
	if l.Error == nil {
		test.Error()
	}
}

func TestFalseValues(test *testing.T) {
	var flag int
	if lexer.Debug {println("--- Start TestFalseValues ---")}
	l := lexer.BeginLexing("A+!B=>C", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 { flag++ }
	}
	if l.Error != nil {
		test.Error(l.Error)
	}
}

func TestInput1(test *testing.T) {
	var flag int
	if lexer.Debug {println("--- Start TestInput1 ---")}
	file := "testdata/input1.txt"
	content, e := ioutil.ReadFile(file)
	str := string(content)
	if e != nil {
		test.Error(fmt.Sprintf("Unable to read file: %s\n", file))
	}
	l := lexer.BeginLexing(str, file)
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 { flag++ }
	}
	if l.Error != nil {
		test.Error(l.Error)
	} else if flag == 0 {
		test.Error(&lexer.LexingError{Lexer: l,Expected: "Token",Got: "nothing",Line: l.Line,Pos: l.PosInLine()})
	}
}
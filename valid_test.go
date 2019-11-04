package main

import (
	"fmt"
	"github.com/Karocyt/expertsystem/internal/lexer"
	"io/ioutil"
	"testing"
)

func TestTokenLoop(test *testing.T) {
	var flag int
	if lexer.Debug {
		println("--- Start TestTokenLoop ---")
	}
	l := lexer.BeginLexing("A+B=>C\n=A\n?B", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error(l.Error)
	} else if flag == 0 {
		test.Error(lexer.LexingError{Lexer: l, Expected: "Token", Got: "nothing"})
	}
}

func TestFalseValues(test *testing.T) {
	var flag int
	if lexer.Debug {
		println("--- Start TestFalseValues ---")
	}
	l := lexer.BeginLexing("S+!R=>F\n=A\n?B", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error(l.Error)
	}
}

func TestFullWords(test *testing.T) {
	var flag int
	if lexer.Debug {
		println("--- Start TestDoubleResult ---")
	}
	l := lexer.BeginLexing("Hello+!World=>Working #Prout\n=A\n?B", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error()
	}
}

func TestCommentEOL(test *testing.T) {
	var flag int
	if lexer.Debug {
		println("--- Start TestCommentEOL ---")
	}
	l := lexer.BeginLexing("Comment+!Soon=>Ok #Prout\n=A\n?B", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error()
	}
}

func TestCommentFirstLastLine(test *testing.T) {
	var flag int
	if lexer.Debug {
		println("--- Start TestCommentFirstLastLine ---")
	}
	l := lexer.BeginLexing("#Blabla\nComment=>Passed #Prout\n=A\n?B", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error()
	}
}

func TestSpaces(test *testing.T) {
	var flag int
	if lexer.Debug {
		println("--- Start TestSpaces ---")
	}
	l := lexer.BeginLexing("  Spaces\t+ \t Everywhere => \tOk \t#Prout\n=A\n?B", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error()
	}
}

func TestInput1(test *testing.T) {
	var flag int
	if lexer.Debug {
		println("--- Start TestInput1 ---")
	}
	file := "testdata/input1.txt"
	content, e := ioutil.ReadFile(file)
	str := string(content)
	if e != nil {
		test.Error(fmt.Sprintf("Unable to read file: %s\n", file))
	}
	l := lexer.BeginLexing(str, file)
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error(l.Error)
	} else if flag == 0 {
		test.Error(lexer.LexingError{Lexer: l, Expected: "Token", Got: "nothing"})
	}
}

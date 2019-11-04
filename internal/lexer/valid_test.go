package lexer

import (
	"testing"
)

func TestTokenLoop(test *testing.T) {
	var flag int
	if Debug {
		println("--- Start TestTokenLoop ---")
	}
	l := BeginLexing("A+B=>C\n=A\n?B", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error(l.Error)
	} else if flag == 0 {
		test.Error(LexingError{Lexer: l, Expected: "Token", Got: "nothing"})
	}
}

func TestFalseValues(test *testing.T) {
	var flag int
	if Debug {
		println("--- Start TestFalseValues ---")
	}
	l := BeginLexing("S+!R=>F\n=A\n?B", "ghostFile")
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
	if Debug {
		println("--- Start TestDoubleResult ---")
	}
	l := BeginLexing("Hello+!World=>Working #Prout\n=A\n?B", "ghostFile")
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
	if Debug {
		println("--- Start TestCommentEOL ---")
	}
	l := BeginLexing("Comment+!Soon=>Ok #Prout\n=A\n?B", "ghostFile")
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
	if Debug {
		println("--- Start TestCommentFirstLastLine ---")
	}
	l := BeginLexing("#Blabla\nComment=>Passed #Prout\n=A\n?B", "ghostFile")
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
	if Debug {
		println("--- Start TestSpaces ---")
	}
	l := BeginLexing("  Spaces\t+ \t Everywhere => \tOk \t#Prout\n=A\n?B", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error != nil {
		test.Error()
	}
}

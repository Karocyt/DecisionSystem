package lexer

import (
	"testing"
)

func TestNoResult(test *testing.T) {
	var flag int
	if Debug {
		println("--- Start TestNoResult ---")
	}
	l := BeginLexing("N+M=>", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error == nil {
		test.Error()
	}
}

func TestNoOperator(test *testing.T) {
	var flag int
	if Debug {
		println("--- Start TestNoOperator ---")
	}
	l := BeginLexing("KL=>X", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error == nil {
		test.Error()
	}
}

func TestDoubleResult(test *testing.T) {
	var flag int
	if Debug {
		println("--- Start TestDoubleResult ---")
	}
	l := BeginLexing("P+!Q=>TY", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error == nil {
		test.Error()
	}
}

func TestCommentMiddle(test *testing.T) {
	var flag int
	if Debug {
		println("--- Start TestCommentMiddle ---")
	}
	l := BeginLexing("P+!Q#=>T", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error == nil {
		test.Error()
	}
}

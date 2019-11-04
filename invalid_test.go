package main

import (
//	"fmt"
	"github.com/Karocyt/expertsystem/internal/lexer"
//	"io/ioutil"
	"testing"
)

func TestNoResult(test *testing.T) {
	var flag int
	if lexer.Debug {
		println("--- Start TestNoResult ---")
	}
	l := lexer.BeginLexing("N+M=>", "ghostFile")
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
	if lexer.Debug {
		println("--- Start TestNoOperator ---")
	}
	l := lexer.BeginLexing("KL=>X", "ghostFile")
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
	if lexer.Debug {
		println("--- Start TestDoubleResult ---")
	}
	l := lexer.BeginLexing("P+!Q=>TY", "ghostFile")
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
	if lexer.Debug {
		println("--- Start TestCommentMiddle ---")
	}
	l := lexer.BeginLexing("P+!Q#=>T", "ghostFile")
	for t := <-l.Tokens; l.State != nil; t = <-l.Tokens {
		if t.Type != 99 {
			flag++
		}
	}
	if l.Error == nil {
		test.Error()
	}
}
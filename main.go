package main

import (
	"fmt"
	"github.com/Karocyt/expertsystem/internal/lexer"
)

func main() {
	l := lexer.BeginLexing("G+(A)+B=>(C|D)")
	fmt.Printf("Hi me! Test: %s", l.Input)
}
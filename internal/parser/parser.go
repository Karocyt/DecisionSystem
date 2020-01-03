package parser

import (
	"github.com/Karocyt/expertsystem/internal/lexer"
)

func Parse(input string, filename string) (count int, e error) {
	l := lexer.BeginLexing(input, filename)
	for t := range l.Tokens {
		if l.Error != nil {
			break
		}
		if t.Type > 0 {
			count++
		}
	}
	e = l.Error
	return
}
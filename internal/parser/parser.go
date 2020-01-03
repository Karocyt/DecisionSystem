package parser

import (
	"github.com/Karocyt/expertsystem/internal/lexer"
	"errors"
	"fmt"
)

const MAX_ITEMS = 50


func Parse(input string, filename string) (count int, e error) {
	l := lexer.BeginLexing(input, filename)
	a := make([]lexer.LexToken, 0, 50)
	for t := range l.Tokens {
		if l.Error != nil {
			break
		}
		if t.Type == lexer.TOKEN_EOL {
			a = make([]lexer.LexToken, MAX_ITEMS)
			fmt.Println("\n")
		} else if t.Type != lexer.TOKEN_EOF {
			a = append(a, t)
			fmt.Println(a[len(a)-1])
			count++
			if count == MAX_ITEMS {
				l.Error = errors.New("Line too long.")
				break
			}
		}



	}
	e = l.Error
	return
}
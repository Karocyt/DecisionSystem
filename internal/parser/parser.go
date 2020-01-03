package parser

import (
	"github.com/Karocyt/expertsystem/internal/lexer"
	"errors"
	"fmt"
)

const MAX_ITEMS = 50


func process_line(a []lexer.LexToken) {
	fmt.Println(a)
	fmt.Println("\n")
}

func Parse(input string, filename string) (count int, e error) {
	l := lexer.BeginLexing(input, filename)
	a := make([]lexer.LexToken, 0, 50)
	for t := range l.Tokens {
		if l.Error != nil {
			break
		}
		if t.Type == lexer.TOKEN_EOL {
			if len(a) > 0 {
				process_line(a)
			}
			a = make([]lexer.LexToken, 0, MAX_ITEMS)
		} else if t.Type != lexer.TOKEN_EOF {
			a = append(a, t)
			count++
			if len(a) == MAX_ITEMS {
				l.Error = errors.New("Line too long.")
				break
			}
		} else {
			l.Error = errors.New("Unexpected error.")
			break
		}



	}
	e = l.Error
	return
}
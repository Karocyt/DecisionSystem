package parser

import (
	"github.com/Karocyt/expertsystem/internal/lexer"
	"errors"
	"fmt"
)


func process_line(a []lexer.LexToken) { //Left to do: Polish notation refactoring, array of tokens
	fmt.Println(a)
	fmt.Println("\n")
}

func Parse(input string, filename string) (count int, e error) {
	l := lexer.BeginLexing(input, filename)
	a := make([]lexer.LexToken, 0)
	for t := range l.Tokens {
		if l.Error != nil {
			break
		}
		if t.Type == lexer.TOKEN_EOL {
			if len(a) > 0 {
				process_line(a)
			}
			a = make([]lexer.LexToken, 0)
		} else if t.Type != lexer.TOKEN_EOF {
			a = append(a, t)
			count++
		} else {
			l.Error = errors.New("Unexpected error.")
			break
		}
	}
	e = l.Error
	return
}
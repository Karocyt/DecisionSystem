package parser

import (
	"github.com/Karocyt/expertsystem/internal/lexer"
	"errors"
	"fmt"
)

type Node struct {
	token 	lexer.LexToken
	state	bool
	left  	lexer.LexToken
	right	lexer.LexToken
}

type Facts struct {
	table 	map[string]Node
	queries	[]Node
}

// func find_operator(a []lexer.lextoken, match string) (i int) {
// 	for i, t := range a {
// 		if t.Value == match {
// 			return
// 		}
// 	}
// 	i = -1
// 	return
// }

func build_tree(a []lexer.LexToken) (tree Node) {
	fmt.Println("Building Tree", a)
	fmt.Println("Operator: ", a[find_operator(a)])
	return tree
}

func process_line(a []lexer.LexToken) { //Left to do: build tree and hashtable
	if a[0].Type == lexer.TOKEN_EQUALS {
		fmt.Println("Process Facts", a)
	} else if a[0].Type == lexer.TOKEN_QUERY {
		fmt.Println("Process Queries", a)
	} else {
		index := 0
		for i, t := range a {
			if t.Type == lexer.TOKEN_IF_ONLY_IF || t.Type == lexer.TOKEN_IMPLIES {
				index = i
			}
		}
		left := a[0 : index]
		right := a[index + 1 : len(a)]
		operator := a[index]
		tree := build_tree(left)
		fmt.Println(left, right, operator, tree)
	}


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
	if len(a) > 0 {
		process_line(a)
	}
	e = l.Error
	return
}
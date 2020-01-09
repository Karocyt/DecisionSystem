package parser

import (
	"github.com/Karocyt/expertsystem/internal/lexer"
	//"errors"
	"fmt"
)

const (
	LEFT_BRACKET  	string = "("
	RIGHT_BRACKET 	string = ")"
	IMPLIES       	string = "=>"
	IF_ONLY_IF    	string = "<=>"
	EQUALS        	string = "="
	QUERY         	string = "?"
	NOT         	string = "!"
	AND 			string = "+"
	OR 				string = "|"
	XOR 			string = "^"
)

type Builder struct {
	register map[string]Node
	queries	[]string
}

func create_node(left []string, right []string) {
	fmt.Println("create_node\n")
}

func prec_dict() func(string) int {
	p := map[string]int{
		IF_ONLY_IF:1,
		IMPLIES:2,
		XOR:3,
		OR:4,
		AND:5,
		NOT:6,
		LEFT_BRACKET:7,
	}
	return func(key string) int {
		return p[key]
	}
}

func find_operator(a []string) (i int) {
	precedence := prec_dict()

	best := -1
	match := 0
	for i, s := range a {
		p := precedence(s)
		if p > best {
			best = p
			match = i
		}
	}
	return match
}

func build_tree(a []string) (tree Node) {
	fmt.Println("Building Tree", a)
	fmt.Println("Operator: ", a[find_operator(a)])
	return tree
}

func process_line(a []string) { //Left to do: build tree and hashtable
	if a[0] == EQUALS {
		fmt.Println("Process Facts", a)
	} else if a[0] == QUERY {
		fmt.Println("Process Queries", a)
	} else {
		index := 0
		for i, t := range a {
			if t == IF_ONLY_IF || t == IMPLIES {
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

func (b Builder) init(l *lexer.Lexer) (e error) {
	b.register = make(map[string]Node)
	b.queries	= make([]string, 0)

	a := make([]string, 0)
	i := 0
	for t := range l.Tokens {
		if t == "\n" {
			if len(a) > 0 {
				process_line(a)
			}
			a = make([]string, 0)
		} else {
			a = append(a, t)
		}
		i++
	}
	if len(a) > 0 {
		process_line(a)
	}
	return
}

func New(input string, filename string) (b Builder, e error) {
	l, e := lexer.BeginLexing(input, filename)
	if e != nil {
		return
	}
	b = Builder{}
	e = b.init(l)
	return
}


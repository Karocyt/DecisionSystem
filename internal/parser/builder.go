package parser

import (
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
	Rules map[string]Node
	Queries	[]string
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

func (b *Builder) build_tree(a []string) (tree Node) {
	index := 0
	for i, t := range a {
		if t == IF_ONLY_IF || t == IMPLIES {
			index = i
		}
	}
	left := a[0 : index]
	right := a[index + 1 : len(a)]
	operator := a[index]
	fmt.Println(left, right, operator, tree)
	fmt.Println("\n")
	return tree
}

func (b *Builder) process_query(a []string) {
	for _, s := range a[1 : len(a)] {
		b.Queries = append(b.Queries, s)
	}
}

func (b *Builder) process_facts(a []string) {
	for _, s := range a[1 : len(a)] {
		b.Rules[s] = Key{s, true, KEY_GIVEN}
	}
}

func (b *Builder) process_rule(a []string) {
	index := 0
	for i, t := range a {
		if t == IF_ONLY_IF || t == IMPLIES {
			index = i
		}
	}
	rule := a[0 : index]
	result := a[index + 1 : len(a)]
	relation := a[index]
	tree := b.build_tree(rule)
	fmt.Println("line:\t\t", rule, relation, result, "\nrule tree:\t", tree)
}

func (b *Builder) process_line(a []string) { //Left to do: build tree and hashtable
	if a[0] == EQUALS {
		b.process_facts(a)
	} else if a[0] == QUERY {
		b.process_query(a)
	} else {
		b.process_rule(a)	
	}
}

// IMPLIES == multiples rules OR
// IOF == multiple rules AND
func (b *Builder) build(tokens chan string) (e error) {
	b.Rules = make(map[string]Node)
	//b.Queries = make([]string, 0)

	a := make([]string, 0)
	i := 0
	for t := range tokens {
		if t == "\n" {
			if len(a) > 0 {
				b.process_line(a)
			}
			a = make([]string, 0)
		} else {
			a = append(a, t)
		}
		i++
	}
	if len(a) > 0 {
		b.process_line(a)
	}
	return
}

func New(input chan string) (b Builder, e error) {
	if e != nil {
		return
	}
	b = Builder{}
	e = b.build(input)
	return
}


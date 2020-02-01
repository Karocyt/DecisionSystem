package solver

import (
	"errors"
	"fmt"
)

type Builder struct {
	Variables map[string]*Key
	Queries	[]string
	Facts 	[]string
}

func map_precedence() map[string]int {
	return map[string]int{
		IF_ONLY_IF:7,
		IMPLIES:6,
		XOR:5,
		OR:4,
		AND:3,
		NOT:2,
		LEFT_BRACKET:1,
	}
}

func next_bracket(a []string, i int) (int) {
	count := 1
	for ; i < len(a); i++ {
		if a[i] == LEFT_BRACKET {
			count ++
		} else if a[i] == RIGHT_BRACKET {
			count--
		}
		if count == 0 {
			return i
		}
	}
	return i
}

func find_operator(a []string) (i int) {
	best := -1
	match := -1
	for i := 0; i < len(a); i++ {
		p := map_precedence()[a[i]]
		if p > best {
			best = p
			match = i
		}
		if a[i] == LEFT_BRACKET {
			i = next_bracket(a, i + 1)
		}
	}
	return match
}

func (b *Builder) Eval_rules(s string) (value bool, e error) {
	k, ok := b.Variables[s]
	if !ok {
		b.Variables[s] = &Key{}
		k = b.Variables[s]
		k.Name = s
		var op False
		k.Child = &op
	}
	return k.Eval(make([]string, 0))
}

func (b *Builder) append_implies(rule Rule) (e error) {
	// Left to do operator in right operand 					/// need to make big op
	// for i, _ := range rule.Right {
	// 	node := b.build_tree(rule.Right[i : i + 1])
	// 	key, ok := node.(*Key)
	// 	if ok {
	fmt.Println(rule.Right, rule.Left)
	if len(rule.Right) == 1	{
		fmt.Println("Ingesting simple rule for", rule.Right)
		node := b.build_tree(rule.Right[0 : 1])
		fmt.Println(node)
		node.(*Key).Child, e = add_op(rule.Left, node.(*Key).Child)
		fmt.Println(node)
	} else if len(rule.Right) == 2 && rule.Right[0] == NOT {
		fmt.Println("Ingesting Not for", rule.Right)
		var op Not
		op.Right = rule.Left
		node := b.build_tree(rule.Right[1 : 2])
		node.(*Key).Child, e = add_op(&op, node.(*Key).Child) // not sure of this (NOT in conclusion) !!! <------
	} else if rule.Right[1] == AND {
		fmt.Println("Ingesting And for", rule.Right)
		node1 := b.build_tree(rule.Right[0 : 1])
		node2 := b.build_tree(rule.Right[2 : 3])
		node1.(*Key).Child, e = add_op(rule.Left, node1.(*Key).Child)
		if e != nil {
			return
		}
		node2.(*Key).Child, e = add_op(rule.Left, node2.(*Key).Child)
	} 													// left to do Or in conclusion
	return
}

func (b *Builder) build_tree(a []string) (tree Node) {
	if len(a) == 1 {
		_, ok := b.Variables[a[0]]
		if !ok {
			b.Variables[a[0]] = &Key{}
		}
		b.Variables[a[0]].Name = a[0]
		return b.Variables[a[0]]
	}
	index := find_operator(a)
	left := a[0 : index]
	right := a[index + 1 : len(a)]
	operator := a[index]

	switch operator {
	case AND:
		var op And
		op.Left, op.Right = b.build_tree(left), b.build_tree(right)
		return &op
	case OR:
		var op Or
		op.Left, op.Right = b.build_tree(left), b.build_tree(right)
		return &op
	case XOR:
		var op Xor
		op.Left, op.Right = b.build_tree(left), b.build_tree(right)
		return &op
	case NOT:
		var op Not
		op.Right = b.build_tree(right)
		return &op
	case LEFT_BRACKET:
		var op Parenthesis
		op.Op = b.build_tree(a[1 : len(a) - 1])
		return &op
	}
	return nil
}

func (b *Builder) process_query(a []string) {
	for _, s := range a[1 : len(a)] {
		b.Queries = append(b.Queries, s)
	}
}

func (b *Builder) take_facts(a []string) {
	if b.Facts == nil {
		b.Facts = a
	} else {
		fmt.Println("Ignoring duplicate Facts line")
	}
}

func (b *Builder) process_facts() (e error) {
	if b.Facts == nil {
		return errors.New("Missing Facts line")
	}
	for _, s := range b.Facts[1 : len(b.Facts)] {
		_, ok := b.Variables[s]
		var op True
		if !ok {
			b.Variables[s] = &Key{Name:s,Child:&op}
		} else {
			b.Variables[s].Child = &op
		}
	}
	return
}

func (b *Builder) process_rule(a []string) (e error) {
	index := 0
	for i, t := range a {
		if t == IF_ONLY_IF || t == IMPLIES {
			index = i
		}
	}
	rule := a[0 : index]
	result := a[index + 1 : len(a)]
	tree := b.build_tree(rule)
	e = b.append_implies(Rule{tree, a[index], result})
	return e
}

func (b *Builder) process_line(a []string) (e error) {
	if a[0] == EQUALS {
		b.take_facts(a)
	} else if a[0] == QUERY {
		b.process_query(a)
	} else {
		e = b.process_rule(a)	
	}
	return
}

// IMPLIES == multiples rules OR
// IOF == multiple rules AND
func (b *Builder) build(tokens chan string) (e error) {
	b.Variables = make(map[string]*Key)

	a := make([]string, 0)
	i := 0
	for t := range tokens {
		if e != nil {
			return e
		}
		if t == "\n" {
			if len(a) > 0 {
				e = b.process_line(a)
			}
			a = make([]string, 0)
		} else {
			a = append(a, t)
		}
		i++
	}
	if e != nil {
		return e
	}
	e = b.process_facts()
	return
}

func New(input chan string) (b Builder, e error) {
	if e != nil {
		return
	}
	b = Builder{}
	e = b.build(input)
	return b, e
}


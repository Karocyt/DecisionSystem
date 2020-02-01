package solver

import (
	//"errors"
	//"fmt"
)

type Builder struct {
	Variables map[string]*Key
	//Rules map[string]Node
	Queries	[]string
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
	//precedence := prec_dict()

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
	// fmt.Println("\tEval_rules")
	// defer fmt.Println("\tEnd Eval rules")
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

func (b *Builder) append_implies(rule Defines) (e error) {
	// Left to do operator in right operand 					/// need to make big op
	for i, _ := range rule.Right {
		node := b.build_tree(rule.Right[i : i + 1])
		key, ok := node.(*Key)
		if ok {
			if len(rule.Right) == 1	{			///// BRICOLAGE TOUT CASSÉ, NE PREND PAS LES
				key.Child = rule.Left // MULTI-RULE
			} else if len(rule.Right) == 2 && rule.Right[0] == NOT {
				var op Not
				op.Right = rule.Left
				key.Child = &op
			}
		}
	}
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
	//fmt.Println("build_tree:\n-left:\t\t", left, "\n-right:\t\t", right, "\n-operator:\t",  operator)

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

func (b *Builder) process_facts(a []string) {
	for _, s := range a[1 : len(a)] {
		_, ok := b.Variables[s]
		if !ok {
			b.Variables[s] = &Key{Name:s, Value:true, State:KEY_GIVEN}
		} else {
			b.Variables[s].Value, b.Variables[s].State = true, KEY_GIVEN
		}
	}
}

func (b *Builder) process_rule(a []string) (e error) {
	// fmt.Println("\tProcess_rule", a)
	// defer fmt.Println("\tEnd process rule", a)
	index := 0
	for i, t := range a {
		if t == IF_ONLY_IF || t == IMPLIES {
			index = i
		}
	}
	rule := a[0 : index]
	result := a[index + 1 : len(a)]
	tree := b.build_tree(rule)
	e = b.append_implies(Defines{tree, a[index], result})
	//fmt.Println("rule tree:\t", tree, "\n")
	return e
}

func (b *Builder) process_line(a []string) (e error) { //Left to do: build tree and hashtable
	if a[0] == EQUALS {
		b.process_facts(a)
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
	//b.Rules = make(map[string]Node)
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


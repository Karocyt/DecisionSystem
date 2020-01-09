package parser

import (
	"fmt"
)

type Node interface{  
    Eval(key string) (bool, error)
}

type And struct {  
    Left, Right Node
}

func (op And) Eval(key string) (bool, error) {
	val1, e := op.Left.Eval(key)
	if e != nil {
		return false, e
	}
	val2, e := op.Right.Eval(key)
	return val1 && val2, e
}

func (op And) String() string {
	return fmt.Sprintf("(%s + %s)", op.Left, op.Right)
}

type Or struct {  
    Left, Right Node
}

func (op Or) Eval(key string) (bool, error) {
	val1, e := op.Left.Eval(key)
	if e != nil {
		return false, e
	}
	val2, e := op.Right.Eval(key)
	return val1 || val2, e
}

func (op Or) String() string {
	return fmt.Sprintf("(%s | %s)", op.Left, op.Right)
}

type Xor struct {  
    Left, Right Node
}

func (op Xor) Eval(key string) (bool, error) {
	val1, e := op.Left.Eval(key)
	if e != nil {
		return false, e
	}
	val2, e := op.Right.Eval(key)
	return val1 != val2, e
}

func (op Xor) String() string {
	return fmt.Sprintf("(%s ^ %s)", op.Left, op.Right)
}

type Not struct {  
    Right Node
}

func (op Not) Eval(key string) (bool, error) {
	val, e := op.Right.Eval(key)
	return !val, e
}

func (op Not) String() string {
	return fmt.Sprintf("!%s", op.Right)
}

type Parenthesis struct {  
    Op Node
}

func (op Parenthesis) Eval(key string) (bool, error) {
	return op.Op.Eval(key)
}

func (op Parenthesis) String() string {
	return fmt.Sprintf("%s", op)
}

type Defines struct {  
    Left Node
    operator string
    Right []string
}

func (op Defines) Apply(b *Builder) (e error) {
	if len(op.Right) == 1 {
		_, e := op.Left.Eval(op.Right[0])
		if e != nil {
			return e
		}
		k := b.Variables[op.Right[0]]
		k.Name = op.Right[0]
		e = k.Set(true)
		if e != nil {
			return e
		}
		b.Variables[op.Right[0]] = k
	} else if len(op.Right) == 2 && op.Right[0] == NOT {
		_, e := op.Left.Eval(op.Right[1])
		if e != nil {
			return e
		}
		k := b.Variables[op.Right[0]]
		k.Name = op.Right[0]
		e = k.Set(false)
		if e != nil {
			return e
		}
		b.Variables[op.Right[0]] = k
	}
	// // Left to do operators in Defines and IOI
	// index := find_operator(op.Right)
	// left := op.Right[0 : index]
	// right := op.Right[index + 1 : len(a)]
	// operator := a[index]
	
	// //for each right, eval left
	// val, e := op.Left.Eval("")
	return e
}

func (op Defines) String() string {
	return fmt.Sprintf("%s %s %s", op.Left, op.operator, op.Right)
}

package solver

import (
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

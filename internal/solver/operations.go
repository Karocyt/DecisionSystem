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
    Eval(keys []string) (bool, error)
}

type And struct {  
    Left, Right Node
}

func (op And) Eval(keys []string) (bool, error) {
	val1, e := op.Left.Eval(keys)
	if e != nil {
		return false, e
	}
	val2, e := op.Right.Eval(keys)
	return val1 && val2, e
}

func (op And) String() string {
	return fmt.Sprintf("(%T + %T)", op.Left, op.Right)
}

type Or struct {  
    Left, Right Node
}

func (op Or) Eval(keys []string) (bool, error) {
	val1, e := op.Left.Eval(keys)
	if !val1 || e != nil {
		return op.Right.Eval(keys) ////////////// ignore errors in left operand, to comment if corrector close-minded
	}
	return val1, e
}

func (op Or) String() string {
	return fmt.Sprintf("(%T | %T)", op.Left, op.Right)
}

type Xor struct {  
    Left, Right Node
}

func (op Xor) Eval(keyss []string) (bool, error) {
	val1, e := op.Left.Eval(keyss)
	if e != nil {
		return false, e
	}
	val2, e := op.Right.Eval(keyss)
	return val1 != val2, e
}

func (op Xor) String() string {
	return fmt.Sprintf("(%T ^ %T)", op.Left, op.Right)
}

type Not struct {  
    Right Node
}

func (op Not) Eval(keys []string) (bool, error) {
	val, e := op.Right.Eval(keys)
	return !val, e
}

func (op Not) String() string {
	return fmt.Sprintf("!%T", op.Right)
}

type Parenthesis struct {  
    Op Node
}

func (op Parenthesis) Eval(keys []string) (bool, error) {
	return op.Op.Eval(keys)
}

func (op Parenthesis) String() string {
	return fmt.Sprintf("%T", op)
}

type True struct {  
}

func (op True) Eval(keys []string) (mybool bool, e error) {
	return true, e
}

func (op True) String() string {
	return fmt.Sprintf("True")
}

type False struct {  
}

func (op False) Eval(keys []string) (mybool bool, e error) {
	return false, e
}

func (op False) String() string {
	return fmt.Sprintf("False")
}

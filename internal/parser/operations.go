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

type Implies struct {  
    Left, Right Node
}

func (op Implies) Eval(key string) (bool, error) {
	// val1, e := op.Left.Eval(key)
	// if e != nil {
	// 	return false, e
	// }
	// val2, e := op.Right.Eval(key)
	//if val1 {

	//}
	val, e := op.Left.Eval("")
	if val {
		apply(op.Right)
	}
	return val, e
}

func (op Implies) String() string {
	return fmt.Sprintf("%s => %s", op.Left, op.Right)
}

type If_Only_If struct {  
    Left, Right Node
}

func (op If_Only_If) Eval(key string) (bool, error) {
	// val1, e := op.Left.Eval(key)
	// if e != nil {
	// 	return false, e
	// }
	// val2, e := op.Right.Eval(key)
	//if val1 {

	//}
	val, e := op.Right.Eval("")
	return val, e
}

func (op If_Only_If) String() string {
	return fmt.Sprintf("%s <=> %s", op.Left, op.Right)
}

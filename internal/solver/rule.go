package solver

import (
	"fmt"
)

type Rule struct {  
    Left Node
    operator string
    Right []string
}


func (op Rule) String() string {
	return fmt.Sprintf("%T %s %T", op.Left, op.operator, op.Right)
}

/*
	Appends an "=>" rule to the pre-existing set via a Or operation set on top
*/
func add_op(to_add Node, child Node) (new Node, e error) {
	if child == nil {
		return to_add, e
	}
	var op Or
	op.Left, op.Right = child, to_add
	return &op, e
}
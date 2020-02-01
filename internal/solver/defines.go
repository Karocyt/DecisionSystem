package solver

import (
	"fmt"
)

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
		_, ok := b.Variables[op.Right[0]]
		if !ok {
			b.Variables[op.Right[0]] = &Key{}
			b.Variables[op.Right[0]].Name = op.Right[0]
		}
		e = b.Variables[op.Right[0]].Set(true)
		if e != nil {
			return e
		}
	} else if len(op.Right) == 2 && op.Right[0] == NOT {
		_, e := op.Left.Eval(op.Right[1])
		if e != nil {
			return e
		}
		_, ok := b.Variables[op.Right[0]]
		if !ok {
			b.Variables[op.Right[0]] = &Key{}
			b.Variables[op.Right[0]].Name = op.Right[0]
		}
		e = b.Variables[op.Right[0]].Set(false)
		if e != nil {
			return e
		}
	}
	// // Left to do operators in Defines and IOI
	// index := find_operator(op.Right)
	// left := op.Right[0 : index]
	// right := op.Right[index + 1 : len(a)]
	// operator := a[index]
	
	// //for each right, eval left
	// val, e := op.Left.Eval("")
	// // or translate right ops into left ops:
	// "A+B => C|D" generate rules:
	//		!C + A +B => D
	//          AND
	//		!D + A + B => C

	// for now taking op.Right[0] (so ignoring leftover)
	return e
}

func (op Defines) String() string {
	return fmt.Sprintf("%s %s %s", op.Left, op.operator, op.Right)
}

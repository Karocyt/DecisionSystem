package solver

import (
	"fmt"
)

type Rule struct {  
    Left Node
    operator string
    Right []string
}

// MANAGE DEFINES CONCATENATION

func (op Rule) String() string {
	return fmt.Sprintf("%s %s %s", op.Left, op.operator, op.Right)
}

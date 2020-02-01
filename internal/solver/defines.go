package solver

import (
	"fmt"
)

type Defines struct {  
    Left Node
    operator string
    Right []string
}

// MANAGE DEFINES CONCATENATION

func (op Defines) String() string {
	return fmt.Sprintf("%s %s %s", op.Left, op.operator, op.Right)
}

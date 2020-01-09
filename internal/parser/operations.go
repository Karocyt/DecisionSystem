package parser

type Node interface{  
    Eval() bool
}

type And struct {  
    Left, Right Node
}

func (op And) Eval() bool {
	return op.Left.Eval() && op.Right.Eval()
} 

type Or struct {  
    Left, Right Node
}

func (op Or) Eval() bool {
	return op.Left.Eval() || op.Right.Eval()
}

type Xor struct {  
    Left, Right Node
}

func (op Xor) Eval() bool {
	return op.Left.Eval() != op.Right.Eval()
}

type Not struct {  
    Left Node
}

func (op Not) Eval() bool {
	return !op.Left.Eval()
}

type Parenthesis struct {  
    Op Node
}

func (op Parenthesis) Eval() bool {
	return op.Op.Eval()
}

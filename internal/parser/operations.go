package parser

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

type Not struct {  
    Left Node
}

func (op Not) Eval(key string) (bool, error) {
	val, e := op.Left.Eval(key)
	return !val, e
}

type Parenthesis struct {  
    Op Node
}

func (op Parenthesis) Eval(key string) (bool, error) {
	return op.Op.Eval(key)
}

type Default struct {  
}

func (op Default) Eval(key string) (bool, error) {
	return false, nil
}

package main

const (
	openBra     string = "("
	closeBra    string = ")"
	not         string = "!"
	and         string = "+"
	or          string = "|"
	xor         string = "^"
	operators   string = and + or + xor + not
	imp         string = "=>"
	ioi         string = "<=>"
	com         string = "#"
	factSymbol  string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	factDeclar  string = "="
	queryDeclar string = "?"
	trueF       int    = 1
	falseF      int    = 0
	unknownF    int    = -1
)

type nodeInfo int

const (
	noInfo           nodeInfo = 1
	skipClimbUp      nodeInfo = 2
	rightAssociative nodeInfo = 3
)

type precedence int

const (
	openBraOp  precedence = 1
	closeBraOp precedence = 1
	orOp       precedence = 2
	xorOp      precedence = 3
	andOp      precedence = 4
	notOp      precedence = 5
)

type infTree struct {
	head       string
	left       *infTree
	right      *infTree
	operator   string
	precedence precedence
}

var env struct {
	rules        []string
	initialFacts []string
	queries      []string
	allFacts     []string
	tree         infTree
}

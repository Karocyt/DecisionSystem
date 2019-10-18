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
	openBraPre  precedence = 1
	closeBraPre precedence = 1
	impPre      precedence = 1
	ioiPre      precedence = 1
	orPre       precedence = 2
	xorPre      precedence = 3
	andPre      precedence = 4
	notPre      precedence = 5
	factPre     precedence = 6
)

type infTree struct {
	head       *infTree
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
	trees        []infTree
}

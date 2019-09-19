package main

const (
	openPar     string = "("
	closePar    string = ")"
	not         string = "!"
	and         string = "+"
	or          string = "|"
	xor         string = "^"
	imp         string = "=>"
	ioi         string = "<=>"
	com         string = "#"
	factSymbol  string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	factDeclar  string = "="
	queryDeclar string = "?"
)

var env struct {
	kb        []string
	trueFacts []string
	queries   []string
}

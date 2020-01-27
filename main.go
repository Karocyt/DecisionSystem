package main

import (
	"errors"
	"fmt"
	"github.com/Karocyt/expertsystem/internal/lexer"
	"github.com/Karocyt/expertsystem/internal/solver"
	"io/ioutil"
	"os"
)

/*

In print result: need to compute and take
fingerprint of output for every order to check if unconsistencies

In builder.go/eval_rules: need to do the same.

Also need to propagate the name of the variable further to check if it's not referring itself later on.

For this purpose, needs to copy the map in a loop as:
for x, y in map, newmap[x] = y



*/

func getInput() (string, error) {
	var s string

	if len(os.Args) == 2 {
		buf, e := ioutil.ReadFile(os.Args[1])
		s = string(buf)
		return s, e
	} else if len(os.Args) > 2 {
		return s, errors.New(fmt.Sprintf("Usage: %s {filename}", os.Args[0]))
	} else {
		return s, errors.New(fmt.Sprintf("Stdin mode ? Not Today.\nUsage: %s {filename}", os.Args[0]))
	}
	return s, nil
}

func print_result(b *solver.Builder) {
	fmt.Println("Results: {")
	defer fmt.Println("}")
	for _, s := range b.Queries {
		val, e := b.Eval_rules(s)
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Printf("%s = %t\n", s, val)
		}
	}
}

func main() {
	content, e := getInput()
	if e != nil {
		fmt.Println(e)
		return
	}
	l := lexer.New(content, os.Args[1])
	if l.Error != nil {
		fmt.Println(e)
		return
	}
	s, e := solver.New(l.Tokens)
	if l.Error != nil {
		e = l.Error
	}
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("\nQueries:\t", s.Queries, "\nVariables:\t", s.Variables)
	print_result(&s)
	return
}

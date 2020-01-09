package main

import (
	"errors"
	"fmt"
	"github.com/Karocyt/expertsystem/internal/lexer"
	"github.com/Karocyt/expertsystem/internal/parser"
	"io/ioutil"
	"os"
)

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

func print_result(b *parser.Builder) {
	fmt.Printf("Results:\n")
	var empty parser.Node
	for _, s := range b.Queries {
		if b.Rules[s] != empty {
			res, e := b.Rules[s].Eval(s)
			if e != nil {
				fmt.Println(e)
				return
			}
			fmt.Printf("%s = %t\n", s, res)
		} else {
			fmt.Printf("%s = %t (Default)\n", s, false)
		}
	}
}

func main() {
	content, e := getInput()
	if e != nil {
		fmt.Println(e)
		return
	}
	l, e := lexer.New(content, os.Args[1])
	if e != nil {
		fmt.Println(e)
		return
	}
	b, e := parser.New(l.Tokens)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Nodes:\t\t", b.Rules, "\nQueries:\t", b.Queries)
	print_result(&b)
	return
}

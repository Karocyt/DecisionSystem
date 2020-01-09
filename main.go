package main

import (
	"errors"
	"fmt"
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

func main() {
	input, e := getInput()
	if e != nil {
		fmt.Println(e)
		return
	}
	nodes, e := parser.New(input, os.Args[1])
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(nodes)
	return
}

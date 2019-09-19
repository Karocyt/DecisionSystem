package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Need only one argument as file name.\n")
		os.Exit(1)
	}
	parser()
}

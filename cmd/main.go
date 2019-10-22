package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 { // dynamic ruleset
		parseDynamic()
	} else if len(os.Args) == 2 { // file ruleset
		parseFile(os.Args[1])
	} else { // error
		fmt.Println("Error. Retry later ...")
		os.Exit(1)
	}

	engine()
}

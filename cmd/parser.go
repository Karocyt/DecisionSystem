package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 * Parse file and initialize the env global variable
 */
func parser(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		statement := strings.Trim(strings.Split(scanner.Text(), com)[0], " \t\n")
		if statement == "" {
			continue
		}
		if strings.HasPrefix(statement, factDeclar) {
			env.initialFacts = strings.Split(strings.TrimPrefix(statement, factDeclar), "")
		} else if strings.HasPrefix(statement, queryDeclar) {
			env.queries = strings.Split(strings.TrimPrefix(statement, queryDeclar), "")
		} else {
			env.rules = append(env.rules, statement)
		}
	}
	fmt.Printf("rules : %q\n", env.rules)
	fmt.Printf("initialFacts : %q\n", env.initialFacts)
	fmt.Printf("queries : %q\n", env.queries)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	initAllFacts()
	buildTree()
	for _, tree := range env.trees {
		fmt.Printf("\nROOT : \n----------------------------\n")
		printNode(&tree, 4)
	}
}

/*
 * Initialize env.allFacts all mentioned facts from file statements
 */
func initAllFacts() {
	// list from initial facts
	env.allFacts = make([]string, len(env.initialFacts))
	copy(env.allFacts, env.initialFacts)

	// list from query facts
	for _, fact := range env.queries {
		if !stringInSlice(fact, env.allFacts) {
			env.allFacts = append(env.allFacts, fact)
		}
	}

	// list from statement facts
	for _, statement := range env.rules {
		for _, stmt := range statement {
			if !stringInSlice(string(stmt), env.allFacts) && stringInSlice(string(stmt), strings.Split(factSymbol, "")) {
				env.allFacts = append(env.allFacts, string(stmt))
			}
		}
	}
}

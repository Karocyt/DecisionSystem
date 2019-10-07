package main

import (
	"fmt"
	"strings"
)

/*
 * infTree structure constructor
 */
func newInfTree() *infTree {
	var t infTree
	return &t
}

/*
 * add head on last tree node
 */
func initInfTreeHead(nodeList []infTree, fact string) {
	nodeList[len(nodeList)-1].head = fact
}

/*
 * add children on last tree node
 */
func initInfTreeChildren(node []infTree, rule string) {
	left := strings.Split(rule, imp)[0]
	children := strings.Fields(left)
	fmt.Printf("%v\n", children)
}

/*
 * add operators on last tree node
 */
func initInfTreeOperators(node []infTree, rule string) {
	left := strings.Split(rule, imp)[0]
	children := strings.Fields(left)
	fmt.Printf("%v\n", children)
}

// /*
//  * Build the inference tree with all facts and statements
//  */
// func buildTree1() {
// 	var nodeList []infTree

// 	// define nodes and their heads
// 	for _, fact := range env.allFacts {
// 		nodeList = append(nodeList, *newInfTree())
// 		initInfTreeHead(nodeList, fact)
// 	}

// 	// define nodes' children and operators
// 	for _, node := range nodeList {
// 		for _, rule := range env.rules {
// 			right := strings.Split(rule, imp)[1]
// 			if strings.Contains(right, node.head) {
// 				initInfTreeChildren(nodeList, rule)
// 				//initInfTreeOperations(nodeList, rule)
// 			}
// 		}
// 	}
// 	fmt.Printf("%q\n", env.allFacts)
// }

/*
 * Build the inference tree with all facts and statements
 * https://www.rhyscitlema.com/algorithms/expression-parsing-algorithm/
 */
func buildTree() {
	var root *infTree
	var node *infTree
	var current *infTree
	var nodePre precedence
	var info nodeInfo
	var previousPre precedence

	root = newInfTree()
	root.precedence = 1
	root.operator = openBra

	current = root

	for _, rule := range env.rules {
		for _, c := range rule {
			info = noInfo
			root = recursiveBuild(root, string(c))

		}
	}
}

func recursiveBuild(root *infTree, c string) *infTree {
	var next *infTree
	var current *infTree

	if root == nil {
		root = newInfTree()
		root.precedence = 1
		root.operator = openBra
		next = root
	}
	if string(c) == openBra {
		node.head = openBra
		node.precedence = openBraOp
		info = skipClimbUp
	}

	return next
}

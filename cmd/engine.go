package main

import "fmt"

func engine() {
	for _, tree := range env.trees {
		fmt.Printf("%+v\n", tree)
	}
}

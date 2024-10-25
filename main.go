package main

import (
	"fmt"
)

var PersonGreeted string = ""

// func to evaluate in "Debug Console REPL"
func getEvenNumbers(max int) (res []int) {
	for i := range max {
		if i%2 == 0 {
			res = append(res, i)
		}
	}
	return
}

func main() {
	fmt.Println("Type your name:")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hi, %s!\n", name)
	PersonGreeted = name
	getEvenNumbers(3) // this call is mandatory
}

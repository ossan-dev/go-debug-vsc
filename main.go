package main

import "fmt"

func main() {
	fmt.Println("Type your name:")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hi, %s!\n", name)
}

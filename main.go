package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, world!")
	a := 13
	b := 45
	c := add(a, b)
	fmt.Println("a + b = ", c)
}

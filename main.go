package main

import (
	"fmt"
	mymath "git-study/my_math"
	mytest "git-study/test"
)

func main() {
	fmt.Println("Hello, world!")

	a := 13
	b := 45
	c := mymath.Add(a, b)
	fmt.Println("a + b = ", c)
	a = 10
	b = 20

	c = mymath.Mult(a, b)
	fmt.Println("a * b = ", c)

	c = mymath.Sub(a, b)
	fmt.Println("a - b = ", c)
	a = 5
	b = 2
	c = mymath.Div(a, b)
	fmt.Println("a / b = ", c)
	mytest.Test()
}

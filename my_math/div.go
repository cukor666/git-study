package mymath

import "fmt"

func Div(a, b int) int {
	if b == 0 {
		fmt.Println("Error: division by zero")
		return 0
	}
	return a / b
}

package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 4
	c := 5

	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	mod := c % b

	fmt.Println("Sum: ", sum)
	fmt.Println("Sub: ", sub)
	fmt.Println("Mul: ", mul)
	fmt.Println("Div: ", div)
	fmt.Println("Mod: ", mod)
}
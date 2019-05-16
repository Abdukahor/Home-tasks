package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var x int
	fmt.Scanf("%d", &x)

	a := x / 10
	b := x % 10
	c := a + b
	d := a * b
	fmt.Println(a,b,c,d)

}
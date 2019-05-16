package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter n: ")
	var n int
	fmt.Scanf("%d", &n)
	x := n % 12 + 1


	fmt.Println(x)

}
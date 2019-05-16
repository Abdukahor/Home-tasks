package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number 1-365: ")
	var input int
	fmt.Scanf("%d", &input)

	output := input % 7

	fmt.Println(output)

}

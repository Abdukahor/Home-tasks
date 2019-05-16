package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var input int	// given in kilograms
	fmt.Scanf("%d", &input)

	output := input / 100	//calculating centner

	fmt.Println(output)

}
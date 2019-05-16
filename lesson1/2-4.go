package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var input int	// given in meters
	fmt.Scanf("%d", &input)

	output := input / 1000	// calculating kilometers

	fmt.Println(output)

}

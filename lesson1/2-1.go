package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var input int	// given in centimeters
	fmt.Scanf("%d", &input)

	output := input / 100	// calculating meters

	fmt.Println(output)

}

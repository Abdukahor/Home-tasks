package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number 1-365 & d: ")
	var input int
	fmt.Scanf("%d", &input)

	var d int
	fmt.Scanf("%d", &d)

	output := (d + input - 2) % 7 + 1



	fmt.Println(output)

}

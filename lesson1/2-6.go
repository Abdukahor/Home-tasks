package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var input int	// given seconds
	fmt.Scanf("%d", &input)

	hours := input / 360	// calculating hours
	minutes := (input - hours * 360) / 60 // calculating minutes
	seconds := input - (hours * 360 + minutes * 60) // calculating seconds

	fmt.Println(hours)
	fmt.Println(minutes)
	fmt.Println(seconds)
}

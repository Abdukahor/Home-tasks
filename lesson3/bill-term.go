package main

import (
	"fmt"
	"strings"
	"time"
	"crypto/md5"
	"encoding/hex"
)

var (
	in1,in2,in3,in4,in5,in6,in7,in8,dr1,dr2,dr3,dr4,dr5,dr6,dr7,dr8 int
	sortsOfPizza =[]string{"Pizza Regina", "Sicilian pizza", "Pizza Hawaii","Pizza al tonno"}
	drinks =[]string{"Cola", "Fanta", "Juice", "Milk",}
)

func main() {
	pLen := len(sortsOfPizza)

	fmt.Println("Welcome to PizzaMaster!")
	fmt.Println("Choose sort of Pizza (number,amount; e.g: 1,4 3,1 2,3):")

	for i, pizza := range sortsOfPizza { // prints ordered sorts of pizza list
		fmt.Printf("\t%d - %s\n", i+1, pizza)
	}
	fmt.Scanf(strings.Repeat("%d,%d ", pLen), &in1, &in2, &in3, &in4, &in5, &in6, &in7, &in8)

	for in1 < 1 || in1 > pLen || in2 < 1 || in2 > 20 { // checking number and amount of pizza
		fmt.Printf("%s", "Please, choose number/amount of pizza correctly \n")
		fmt.Scanf(strings.Repeat("%d,%d ", pLen), &in1, &in2, &in3, &in4, &in5, &in6, &in7, &in8)
	}

	fmt.Println("What you want to drink?")
	fmt.Println("Choose drinkings (number,amount; e.g: 1,4 3,1 2,3):")

	for i, sort := range drinks { // prints ordered drinks list
		fmt.Printf("\t%d - %s\n", i+1, sort)
	}

	fmt.Scanf(strings.Repeat("%d,%d ", len(drinks)), &dr1, &dr2, &dr3, &dr4, &dr5, &dr6, &dr7, &dr8)

	for dr1 < 1 || dr1 > len(drinks) || dr2 < 1 || dr2 > 20{ //checking number and amount of drinks
		fmt.Printf("%s","Please, enter drinking number/amount correctly \n")
		fmt.Scanf(strings.Repeat("%d,%d ", len(drinks)), &dr1, &dr2, &dr3, &dr4, &dr5, &dr6, &dr7, &dr8)

	}

	cheque()
}

func cheque()  {
	pizPrice :=[]int{6,10,15,12,}
	drPrice :=[]int{1,1,2,1,}

	n1 :=[]int{in1,in3,in5,} // inputed number of pizza
	n2 :=[]int{in2,in3,in6,} // inputed amount of pizza
	currentTime := time.Now().Format("02.01.2006 15:04:05")
	total := 0

	for i, b := range n1 { // prints all choosen pizze,amout and total price
		if b != 0 {
			pizza := sortsOfPizza[b-1]
			price := pizPrice[b-1]*n2[i]
			total = total + price
			fmt.Printf("%s x %d : %d $\n", pizza, n2[i], price)
		}
	}

	d1 :=[]int{dr1,dr3,dr5,dr7,} // inputed number of drink
	d2 :=[]int{dr2,dr4,dr6,dr8,} // inputed amount of drink


	for i, b := range d1 { // prints all choosen drinks,amout and total price
		if b != 0 {
			drink := drinks[b-1]
			price := drPrice[b-1] * d2[i]
			total = total + price
			fmt.Printf("%s x %d : %d $\n", drink, d2[i], price)
		}
	}
	fmt.Println("Customer : Sherzod",)
	fmt.Println("Servicing: Abdukahor")
	fmt.Print(strings.Repeat("-", 32))
	fmt.Println("\ntotal:",total, "$")
	fmt.Println("Date:", currentTime)
	fmt.Println(md5Hash(string(total)))
	fmt.Println(md5Hash(currentTime))
}

func md5Hash(x string) string {
	hasher := md5.New()
	hasher.Write([]byte(x))
	return hex.EncodeToString(hasher.Sum(nil))
}
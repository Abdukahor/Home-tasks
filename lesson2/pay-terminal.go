package main

import (
	"fmt"
	"strconv"
	 "strings"
)

var x,num,transact int
var t,operator string

func main()  {
	list()
}

func list() {
	fmt.Println("    Welcome to our pay terminal!\v")
	fmt.Println("\tChoose your operator:")

	operators :=[]string{"Megafon", "Beeline", "Babilon", "Tcell",}
	err1 := "Please, choose right operator number"

	for i, operator:= range operators{
		fmt.Printf("\t   %d - %s\n", i + 1, operator)
	}

	fmt.Scanf("%d\n", &x)
	for x < 1 || x > len(operators)  {
		fmt.Println(err1)
		fmt.Println(strings.Repeat("-", len(err1))+"\v")
		fmt.Scanf("%d\n",&x)
	}
	operator = operators[ x - 1 ]
	number(operator)

}

func number(operator string) {
	fmt.Printf("Enter your number (%s):", operator)
	fmt.Scanf("%d\n", &num)

	t = strconv.Itoa(num) //converting number to string
	err2 := "Please, enter correct phone number"

	if num == 0 {
		list()
	}else {
		if len(t) !=9 || checkPref() == false  {
			fmt.Println(err2)
			fmt.Println(strings.Repeat("-", len(err2)) + "\v")
			number(operator)
		}else {
			sum()
		}
	}
}

func checkPref() bool {
	megafon :=[]string{"888","904","906","907","909",}
	beeline :=[]string{"915","916","917","918","919",}
	babilon :=[]string{"985","986","987","988","989",}
	tcell   :=[]string{"933","934","935","937","939",}

	var opArr[]string

	switch x {
	case 1: opArr = megafon
	case 2: opArr = beeline
	case 3: opArr = babilon
	case 4: opArr = tcell
	}

	pref := false
	for i:=0; i<len(opArr); i++{

		pref = strings.HasPrefix(t, opArr[i])
		if pref == true {
			break
		}
	}
	return pref
}

func sum() {
	err3 := "Please, enter amount of money properly"

	fmt.Print("Enter sum:")
	fmt.Scanf("%d\n", &transact)

	if transact == 0 {
		number(operator)
	}else {
		if transact < 0 || transact > 1001{
			fmt.Println(err3)
			fmt.Println(strings.Repeat("-", len(err3)) + "\v")
			sum()
		} else {
			success()
		}
	}
}

func success()  {
	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("Operator: %s\n", operator)
	fmt.Printf("Number  : %d\n", num)
	fmt.Printf("Balance : %d\n", transact)
	fmt.Println("Operation Status: Success")
	fmt.Print(strings.Repeat("-", 25))
}

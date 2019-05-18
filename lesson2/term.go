package main

import (
	"fmt"
	"strconv"
	 "strings"
	"time"
)

var ( operatorNum, phoneNum, transact int
		t, operator string
	operators =[]string{"Megafon", "Beeline", "Babilon", "Tcell",} // list of operators
	err =[]string{ "Please, choose right operator number","Please, enter correct phone number","Please, enter sum properly",} // list of errors may occur
)

func main() {
	fmt.Println("    Welcome to our pay terminal!\v")
	fmt.Println("\tChoose your operator:")

	for i, operator:= range operators {  // prints ordered operators list
		fmt.Printf("\t   %d - %s\n", i + 1, operator)
	}

	fmt.Scanf("%d\n", &operatorNum)

	for operatorNum < 1 || operatorNum > len(operators)  {	// checks operators number err
		fmt.Println(err[0])
		fmt.Println(strings.Repeat("-", len(err[0])) + "\v")
		fmt.Scanf("%d\n",&operatorNum)
	}
	operator = operators[ operatorNum - 1 ]
	number()
}

func number() {
	fmt.Printf("|%s| Enter your number (0 to go back):\n", operator)
	fmt.Scanf("%d\n", &phoneNum)

	t = strconv.Itoa(phoneNum) //converting number to string

	if phoneNum == 0 {	// go back to prev menu
		main()
	}else {
		if len(t) !=9 || checkPref() == false {  // checking phoneNumber err
			fmt.Println(err[1])
			fmt.Println(strings.Repeat("-", len(err[1])) + "\v")
			number()
		}else {
			sum()
		}
	}
}

func checkPref() bool { // checking prefix of selected operator
	opArr := map[string][]string{
		"Megafon" : {"888","904","906","907","909",},
		"Beeline" : {"915","916","917","918","919",},
		"Babilon" : {"985","986","987","988","989",},
		"Tcell"   : {"933","934","935","937","939",},
	}

	for _, num1 := range opArr[operator]{

		if strings.HasPrefix(t, num1){
			return true
		}
	}
	return false
}

func sum() {

	fmt.Print("Enter sum (0 to go back):")
	fmt.Scanf("%d\n", &transact)

	if transact == 0 { // go back to prev menu
		number()
	}else {
		if transact < 0 || transact > 1001 { // checking sum err
			fmt.Println(err[2])
			fmt.Println(strings.Repeat("-", len(err[2])) + "\v")
			sum()
		}else {
			success()
		}
	}
}

func success()  {
	currentTime := time.Now().Format("02.01.2006 15:04:05")

	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("Operator: %s\n", operator)
	fmt.Printf("Number  : %d\n", phoneNum)
	fmt.Printf("Balance : %d\n", transact)
	fmt.Println("Date:" , currentTime)
	fmt.Println("Operation Status: Success")
	fmt.Print(strings.Repeat("-", 25))
}
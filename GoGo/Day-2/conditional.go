package main

import (
	"fmt"
	"time"
)

func main() {
	//conditionalStatements()

	today := time.Now()
	fmt.Println(today)
	todate := today.Day()
	fmt.Println(todate)

}

func conditionalStatements() {
	ifELse()
	nestedIfElse()
}
func ifELse() {
	// program to get even and odd numbers
	var x int = 10
	if x/2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("The number is odd")
	}
}

func nestedIfElse() {
	var y = 89
	if y/2 == 0 {
		fmt.Println("the number is not even")
	} else if y/2 == 1 {
		fmt.Println("the number equal to the numeber given")
	} else {
		fmt.Println("Odd hai bhhai")
	}
}

func switchStatememts() {
	var a int = 1
	switch a {
	case 1:
		fmt.Println("Its 1")
		break
	case 2:
		fmt.Println("Its 2")
		break
	default:
		fmt.Println("Into default")
		break
	}
}

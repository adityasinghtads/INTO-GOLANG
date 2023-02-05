package main

import (
	"fmt"
)

func main() {
	// firsttype()
	// secongType()
	// thirdType()
	//breakContinue()
	doWhileinGo()
}

// there are three types of loops in Golang

func firsttype() {
	// Using all three inputs
	sum := 0
	for i := 0; i < 11; i++ {
		sum = sum + i
	}
	//here in the for statement
	// first stands for init statement, second is condition computing
	// third is after condition.
	fmt.Println(sum)
}

func secongType() {
	// while loop in Golang
	n := 1
	for n < 5 {
		n = n * 2
		// above can also be written as
		n *= 2
	}
	// in above case if we skip all the functions it will lead to infinite loop.

}

func thirdType() {
	// for each loop
	strings := []string{"Hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}

// working on break and continue statements

func breakContinue() {
	//sub := 0
	// for i := 0; i < 10; i++ {
	// 	if i%2 != 0 {
	// 		break
	// 	} else {
	// 		fmt.Println("Hey it is an even number")
	// 	}
	// }
	x := 10
	if x%2 == 0 {
	} else {
		fmt.Println("its even")
	}
}

// we have break and continue only in case of for loop.

// we dont have a do while loop in golang.
// same can be recreate / dowhile in golang

func doWhileinGo() {
	for {
		i := 0
		fmt.Println(i)
		i = i + 1
		if i == 3 {
			break
		} else {
			continue
		}
	}
}

// Do while can be done using an if statement to break the infinite statement.

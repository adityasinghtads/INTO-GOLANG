package main

import "fmt"

// ********** Defer statements *********
// defer is the keyword which means the code or lne will run at the end of the program after all others are done

func mul(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("start")
	//defer mul(3, 2)
	defer fmt.Println(mul(3, 2))
	fmt.Println(mul(2, 2))
	//mul(2, 2)
	fmt.Println("End")
}

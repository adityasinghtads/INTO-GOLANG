package main

import (
	"fmt"
)

func main() {
	forLoop()
	//forInfinite()
	forAsWhile()
	forWithRange()
	forWithRangeString()
}

func forLoop() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Looping from :", i)
		i++
	}
}

func forInfinite() {
	for {
		fmt.Println("I am Aditya Singh.")
	}
}

func forAsWhile() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func forWithRange() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, j := range arr {
		fmt.Print(i, "  ", j)
	}

}

func forWithRangeString() {
	for index, char := range "Aditya" {
		fmt.Println("The Index value and Char", index, char)
	}
}

// there is no do-while in Golang, but we can replicate it in this way using for and if...
func doWhile() {
	for {
		i := 0
		fmt.Println(i)
		i = i + 1
		if i < 3 {
			break
		}
	}
}

func forUsingChannel() {

}

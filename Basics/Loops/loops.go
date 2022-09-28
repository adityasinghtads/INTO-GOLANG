package main

import (
	"fmt"
	"log"
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

// using Underscore in loops in Golang - which basically determines that the same can be left empty.. as shown in the below example

func forWithUnderscore() {

	animals := []string{"ant", "cat", "dog", "Horse"}

	animals1 := make(map[string]string)
	animals1["dog"] = "Koko"

	var firstLine = "Once upon a time"

	for i, l := range firstLine {
		log.Println(i, ": ", l)
	}

	for _, animal := range animals1 {
		log.Println(animal)
	}

	log.Println(animals)

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

}

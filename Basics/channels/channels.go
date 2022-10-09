package main

import (
	"log"

	"github.com/adityasinghtads/myniceprogram/helpers"
)

const numPool = 100

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)
	// Good practice to complete with Closing a Defer

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}

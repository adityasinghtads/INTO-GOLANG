package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	// same are Arrays..
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	iterateSlice()

	// in Go we dont use array we use slice
	var mySlice []string

	mySlice = append(mySlice, "Singh")
	mySlice = append(mySlice, "Dhanraj")
	mySlice = append(mySlice, "Aditya")

	log.Println(mySlice)

	sort.Ints(slice)

	// The other way to assign slice is

	uslice := []int{1, 2, 3, 4, 5, 6}
	log.Println(uslice)

	log.Println(uslice[0:2])

}

func iterateSlice() {
	myslice := []int{12, 23, 34, 45, 56, 67, 78, 89}
	for index, element := range myslice {
		fmt.Println(index, " at =>", element)
	}
}

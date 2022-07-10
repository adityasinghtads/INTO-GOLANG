package main

import "fmt"

func main() {
	// same are Arrays..
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	iterateSlice()

}

func iterateSlice() {
	myslice := []int{12, 23, 34, 45, 56, 67, 78, 89}
	for index, element := range myslice {
		fmt.Println(index, " at =>", element)
	}
}

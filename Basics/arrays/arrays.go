package main

import "fmt"

func main() {

	declaringArray()
	multiDimentionalArray()
	print2dArray()
	lengthOfarray()
	//ellipsis in an array - the number of elements are not specidied in before
	ellipsisArray()
	iterateArray()
}

func declaringArray() {
	//delaring an array in Go lang
	myarr := [3]string{"aditya", "is", "TADS"}
	fmt.Println(myarr)

	// Delaring another way
	var myarr1 [3]string
	myarr1[0] = "aditya"
	myarr1[1] = "adit"
	myarr1[2] = "adi"
	fmt.Println(myarr1)

	// 2 Dimentional arrays
	arr1 := [2][2]int{{1, 2}, {2, 1}}
	fmt.Println(arr1)

	// Incase we dont assign any valu to the array specified the default value is taken as zero
	var arr21 string
	fmt.Println(arr21)

	var arr123 [2]string
	fmt.Println(arr123)
}

func multiDimentionalArray() {
	// can be of any dimention
	var array [3][3][3]int
	fmt.Println(array)

}

func print2dArray() {
	fmt.Println("we are into 2d array")
	//var array[2][2]int -- lengthy way
	array := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(array)
}

func lengthOfarray() {
	// taking the length of array..
	array := [3]int{1, 2, 3}
	fmt.Println(len(array))

	// length of empty array
	var array1 [2]int
	fmt.Println(len(array1))
}

func ellipsisArray() {
	myarray := [...]string{"aditya", "adi", "singh"}
	fmt.Println(myarray)
	fmt.Println(len(myarray))
}

func iterateArray() {
	myarray := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for x := 0; x < len(myarray); x++ {
		fmt.Printf("%d\n", myarray[x])
	}
}

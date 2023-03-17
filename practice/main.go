package main

import (
	"fmt"
	"reflect"
)

func main() {
	itsAditya()
	var tads [3]int
	tads = declArray()
	compareArray(tads)
}

func itsAditya() {
	var aditya string = "Aditya"
	fmt.Println("Hey hello")
	fmt.Println(aditya)
	Aditya()
}

func Aditya() string {
	var a string = "Hey this is aditya"
	fmt.Println(a)
	if reflect.TypeOf(a).Kind() == reflect.String {
		fmt.Println("this is a string")
	} else {
		fmt.Println("this is not a string")
	}
	return a
}

func adityaSingh() {

}

func datatypes() {
	var a int = 1
	var b float64 = 1.0
	var abc bool = true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(abc)
}

// Learning Arrays in Golang.

func declArray() [3]int {
	myarr := [...]int{1, 2, 3}
	fmt.Println(myarr)

	//declaring the other way
	var arr1 [3]string
	arr1[0] = "Aditya"
	arr1[1] = "Aditya"
	arr1[2] = "Aditya"

	fmt.Println(arr1)
	fmt.Println(len(arr1))

	//Iterate through an array
	myarray := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(myarr); i++ {
		fmt.Println(myarray[i])
		// In case we need output in single line
		fmt.Printf("%d", myarray[i])
	}

	return myarr
}

func compareArray() {

}

//// Learning Maps in Golang

func mapsInGolang() {

}

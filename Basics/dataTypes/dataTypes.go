package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a bool = true
	var b int = 1
	var c float32 = 3.14
	var d string = "hi This is Aditya"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func boolean() {

}

func byteSameAsRune() {
	fmt.Println("Byte is same as uint8 and rune")
}

func myRune() {
	r := 'a'
	fmt.Printf("Size of %d\n", unsafe.Sizeof(r))
	fmt.Printf("Size of %s\n", reflect.TypeOf(r))
	fmt.Printf("Size of %c\n", r)

	// using rune
	s := "0b£"
	fmt.Printf("%U\n", []rune(s))
	fmt.Println([]rune(s))
}

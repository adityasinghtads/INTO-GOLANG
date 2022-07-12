package main

import "fmt"

func main() {
	pointerIntro()
}

func pointerIntro() {

	b := 8
	a := &b
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)
}

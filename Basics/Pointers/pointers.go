package main

import "fmt"

func main() {
	//pointerIntro()
	pointerToFunc()
	pointerToStruct()
}

// ⁡⁣⁢⁣********** Intro Pointers in Go *********⁡

func pointerIntro() {
	b := 8
	a := &b
	fmt.Println(a)  // address of b
	fmt.Println(&a) // address of a
	fmt.Println(b)  // b value
	fmt.Println(&b) // address of b
	fmt.Println(*a) // value stored at address of a

	z := 10
	p := &z
	fmt.Println(p)
	fmt.Println(*p)

	h := 9
	g := &h
	fmt.Println(h)

	*g = 100
	fmt.Println(h)
}

// ⁡⁣⁢⁣*********** Pointers to function **********⁡

func pointerToFunc() {
	x := 100
	fmt.Println("Value of x before calling:", x)
	var b *int = &x
	callingFunc(b)
	fmt.Println("Value of x after calling:", x)
}
func callingFunc(y *int) {
	*y = 99
}

// *********** Pointers to strct **************
type student struct {
	name string
	age  int
}

func pointerToStruct() {
	stu := student{"Aditya", 2}
	fmt.Println(stu)

	ptr := &stu
	fmt.Println(ptr)
	fmt.Println(ptr.name)
	fmt.Println((*ptr).age)
	ptr.name = "Vedanta"
	fmt.Println("After change name :", ptr.name)
	fmt.Println("After change name :", (*ptr).name)
}

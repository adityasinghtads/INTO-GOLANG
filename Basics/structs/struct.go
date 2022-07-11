package main

import "fmt"

func main() {
	structIntro()
	accessUsingDot()
	pointerToStuct()
	nestedStruct()
	ananymonusStruct()
}

// ⁡⁣⁢⁣************** Structure Intro *************⁡

func structIntro() {
	type Address struct {
		name    string
		street  int
		city    string
		state   string
		pincode int
	}

	type student struct {
		fname, lname   string
		class, pnumber int
	}

	var a Address
	fmt.Println(a)

	b := Address{"Aditya", 2, "Hyderabad", "Telangana", 50059}
	fmt.Println(b)

	c := Address{name: "Aditya"}
	fmt.Println(c)

}

// ⁡⁣⁢⁣************ Struct access using Dot operator *****************⁡

func accessUsingDot() {
	type Address struct {
		name    string
		street  int
		city    string
		state   string
		pincode int
	}
	a := Address{name: "Vedanta", street: 23}
	println(a)

	fmt.Println("Name:", a.name)
	fmt.Println("Street", a.street)
}

// ⁡⁣⁢⁣************ Poiner to Struct *****************⁡

func pointerToStuct() {
	type Address struct {
		name    string
		street  int
		city    string
		state   string
		pincode int
	}

	emp1 := &Address{name: "Aditya Singh"}
	fmt.Println("Name is ", (*emp1).name)
	// Other way to define this is
	fmt.Println("name is ", emp1.name)
	// Checking out other things as well
	fmt.Println("Complete", emp1)

}

// ⁡⁣⁢⁣************ Nested Structure *****************⁡

func nestedStruct() {
	type subject struct {
		name  string
		hours int
	}
	type teacher struct {
		name    string
		age     int
		details subject
	}

	Manya := teacher{"Manya", 25, subject{"science", 25}}
	fmt.Println(Manya)

	Divya := teacher{name: "divya", age: 26, details: subject{"maths", 26}}
	fmt.Println(Divya)
}

func ananymonusStruct() {
	// rule for anonymous Struct is the value for the struct should be defined immedatiely afte the struct.
	aditya := struct {
		name string
		age  int
	}{
		name: "TADS",
		age:  22,
	}
	fmt.Println(aditya)
}

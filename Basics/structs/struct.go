package main

import "fmt"

func main() {
	structIntro()
	accessUsingDot()
}

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

func accessUsingDot() {
	type Address struct {
		name    string
		street  int
		city    string
		state   string
		pincode int
	}

}

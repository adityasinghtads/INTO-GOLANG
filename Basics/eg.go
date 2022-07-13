package main

import "fmt"

func main() {
	var f int
	fmt.Scanln(&f)
	//your code goes here
	switch f {
	i:=f 
	case i>= 20:
		fmt.Println("Audible")
	case i < 20:
		fmt.Println("Infrasound")
	case i > 20000:
		fmt.Println("Ultrasound")
	case i < 0:
		fmt.Println("Wrong Input")
	}
}

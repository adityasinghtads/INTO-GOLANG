package main

import "fmt"

type Coord struct {
	x int
	y int
}

func main() {
	a := Coord{7, 5}
	fmt.Println(a.x - a.y)
}

// ********** Salary using struct ***********

// type Employee struct {
// 	name     string
// 	age      int
// 	position string
// 	salary   float32
// }

// func main() {
// 	e1 := Employee{"James", 42, "Manager", 0}

// 	var x float32
// 	fmt.Scanln(&x)
// 	e1.salary = x

// 	fmt.Println("========")
// 	fmt.Println(e1.name + "(" + e1.position + ")")
// 	fmt.Println(e1.age)
// 	fmt.Println(e1.salary)
// 	fmt.Println("========")
// }

// package main

// import "fmt"

// func main() {
// 	var num float32
// 	var factor float32

// 	fmt.Scanln(&num)
// 	fmt.Scanln(&factor)

// 	scale(&num, factor)
// 	fmt.Println(num)
// }
// func scale(a *float32, b float32) {
// 	*a = *a * b
// }

// //************** pointes general *******
// // package main

// // import "fmt"

// // func main() {
// // 	var x int
// // 	fmt.Scanln(&x)
// // 	fmt.Println(&x)
// // }

// // ******** Pointers *******
// // func main() {
// // 	a := 4
// // 	p := &a
// // 	a += 2
// // 	*p = *p - 1
// // 	fmt.Println(a)
// // }
// // func mars_age(a int) int {
// // 	daysOnEarth := a * 365
// // 	yearsOnMars := daysOnEarth / 687
// // 	return yearsOnMars
// // }

// // *********  MARS AGE *************
// // func main() {
// // 	var age int
// // 	fmt.Scanln(&age)

// // 	mars := mars_age(age)
// // 	fmt.Println(mars)
// // }
// // func mars_age(a int) int {
// // 	daysOnEarth := a * 365
// // 	yearsOnMars := daysOnEarth / 687
// // 	return yearsOnMars
// // }

// // package main

// // import "fmt"

// // func double(a int) {
// // 	fmt.Println(a * 2)
// // }
// // func main() {
// // 	for i := 4; i > 0; i-- {
// // 		defer double(i)
// // 	}
// // }

// // func main() {
// // 	var f int
// // 	fmt.Scanln(&f)
// // 	//your code goes here
// // 	switch {
// // 	case f < 0:
// // 		fmt.Println("Wrong Input")
// // 	case f > 20000:
// // 		fmt.Println("Ultrasound")
// // 	case f < 20:
// // 		fmt.Println("Infrasound")
// // 	default:
// // 		fmt.Println("Audible")
// // 	}
// // }

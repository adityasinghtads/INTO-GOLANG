package main

import (
	"fmt"
	"log"

	"github.com/adityasinghtads/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
	fmt.Println(myVar.TypeName)
}

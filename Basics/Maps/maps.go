package main

import (
	"fmt"
	"log"
)

func main() {
	mapIntro()
	mapUsingMake()
	updateMap()
	retriveMapData()
	udemymaps()
	mapsAndStruct()

	// right way to use map
	myMap := make(map[string]string)
	myMap["dog"] = "samson"
	myMap["other-dog"] = "cassie"

	myMap["dog"] = "KOKO"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

}

func mapsAndStruct() {

	type User struct {
		FirstName string
		LastName  string
	}

	myMap := make(map[string]User)

	me := User{
		FirstName: "Aditya",
		LastName:  "Singh",
	}

	myMap["me"] = me
	log.Println(myMap["me"].FirstName)
	log.Println(myMap["me"].LastName)

	var myNewVar float32

	myNewVar = 11.1

	//Maps are immutable

	log.Println(myNewVar)

}

func udemymaps() {
	myMap := make(map[string]int)
	myMap["first"] = 1
	myMap["second"] = 2

	log.Println(myMap["first"], myMap["second"])
}

func mapIntro() {
	var map1 map[int]int
	if map1 == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	map2 := map[int]string{
		10: "dog",
		11: "cat",
		12: "Adi",
		13: "singh",
	}
	fmt.Println("Map1:", map2)
}

func mapUsingMake() {
	var mmap = make(map[int]int)
	fmt.Println(mmap)
	mmap[1] = 1
	mmap[2] = 2
	fmt.Println(mmap)

	// iterating in MAP
	for id, value := range mmap {
		fmt.Println(id, value)
	}
}

func updateMap() {
	my_map := map[int]string{
		1: "adi",
		2: "singh",
		3: "Aditya",
		4: "Birds",
	}
	fmt.Println("Original Map", my_map)
	//update
	my_map[1] = "parrot"
	my_map[2] = "hi"
	//add
	my_map[8] = "eight"
	my_map[9] = "nine"

	fmt.Println("Maps after Updation", my_map)
}

func retriveMapData() {
	var mymap map[int]string
	mymap[1] = "one"
	mymap[2] = "two"
	mymap[3] = "three"
	mymap[4] = "four"

	value1 := mymap
	fmt.Println(value1)
}

func iterateMap() {
	mymap := map[int]string{
		1: "Aditya",
		2: "Singh",
		3: "TADS",
	}
	for key, element := range mymap {
		fmt.Println(key, "=>", element)
	}
}

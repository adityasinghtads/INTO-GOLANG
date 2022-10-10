package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home Page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is at the about Pafe and 2+2 is %d", sum))
}

// addValues adds two intergers and returns sum
func addValues(x, y int) int {
	return x + y
}

//main is the man application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/", About)

	fmt.Println(fmt.Sprintf("Starting application on Port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello world!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number f bytes writen: %d", n))
	// })

}

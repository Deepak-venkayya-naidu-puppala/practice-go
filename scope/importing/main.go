package main

import "fmt"

// import "fmt" // cannot import again 

import format "fmt" // aliasing

func main() {
	fmt.Println("hello world frm fmt")
	format.Println("hello world from format")
}
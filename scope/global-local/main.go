package main

import "fmt"

func util() {
	value:=1000
	fmt.Println("Value of value inside util: %v", value)
}

var value = 2000

func main() {
	fmt.Println("inside main %v", value)

	util()

	value:=3000

	fmt.Println("Inside main after value is decalred and defined %v ", value)
}
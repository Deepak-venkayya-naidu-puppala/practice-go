package main 

import "fmt"

var value = 100

func main() {
	 fmt.Println(value)
	 util()
}

func util() {

	fmt.Println(" from util in main.go....")
}
package main


import "fmt" //file scoped

const bool_value = true //package scoped

func main() {
	in_main := true //block scoped
	fmt.Println(bool_value)
	fmt.Println(is_accessable)
	fmt.Println(in_main)
	//fmt.Println(integer) can't be accessed
}
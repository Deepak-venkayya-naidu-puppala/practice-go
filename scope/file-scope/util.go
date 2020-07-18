package main

var is_accessable  = true // package scoped

func function() {
	integer:=10 
	integer+=1 //block scoped

	// fmt.Println(integer)
}
package main

import "fmt"

func main() {
	fmt.Println("Hello");

	fmt.Println("world!"); // go supports semicolons;

	if 5 > 1 {
		fmt.Println("5 is greater than 1")
	}

	if 5 > 6 {
		fmt.Println("5 is greater than 6");
	} else {
		fmt.Println("5 is not greater than 6");
	}

	if 5 > 5 {
		fmt.Println("5 is greater than 5")
	} else if 5 < 5 {
		fmt.Println("5 is less than 5")
	} else {
		fmt.Println("5 is neihter less than nor greater than 5")
	}
}




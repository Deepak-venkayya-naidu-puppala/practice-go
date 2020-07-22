package main


import "fmt"

func main() {
	a, b, c := 1000, 2000, 3000
	brand := "AX"
	var branch string;
	
	fmt.Println("Printing using Println .....:")
	fmt.Println(a, b, c)
	fmt.Println("brand:", brand)
	fmt.Println("Branch:", branch)

	fmt.Printf("Printing using Printf ......:")
	fmt.Printf("%d %d %d", a, b, c)
	fmt.Printf("Brand: %s", brand)
	fmt.Printf("Branch: %s", branch)

// ---------> Printing types
	fmt.Printf("Type of a, b, c are %T, %T, %T \n", a, b, c);
	fmt.Printf("Types of brand and branch are : %T, %T\n", brand, branch)


}
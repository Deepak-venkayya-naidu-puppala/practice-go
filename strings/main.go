package main

import ( 
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println("String")
	fmt.Println("Raw string literal")

	fmt.Println("String 1" + " concatenating with " + "string 2")

	fmt.Println("ASCII value of 3 is " + strconv.Itoa(3))

	eq := strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)

	fmt.Println("lengh of string \"Hello world\" :" , len("Hello world"))


	fmt.Println("operations on unicode string literal..")
	// len function counts the bytes in a string value.
	//
	// This string literal contains unicode characters.
	//
	// And, unicode characters can be 1-4 bytes.
	// So, "İnanç" is 7 bytes long, not 5.
	//
	// İ = 2 bytes
	// n = 1 byte
	// a = 1 byte
	// n = 1 byte
	// ç = 2 bytes
	// TOTAL = 7 bytes
	name := "İnanç"
	fmt.Printf("%q is %d bytes\n", name, len(name))

	// To get the actual characters (or runes) inside
	// a utf-8 encoded string value, you should do this:
	fmt.Printf("%q is %d characters\n",
		name, utf8.RuneCountInString(name))
}
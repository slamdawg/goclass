package main

import "fmt"

func main() {
	// fmt.Printf("%d - %b - %x", 42, 42, 42)
	// fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	// fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	// fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}

//%d = decimal
//%b = binary
//%x = hexidecimal
//%q = a double-quoted string safely escaped with Go syntax
//%s = the uninterpreted bytes of the string or slice

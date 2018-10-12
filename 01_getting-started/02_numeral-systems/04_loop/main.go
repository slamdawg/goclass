package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}

//%d = decimal
//%b = binary
//%x = hexidecimal
//%q = a double-quoted string safely escaped with Go syntax
//%s = the uninterpreted bytes of the string or slice

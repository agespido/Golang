package main

// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!") // Ends the line with '\n'
	fmt.Printf("Hello Sun!\n")  // Analogous to C/C++ printf()
	fmt.Print("Hello Moon!\n")  // Similar to Println, but it does not print '\n' at the end of the line nor spaces between variables

	fmt.Println("Today is", 13, "December")
	fmt.Printf("Today is %d December\n", 13)
	fmt.Print("Today is", 13, "December\n")
}

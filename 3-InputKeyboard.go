package main

import (
	"bufio"
	"fmt"
	"os"
)

var num1 int
var num2 int
var res int
var text string

func main() {

	// Scan numeric variables
	fmt.Println("Num 1: ")
	fmt.Scanln(&num1)
	fmt.Println("Num 2: ")
	fmt.Scanln(&num2)
	fmt.Println("Tell me something: ")

	// Scan text -> Scanln will not work for text strings
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}

	// Output
	res = num1 + num2
	fmt.Println() // Prints an empty line
	fmt.Println("Scanned text:", text)
	fmt.Println("Result of the addition:", res)

}

package main

import "fmt"

// Global variable declaration
var num int
var text string
var status bool = true

func main() {
	// Declare a single variable
	var num2 int
	// Declare and initialize several variables in just a single line
	num3, num4, text2, status := 4, 10, "Hi there!", false

	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(text2)
	fmt.Println(status)

	showStatus()
}

func showStatus() {
	fmt.Println(status)
}

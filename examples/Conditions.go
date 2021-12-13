package main

import "fmt"

func main() {
	var num int
	fmt.Println("Square a number between 1 and 100: ")
	fmt.Scanln(&num)

	// if-else confition example
	if num2 := num * num; num2 > 100*100 {
		fmt.Println("Your number is too largo!")
	} else if num2 <= 0 {
		fmt.Println("Yout number is too small!")
	} else {
		fmt.Println(num, "squared is", num2)
	}
}

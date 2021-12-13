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

	fmt.Println("Write a number between 1 and 12, and I will tell you which month it is: ")
	fmt.Scanln(&num)
	// Switch condition example
	switch num {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Unknown")
	}

}

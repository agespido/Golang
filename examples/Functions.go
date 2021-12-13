package main

import "fmt"

func main() {
	num := 5
	fmt.Println("The square of", num, "is", square(num))

	num1, num2 := 1, -2
	sum, status := checkSignOfSum(num1, num2)
	fmt.Println("The result of", num1, "+", num2, "=", sum)
	fmt.Println("Is", sum, "greater than 0?", status)

	fmt.Println("Summatory:", sumatory(1))
	fmt.Println("Summatory:", sumatory(4, 6))
	fmt.Println("Summatory:", sumatory(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

// Gets num as an input, and returns z
func square(num int) (z int) {
	z = num * num
	return // No need to specify which variable is returned, as it is specified in th declaration
}

// Given 2 numbers, the function return the resulto of adding them and true if the result is postive, false otherwise.
func checkSignOfSum(num1 int, num2 int) (int, bool) {
	sum := num1 + num2
	if sum > 0 {
		return sum, true
	} else {
		return sum, false
	}
}

// Returns the summatory of a variable number of inputs
func sumatory(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

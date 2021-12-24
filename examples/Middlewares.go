package main

import "fmt"

func main() {
	result := middleware(sum)(6, 2)
	fmt.Println(result)
	result = middleware(sub)(6, 2)
	fmt.Println(result)
	result = middleware(mult)(6, 2)
	fmt.Println(result)
	result = middleware(div)(6, 2)
	fmt.Println(result)
}

func middleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		return f(a, b)
	}
}

func sum(a, b int) int {
	fmt.Println("Addition")
	return a + b
}

func sub(a, b int) int {
	fmt.Println("Substraction")
	return a - b
}

func mult(a, b int) int {
	fmt.Println("Multiplication")
	return a * b
}
func div(a, b int) int {
	fmt.Println("Division")
	return a / b
}

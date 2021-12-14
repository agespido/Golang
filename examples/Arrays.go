package main

import "fmt"

func f(arr ...int) {
	fmt.Println("Array length:", len(arr), "->", arr)
}

func f2(arr []int) {
	fmt.Println("Array length:", len(arr), "->", arr)
}

func main() {

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array length:", len(arr), "->", arr)

	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Array length:", len(arr2), "->", arr2)

	var arr3 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Array length:", len(arr3), "->", arr3)

	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7}
	arr5 := arr4

	arr5[0] = 10
	arr5[1] = 11
	/*
		Unlike in C/C++, when one array is assigned the value of another,
		it does not create a pointer, but a complete copy of the array.
	*/
	fmt.Println("Array length:", len(arr4), "->", arr4)
	fmt.Println("Array length:", len(arr5), "->", arr5)

	// Passing arrays to functions
	f(arr[:]...)
	f2(arr[:])
}

package main

import (
	"fmt"
)

/*
Note:
	"While" and "do-while" loops do not exist in Golang.
	All the loops must be handled by means of "for" loops.
*/

func main() {
	// Classic for loop
	max := 10
	for i := 1; i <= max; i++ {
		fmt.Println(i, "out of", max, "iterations")
	}
}

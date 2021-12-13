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
	num, secret, credits := 0, 100, 5
	// While-style loop
	for {
		fmt.Println("Try to guess the secret number:")
		fmt.Scanln(&num)
		if num == secret {
			fmt.Println("YOU WIN!", num, "is the secret number!")
			break
		}
		fmt.Print(num, " is not the secret number.")
		if credits > 0 {
			credits--
			fmt.Print("Try again!!\n")
			continue
		} else {
			fmt.Println()
			fmt.Println("GAME OVER")
			break
		}

	}
}

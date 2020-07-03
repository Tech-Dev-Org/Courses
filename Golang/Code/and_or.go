package main

import (
	"fmt"
)

func main() {
	x := 10
	//if else
	if x > 5 {
		fmt.Println("X is big")
	} else {
		fmt.Println("x is less")
	}
	// and
	if x > 5 && x < 15 {
		fmt.Println("x is less 15 and greater than 5")
	}
	//or
	if x > 30 || x < 20 {
		fmt.Println("x is out of range")
	}
}

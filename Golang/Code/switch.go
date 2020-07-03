// example of "switch" statement

package main

import (
	"fmt"
)

func main() {
	//using choice
	x := 2
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("many")

	}
	//using conditions
	switch {
	case x > 10:
		fmt.Println("above 10")
	case x > 5:
		fmt.Println("above 5")
	default:
		fmt.Println("below 5")
	}
}

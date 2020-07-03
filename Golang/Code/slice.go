package main

import (
	"fmt"
)

func main() {
	loons := []string{"car", "motor", "truck"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	//length
	fmt.Println(len(loons))

	fmt.Println("------")
	// start 0 indexing
	fmt.Println(loons[1])

	fmt.Println("------")
	//  indexing
	fmt.Println(loons[1:])

	fmt.Println("------")
	//for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("------")
	// single value range
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("------")
	// double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("------")
	// ignore index using _
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("------")
	// append new element
	loons = append(loons, "cycle")
	fmt.Println(loons)

}

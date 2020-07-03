package main

import (
	"fmt"
)

func main() {
	book := "the color of magic"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	//String in go is immutable
	//book[0] = 116

	//slice
	fmt.Println(book[2:9])

	//concat string
	fmt.Println("hello " + book[4:])

	//multiple line
	milti_book := `
	this is first.
	this is second
	`
	fmt.Println(milti_book)

	//convert int to string
	n := 47
	s := fmt.Sprintf("%d", n)
	fmt.Printf("s = %v (type %T)\n", s, s)
	fmt.Printf("s = %q", s) //check string

}

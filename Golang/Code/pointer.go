package main

import "fmt"

func main() {
	x := 1
	var y int
	var ip *int //ip is pointer to x

	ip = &x //ip is now point to x
	y = *ip //y is now 1

	//print value
	fmt.Printf("%d", y)
}

package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"google": 99.2,
		"ibm":    95.7,
		"tcs":    90.6, //must have trailing comma in multi line
	}

	//number of items
	fmt.Println(len(stocks))

	//get value
	fmt.Println(stocks["tcs"])

	//get zero value if no fount
	fmt.Println(stocks["wipro"]) //0

	//use to value to see if found
	value, ok := stocks["wipro"]
	if !ok {
		fmt.Println("Wipro not found")
	} else {
		fmt.Println(value)
	}

	//set
	stocks["wipro"] = 89.5
	fmt.Println(stocks)

	//delete
	delete(stocks, "tcs")
	fmt.Println(stocks)

	//for single value key
	for key := range stocks {
		fmt.Println(key)
	}

	//for double value key , values
	for key, value := range stocks {
		fmt.Println(key, "=", value)
	}

}

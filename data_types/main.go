package main

import (
	"fmt"
)

func main() {
	//INT type
	var i1 int8 = 127
	fmt.Println(i1)
	//Array
	var number = [5]int{1, 2, 3, 4}

	fmt.Println(number)

	fmt.Printf("%T\n", number)

	//slice
	var cities = []string{"hello", "hello1"}
	fmt.Println(cities)

	//MAPS
	var names = map[string]float64{
		"1": 1.05,
		"2": 2.03,
	}

	fmt.Println(names)

}

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

	//STRUC
	type person struct {
		id   int
		name string
	}
	var me person
	me.id = 1
	me.name = "rohit"

	//POINTER
	var x int = 2
	ptr := &x
	fmt.Printf("%T\nvalue %v\n", ptr, ptr)
	fmt.Println(me)

	//FUNCTION

	fmt.Printf("%T and value %v\n", myfunc, myfunc)

}

func myfunc() {

}

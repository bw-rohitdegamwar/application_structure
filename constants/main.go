package main

import "fmt"

func main() {
	const days int = 7
	const pi float64 = 3.14
	var i int
	fmt.Println(i)

	a, b := 5, 1

	fmt.Println(a / b)

	const c, d = 6, 1

	fmt.Println(c / d)
	// conatsnts rule
	// you can not change a constant
	// you can not initalise at runtime
	// constant belong to compile time
	// you can not use a variable to initialise the constant

	const t = len("test")

	fmt.Println(t)

	const p = 6.7
	const k float64 = 6.4

}

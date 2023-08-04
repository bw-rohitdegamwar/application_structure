package main

import "fmt"

func main() {
	var name = "Dan"
	x := 10
	fmt.Println("Started with ", x, name)
	s := "Learing go"
	_ = s

	car, cost := "audi", 50000

	fmt.Println("Car :", car, " Cost: ", cost)

	car, price := "BMW", 5000000

	fmt.Println("Car :", car, " price: ", price)

	var opened = false

	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)
	var a, b, c int
	fmt.Println(a, b, c)

	var i, j = 5, 8

	j, i = i, j // swapping variables

	var sum float64 = 60.9999999999
	sum = sum + 5 + 2.5

	fmt.Println(sum)

	fmt.Println(i, j)

}

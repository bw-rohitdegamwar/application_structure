package main

import "fmt"

func f1() {
	fmt.Println("This is f1() nfunction.")
}

func f2(a int, b int) {
	fmt.Println("Sum : ", a+b)
}

func f3(a, b, c int, d, e float32, f string) {
	fmt.Println(a, b, c, d, e, f)
}

func f4(a, b float32) float32 {
	return a + b
}

func f5(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	f1()
	f2(6, 7)
	f3(1, 2, 3, 4.5, 6.7, "Test")
	fmt.Println(f4(3.2, 4.5))
	c, d := f5(8, 7)
	fmt.Println(c, d)
}

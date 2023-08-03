package main

import "fmt"

const secInHour = 3600

func main() {
	fmt.Printf("Hello Go world\n")
	distance := 60.8 //distance in kilometer
	fmt.Printf("Distance in mile %f \n", distance*0.621)
}

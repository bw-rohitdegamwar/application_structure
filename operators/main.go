package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99)
	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 44.2)

	fmt.Println(myStr)

	var x, err = strconv.ParseFloat("44.899", 64)

	fmt.Println(x, err)

}

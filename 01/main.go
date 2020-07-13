package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if z := x * y; z%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

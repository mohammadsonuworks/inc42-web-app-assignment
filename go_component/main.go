package main

import (
	"fmt"
)

// Add function that takes two integers and returns their sum
func Add(a, b int) int {
	return a + b
}

func main() {
	x, y := 5, 6
	result := Add(x, y)
	fmt.Printf("The sum of %d and %d is %d\n", x, y, result)
}

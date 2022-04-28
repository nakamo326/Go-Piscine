package main

import (
	"fmt"
	"piscine"
)

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
func main() {
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
	fmt.Println(piscine.Fibonacci(6))
	fmt.Println(piscine.Fibonacci(10))
}

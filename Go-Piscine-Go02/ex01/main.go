package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))
	fmt.Println(piscine.RecursiveFactorial(12))
	fmt.Println(piscine.RecursiveFactorial(13))

	fmt.Println(piscine.RecursiveFactorial(20))
	fmt.Println(piscine.RecursiveFactorial(21))
}

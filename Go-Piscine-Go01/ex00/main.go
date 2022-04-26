package main

import (
	"fmt"
	"piscine"
)

func main() {
	n := 0
	piscine.PointOne(&n)
	fmt.Println(n)
	n = 200
	piscine.PointOne(&n)
	fmt.Println(n)
}

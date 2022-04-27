package piscine

import "fmt"

func SortIntegerTable(table []int) {
	size := 0
	for range table {
		size++
	}
	var i, j, k int
	for i = 1; i < size; i++ {
		// check insert position. i) elm to insert, j) pos to insert
		for j = i - 1; j >= 0; j-- {
			if table[j] < table[i] {
				break
			}
		}
		for k = i; k > j+1; k-- {
			table[k], table[k-1] = table[k-1], table[k]
		}
		fmt.Println(table)
	}
}

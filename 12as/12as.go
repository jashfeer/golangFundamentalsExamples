package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Enter array limit ")
	var limit int
	fmt.Scan(&limit)
	var array = make([]int, limit)
	fmt.Println("Enter array values :")
	for i := 0; i < limit; i++ {
		fmt.Scan(&array[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	fmt.Println("Sorted array:", array)
}

package main

import "fmt"

func main() {
	fmt.Println("Enter the limit ")
	var limit int
	fmt.Scan(&limit)
	fmt.Println("Enter array values")
	var array = make([]int, limit)
	for i := 0; i < limit; i++ {
		fmt.Scan(&array[i])
	}
	count := 0
	for i := 0; i < limit; i++ {
		if array[i]%2 == 0 {
			count++
		}
	}
	fmt.Println("Number of even numbers in the given array is :", count)
}

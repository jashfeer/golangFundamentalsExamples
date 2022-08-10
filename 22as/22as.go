package main

import "fmt"

var size, num int
var array1 [20][20]int
var array2 [20][20]int

func main() {
	fmt.Println("Enter array size ")
	fmt.Scan(&size)
	getarray()
	display()
}

func getarray() {
	fmt.Println("Enter array1 values")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&array1[i][j])
		}
	}
	fmt.Println("Enter array2 values")

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&num)
			array2[i][j] = num + array1[i][j]
		}
	}
}
func display() {
	fmt.Println("Sum of 2 array ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(" ", array2[i][j])
		}
		fmt.Println()
	}
}

package main

import "fmt"

var array [10][10]int
var size int

func main() {
	getArray()
	displayArray()
}
func getArray() {
	fmt.Print("Enter array size: ")
	fmt.Scan(&size)
	fmt.Println("Enter array values ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&array[i][j])
		}
	}
}

func displayArray() {
	fmt.Print("\nArray elements are:\n")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {

			fmt.Print(" ", array[i][j])
		}
		fmt.Println()
	}

}

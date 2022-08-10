package main

import "fmt"

func main() {
	fmt.Print("Enter the array limit : ")
	var limit int
	fmt.Scanln(&limit)
	fmt.Println("Enter Array 1 values ")
	var array1 = make([]int, limit)
	for i := 0; i < limit; i++ {
		fmt.Scan(&array1[i])
	}
	fmt.Println("Enter Array 2 values ")
	var array2 = make([]int, limit)
	for i := 0; i < limit; i++ {
		fmt.Scan(&array2[i])
	}
	array1, array2 = array2, array1
	fmt.Println("Arrays after swapping ")
	fmt.Println("Array 1", array1)
	fmt.Println("Array 2", array2)
}

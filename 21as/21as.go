package main

import "fmt"

func main() {
	fmt.Print("Enter the array limit : ")
	var limit int
	fmt.Scan(&limit)
	fmt.Println("Enter array values")
	array1 := make([]int, limit)
	array2 := make([]int, limit-1)
	for i := 0; i < limit; i++ {
		fmt.Scan(&array1[i])
	}
	for i := 0; i < limit-1; i++ {
		array2[i] = array1[i] * array1[i+1]
	}
	fmt.Println("Multipled array is : ", array2)
}

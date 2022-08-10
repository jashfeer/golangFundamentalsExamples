package main

import "fmt"

func main() {
	array := getarray()
	displayarray(array)

}

func getarray() []int {
	fmt.Print("Enter array limit : ")
	var limit int
	fmt.Scan(&limit)
	var ar = make([]int, limit)
	fmt.Println("Enter array values ")
	for i := 0; i < limit; i++ {
		fmt.Scan(&ar[i])
	}
	return ar
}

func displayarray(arr []int) {
	fmt.Println("Enterd array is : ", arr)
}

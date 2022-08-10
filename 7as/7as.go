package main

import "fmt"

func main() {
	fmt.Print("Enter a number :")
	var number int
	fmt.Scan(&number)
	for i := 0; i <= 10; i++ {
		result := i * number
		fmt.Printf("%v X %v = %v\n", i, number, result)
	}
}

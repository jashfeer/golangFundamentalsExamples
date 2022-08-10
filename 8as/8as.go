package main

import "fmt"

func main() {
	fmt.Print("Enter a limit : ")
	var limit int
	fmt.Scan(&limit)
	sum := 0
	for i := 0; i <= limit; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Println("Sum of odd numbers is ", sum)
}

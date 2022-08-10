package main

import "fmt"

func main() {
	var Number1 int
	var Number2 float32
	var sum float32
	fmt.Print("Enter 2 Numbers : ")
	fmt.Scan(&Number1, &Number2)
	sum = float32(Number1) + Number2
	fmt.Print("sum of 2 numbers is : ", sum)
}

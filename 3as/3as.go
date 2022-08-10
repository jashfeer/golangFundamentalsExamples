package main

import "fmt"

func main() {
	var P int
	var R float32
	var n float32
	var SI float32
	fmt.Print("Enter the principal amount : ")
	fmt.Scan(&P)
	fmt.Print("Enter interest rate : ")
	fmt.Scan(&R)
	fmt.Print("Enter number of years : ")
	fmt.Scan(&n)
	SI = (float32(P) * R * n) / 100
	fmt.Print("Simple Interest is : ", SI)
}

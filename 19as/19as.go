package main

import "fmt"

func main() {
	fmt.Print("Enter the annual income :")
	var income float32
	fmt.Scan(&income)
	var tax float32
	if income < 250000 {
		fmt.Println("No tax")
	} else if 250000 < income && income <= 500000 {
		tax = (income * 5) / 100
		fmt.Println("Income tax amount : ", tax)
	} else if 500000 < income && income <= 1000000 {
		tax = (income * 20) / 100
		fmt.Println("Income tax amount : ", tax)
	} else if 1000000 < income && income <= 5000000 {
		tax = (income * 30) / 100
		fmt.Println("Income tax amount : ", tax)
	}
}

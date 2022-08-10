package main

import "fmt"

func main() {
	var number1, number2 float32
	var operation int
	fmt.Printf("Enter two numbers : ")
	fmt.Scanln(&number1)
	fmt.Scanln(&number2)
	fmt.Print("1 for addition\n2 for subtraction\n3 for multiplication\n4 for division\nChoose operation :")
	fmt.Scan(&operation)

	switch operation {
	case 1:
		addition(number1, number2)
	case 2:
		subtraction(number1, number2)
	case 3:
		multiplication(number1, number2)
	case 4:
		division(number1, number2)
	default:
		fmt.Print("Invalid entry!")
	}

}

func addition(number1 float32, number2 float32) {
	fmt.Println("Result: ", number1+number2)
}
func subtraction(number1 float32, number2 float32) {
	fmt.Println("Result: ", number1-number2)
}
func multiplication(number1, number2 float32) {
	fmt.Println("Result: ", number1*number2)
}
func division(number1, number2 float32) {
	fmt.Println("Result: ", number1/number2)
}

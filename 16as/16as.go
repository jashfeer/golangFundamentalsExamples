package main

import "fmt"

func main() {
	fmt.Print("Enter the number : ")
	var number int
	fmt.Scan(&number)
	flag := true
	if number < 2 {
		fmt.Println("Given number is not a prime")
		flag = false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				flag = false
				fmt.Println("Given number is not a prime")
				break
			}
		}
		if flag == true {
			fmt.Println("Given number is a prime")
		}
	}
}

package main

import "fmt"

func main() {
	fmt.Print("Enter any char : ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Entered char is : ", name)
}

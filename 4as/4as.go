package main

import "fmt"

func main() {
	fmt.Print("Enter your mark : ")
	var mark float32
	fmt.Scan(&mark)
	if mark >= 50 {
		fmt.Print("You are passed")
	} else {
		fmt.Print("You are failed")
	}
}

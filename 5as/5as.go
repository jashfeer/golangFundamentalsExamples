package main

import "fmt"

func main() {
	fmt.Print("Enter your mark : ")
	var mark float32
	fmt.Scan(&mark)
	if mark >= 90 {
		fmt.Print("A Grade")
	} else if mark >= 80 && mark < 90 {
		fmt.Print("B Grade")
	} else if mark >= 70 && mark < 80 {
		fmt.Print("C Grade")
	} else if mark >= 60 && mark < 70 {
		fmt.Print("D Grade")
	} else if mark >= 50 && mark < 60 {
		fmt.Print("E Grade")
	} else {
		fmt.Print("Failed")
	}
}

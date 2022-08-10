package main

import "fmt"

func main() {
	fmt.Print("  \n 1.Sunday \n 2.Monday \n 3.Tuesday \n 4.Wendsday \n 5.Thursday \n 6.Friday \n 7.Saturday \n")
	fmt.Print("Select the day : ")
	var day int
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Print("You are selected Sunday")
	case 2:
		fmt.Print("You are selected Monday")
	case 3:
		fmt.Print("You are selected Tuesday")
	case 4:
		fmt.Print("You are selected Wendsday")
	case 5:
		fmt.Print("You are selected Thursday")
	case 6:
		fmt.Print("You are selected Friday")
	case 7:
		fmt.Print("You are selected Saturday")
	default:
		fmt.Print("Invalid Entry")
	}
}

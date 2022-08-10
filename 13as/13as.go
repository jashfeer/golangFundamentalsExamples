package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the srting")
	var str string
	fmt.Scanln(&str)
	var length = len(str)
	flag := 0
	for i := 0; i < length; i++ {
		first := string(str[i])
		last := string(str[length-1-i])
		if first != last {
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("String is palindrome")
	} else {
		fmt.Println("String is not palindrome")
	}
}

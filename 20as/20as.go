package main

import "fmt"

func main() {
	p := 1
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("  ", p)
			p++
		}
		println()
	}
}

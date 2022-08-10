package main

import "fmt"

func main() {
	fmt.Println("Enter  mark of written test, lab exams and assignments :")
	var written, lab, assignment float32
	fmt.Scan(&written, &lab, &assignment)
	Grade := (written * 70 / 100) + (lab * 20 / 100) + (assignment * 10 / 100)
	fmt.Println("Grade of the student is : ", Grade)
}

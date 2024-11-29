package main

import "fmt"

func main() {
	var student struct {
		id    int
		name  string
		grade float32
	}
	student.id = 202444077
	student.name = "강태혁"
	student.grade = 4.5
	fmt.Println(student.grade)
}

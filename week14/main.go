package main

import "fmt"

type student struct {
	id    int
	name  string
	grade float32
}

func main() {
	var student1 student
	student1.id = 202444077
	student1.name = "강태혁"
	student1.grade = 4.5
	fmt.Println(student1.grade)

	var student2 student
	student2.id = 202444666
	student2.name = "니얼굴"
	student2.grade = 4.5
	fmt.Println(student2.id)
}

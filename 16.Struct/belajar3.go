package main

import "fmt"

type person struct {
	name string
	age  int
}

type student_embed struct {
	grade int
	age   int
	person
}

func main() {
	var s1 = student_embed{}
	s1.name = "Udin Bejo"
	s1.grade = 4
	s1.age = 19
	s1.person.age = 21

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("grade :", s1.grade)

	// Anonymous Struct

	var s4 = struct {
		person
		grade int
	}{}

	s4.person = person{"wick", 32}
	s4.grade = 4

	fmt.Println("name  :", s4.person.name)
	fmt.Println("age   :", s4.person.age)
	fmt.Println("grade :", s4.grade)
}

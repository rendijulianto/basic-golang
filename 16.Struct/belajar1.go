package main

import "fmt"

type studenta struct {
	name  string
	grade int
}

func main() {

	var s1 = studenta{}
	s1.name = "udin"
	s1.grade = 2
	var s2 = studenta{"ethan", 2}

	var s3 = studenta{name: "jason"}
	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

}

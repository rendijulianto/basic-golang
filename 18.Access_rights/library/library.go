package library

//func SayHello(name string) {
//	fmt.Println("Hallo Guys")
//	introduce(name)
//}
//
//func introduce(name string) {
//	fmt.Println("Nama Saya ", name)
//}

//type Student struct {
//	Name  string
//	Grade int
//}

import "fmt"

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "John Wick"
	Student.Grade = 2

	fmt.Println("--> library/library.go imported")
}

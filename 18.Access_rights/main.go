//package main
//
//import (
//	. "awesomeProject/18.Access_rights/library"
//	f "fmt"
//)
//
//func main() {
//	//library.SayHello("rendi")
//	//var s1 = library.Student{"Rendi", 21}
//	var s1 = Student{"Rendi", 21}
//	f.Println("Name", s1.Name)
//	f.Println("grade", s1.Grade)
//}

package main

import (
	"awesomeProject/18.Access_rights/library"
	"fmt"
)

func main() {
	//sayHello("ethan")
	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}

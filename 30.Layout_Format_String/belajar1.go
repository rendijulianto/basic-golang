package main

import "fmt"

type stundent struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

func main() {
	var data = stundent{
		name:        "Rendi Julianto",
		height:      168,
		age:         19,
		isGraduated: false,
		hobbies:     []string{"Ngoding", "Sleeping", "Eating"},
	}

	fmt.Printf("%b\n", data.age)
	//Octal
	fmt.Printf("%o\n", data.age)
	//Address pointer
	fmt.Printf("%p\n", &data.name)

	fmt.Printf("%q\n", `" name \ height "`)

	fmt.Printf("%s\n", data.name)

	fmt.Printf("%t\n", data.isGraduated)

	fmt.Printf("%T\n", data.name)
	// string

	fmt.Printf("%T\n", data.height)
	// float64

	fmt.Printf("%T\n", data.age)
	// int32

	fmt.Printf("%T\n", data.isGraduated)
	// bool

	fmt.Printf("%T\n", data.hobbies)
	// []string

	fmt.Printf("%v\n", data)

	fmt.Printf("%+v\n", data)

	fmt.Printf("%#v\n", data)

	fmt.Printf("%x\n", data.age)

}

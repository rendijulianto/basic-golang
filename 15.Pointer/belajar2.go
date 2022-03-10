package main

import "fmt"

func main() {
	var number = 4
	fmt.Println("Before : ", number)

	change(&number, 10)

	fmt.Println("After : ", number)

}

func change(original *int, value int) {
	*original = value
}

package main

import "fmt"

func main() {

	var point = 6

	switch point {
	case 8:
		fmt.Println("Perpect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

}

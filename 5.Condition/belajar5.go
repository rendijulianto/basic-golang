package main

import "fmt"

func main() {
	var point = 6

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("You Can Be Better")
		}
	}
}

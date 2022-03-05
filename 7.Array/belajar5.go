package main

import "fmt"

func main() {

	var fruits = [4]string{"Apple", "Grape", "BANANA", "MELON"}

	for i, fruit := range fruits {
		fmt.Printf("Element %d : %s\n", i, fruit)
	}
}

package main

import "fmt"

func main() {
	var fruits = [4]string{"Apple", "Grape", "Banana", "Melon"}

	for _, fruit := range fruits {
		fmt.Printf("Nama Buah : %s\n", fruit)
	}
}

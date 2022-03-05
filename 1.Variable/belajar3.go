package main

import "fmt"

func main() {
	//Multi Variable Deklarasi Type
	var first, second, third string

	first, second, third = "Satu", "Dua", "Tiga"

	fmt.Printf("Hallo %s %s %s", first, second, third)

	//Multi variable inference
	one, isFriday, twoPointTwo, Say := 1, true, 2.2, "Hallo Bang"
	fmt.Println("")
	fmt.Printf("Hallo %d %f %s", one, isFriday, twoPointTwo, Say)
}

package main

import "fmt"

func main() {

	var fruitsA = []string{"Apple", "Grape"}
	var fruitsB = [2]string{"Banana", "Melon"}
	var fruitsC = [...]string{"Papaya", "Grape"}
	fmt.Println(fruitsA)
	fmt.Println(fruitsB)
	fmt.Println(fruitsC)

}

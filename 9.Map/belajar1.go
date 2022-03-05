package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("Januari", chicken["januari"])
	fmt.Println("Mei", chicken["mei"])
}

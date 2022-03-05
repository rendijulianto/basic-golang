package main

import "fmt"

func main() {
	//var data map[string]int
	//data["one"] = 1
	//Akan muncul Error

	var data = map[string]int{}
	data["one"] = 1

	fmt.Println(data)
}

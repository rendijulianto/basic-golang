package main

import "fmt"

func main() {
	var person = []map[string]interface{}{
		{"name": "Rendi", "Age": 23},
		{"name": "Rendi A", "Age": 20},
		{"name": "BNpr", "Age": 21},
	}
	for _, each := range person {
		fmt.Println(each["name"], "Age is", each["Age"])
	}
}

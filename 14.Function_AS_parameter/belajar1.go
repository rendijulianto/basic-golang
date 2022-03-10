package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"Wick", "Jason", "Ethan", "Odino"}

	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Data asli \t\t:", data)
	fmt.Println("Filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("Filter Jumlah Huruf \"5\"\t:", dataLength5)

}

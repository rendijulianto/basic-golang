package main

import (
	"fmt"
	"math"
)

func main() {
	var area, cir = calculate1(10)
	fmt.Println(area)
	fmt.Println(cir)
}

func calculate1(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

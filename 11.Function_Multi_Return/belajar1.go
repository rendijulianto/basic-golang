package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("Luas Lingkaran \t\t: %.2f \n", area)
	fmt.Printf("Keliling Lingkaran \t: %.2f \n", circumference)
}

func calculate(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	//hitung keliling
	var circumference = math.Pi * d

	//kembalikan 2 nilai
	return area, circumference
}

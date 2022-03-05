package main

import "fmt"

func main() {
	//Seleksi Kondisi
	var point = 8

	if point == 10 {
		fmt.Println("Lulus Dengan Nilai Sangat Sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Println("Tidak lulus. nilai anda %d\n", point)
	}
}

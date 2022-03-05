package main

import "fmt"

func main() {

}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Println("Invalid")
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d \n", m, n, res)
}

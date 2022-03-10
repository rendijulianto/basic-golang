package main

import "fmt"

func main() {
	//Get Min Max
	var getMinMax = func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers = []int{2, 4, 56, 45, 23, 2, -1, 3, 5}
	var min, max = getMinMax(numbers)
	fmt.Printf("data :%v\n min : %v max : %d\n", numbers, min, max)
}

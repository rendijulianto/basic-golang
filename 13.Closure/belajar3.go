package main

import "fmt"

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

func main() {
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3, 9}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers \t:", numbers)
	fmt.Printf("find \t: %d\n", max)

	fmt.Println("Found \t:", howMany)
	fmt.Println("Value \t:", theNumbers)
}

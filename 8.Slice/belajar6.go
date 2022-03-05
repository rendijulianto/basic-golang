package main

import "fmt"

func main() {
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2
}

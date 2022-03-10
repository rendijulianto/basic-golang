package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(10)
	fmt.Println("Random ke-1:", rand.Int())
	fmt.Println("Random ke-2:", rand.Int())
	fmt.Println("Random ke-3:", rand.Int())

	fmt.Println("Unique Seed")
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Random ke-4:", rand.Int())
	fmt.Println("Random ke-5:", rand.Int())
	fmt.Println("Random ke-6:", rand.Int())
}

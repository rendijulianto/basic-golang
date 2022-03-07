package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Println("Time1 %v", time1)

	var time2 = time.Date(2022, 03, 07, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

}

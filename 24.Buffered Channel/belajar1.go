package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	messages := make(chan int, 2)

	go func() {
		for {
			i := <-messages
			fmt.Println("Receive Data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Send Data", i)
		messages <- i
	}
}

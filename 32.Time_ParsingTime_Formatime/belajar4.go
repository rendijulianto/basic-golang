package main

import (
	"fmt"
	"time"
)

func main() {
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(date.String())
}

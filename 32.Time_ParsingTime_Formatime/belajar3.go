package main

import (
	"fmt"
	"time"
)

func main() {

	var layoutFormat, value string

	var date time.Time

	layoutFormat = "2002-17-07 12:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())

}
